// Copyright (c) 2017-2019 Snowflake Computing Inc. All right reserved.

package gosnowflake

import (
	"context"
	"database/sql"
	"database/sql/driver"
	"encoding/json"
	"github.com/google/uuid"
	"net/url"
	"strconv"
	"strings"
	"sync/atomic"
)

const (
	statementTypeIDDml              = int64(0x3000)
	statementTypeIDInsert           = statementTypeIDDml + int64(0x100)
	statementTypeIDUpdate           = statementTypeIDDml + int64(0x200)
	statementTypeIDDelete           = statementTypeIDDml + int64(0x300)
	statementTypeIDMerge            = statementTypeIDDml + int64(0x400)
	statementTypeIDMultiTableInsert = statementTypeIDDml + int64(0x500)
)

const (
	sessionClientSessionKeepAlive          = "client_session_keep_alive"
	sessionClientValidateDefaultParameters = "CLIENT_VALIDATE_DEFAULT_PARAMETERS"
	serviceName                            = "service_name"
)

type snowflakeConn struct {
	cfg             *Config
	rest            *snowflakeRestful
	SequenceCounter uint64
	QueryID         string
	SQLState        string
}

// isDml returns true if the statement type code is in the range of DML.
func (sc *snowflakeConn) isDml(v int64) bool {
	switch v {
	case statementTypeIDDml, statementTypeIDInsert,
		statementTypeIDUpdate, statementTypeIDDelete,
		statementTypeIDMerge, statementTypeIDMultiTableInsert:
		return true
	}
	return false
}

func (sc *snowflakeConn) exec(
	ctx context.Context,
	query string, noResult bool, isInternal bool, parameters []driver.NamedValue) (*execResponse, error) {
	var err error
	counter := atomic.AddUint64(&sc.SequenceCounter, 1) // query sequence counter

	req := execRequest{
		SQLText:    query,
		AsyncExec:  noResult,
		SequenceID: counter,
	}
	req.IsInternal = isInternal
	tsmode := "TIMESTAMP_NTZ"
	idx := 1
	if len(parameters) > 0 {
		req.Bindings = make(map[string]execBindParameter, len(parameters))
		for i, n := 0, len(parameters); i < n; i++ {
			t := goTypeToSnowflake(parameters[i].Value, tsmode)
			glog.V(2).Infof("tmode: %v\n", t)
			if t == "CHANGE_TYPE" {
				tsmode, err = dataTypeMode(parameters[i].Value)
				if err != nil {
					return nil, err
				}
			} else {
				v1, err := valueToString(parameters[i].Value, tsmode)
				if err != nil {
					return nil, err
				}
				req.Bindings[strconv.Itoa(idx)] = execBindParameter{
					Type:  t,
					Value: v1,
				}
				idx++
			}
		}
	}
	glog.V(2).Infof("bindings: %v", req.Bindings)

	headers := make(map[string]string)
	headers["Content-Type"] = headerContentTypeApplicationJSON
	headers["accept"] = headerAcceptTypeApplicationSnowflake // TODO v1.1: change to JSON in case of PUT/GET
	headers["User-Agent"] = userAgent
	if serviceName, ok := sc.cfg.Params[serviceName]; ok {
		headers["X-Snowflake-Service"] = *serviceName
	}

	jsonBody, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	var data *execResponse

	requestID := uuid.New()
	data, err = sc.rest.FuncPostQuery(ctx, sc.rest, &url.Values{}, headers, jsonBody, sc.rest.RequestTimeout, &requestID)
	if err != nil {
		return nil, err
	}
	var code int
	if data.Code != "" {
		code, err = strconv.Atoi(data.Code)
		if err != nil {
			code = -1
			return nil, err
		}
	} else {
		code = -1
	}
	glog.V(2).Infof("Success: %v, Code: %v", data.Success, code)
	if !data.Success {
		return nil, &SnowflakeError{
			Number:   code,
			SQLState: data.Data.SQLState,
			Message:  data.Message,
			QueryID:  data.Data.QueryID,
		}
	}
	glog.V(2).Info("Exec/Query SUCCESS")
	sc.cfg.Database = data.Data.FinalDatabaseName
	sc.cfg.Schema = data.Data.FinalSchemaName
	sc.cfg.Role = data.Data.FinalRoleName
	sc.cfg.Warehouse = data.Data.FinalWarehouseName
	sc.QueryID = data.Data.QueryID
	sc.SQLState = data.Data.SQLState
	sc.populateSessionParameters(data.Data.Parameters)
	return data, err
}

func (sc *snowflakeConn) Begin() (driver.Tx, error) {
	return sc.BeginTx(context.TODO(), driver.TxOptions{})
}

func (sc *snowflakeConn) BeginTx(ctx context.Context, opts driver.TxOptions) (driver.Tx, error) {
	glog.V(2).Info("BeginTx")
	if opts.ReadOnly {
		return nil, &SnowflakeError{
			Number:   ErrNoReadOnlyTransaction,
			SQLState: SQLStateFeatureNotSupported,
			Message:  errMsgNoReadOnlyTransaction,
		}
	}
	if int(opts.Isolation) != int(sql.LevelDefault) {
		return nil, &SnowflakeError{
			Number:   ErrNoDefaultTransactionIsolationLevel,
			SQLState: SQLStateFeatureNotSupported,
			Message:  errMsgNoDefaultTransactionIsolationLevel,
		}
	}
	if sc.rest == nil {
		return nil, driver.ErrBadConn
	}
	_, err := sc.exec(ctx, "BEGIN", false, false, nil)
	if err != nil {
		return nil, err
	}
	return &snowflakeTx{sc}, err
}

func (sc *snowflakeConn) cleanup() {
	glog.Flush() // must flush log buffer while the process is running.
	sc.rest = nil
	sc.cfg = nil
}

func (sc *snowflakeConn) Close() (err error) {
	glog.V(2).Infoln("Close")
	sc.stopHeartBeat()

	err = sc.rest.FuncCloseSession(sc.rest)
	if err != nil {
		glog.V(2).Info(err)
	}
	sc.cleanup()
	return nil
}
func (sc *snowflakeConn) PrepareContext(ctx context.Context, query string) (driver.Stmt, error) {
	glog.V(2).Infoln("Prepare")
	if sc.rest == nil {
		return nil, driver.ErrBadConn
	}
	stmt := &snowflakeStmt{
		sc:    sc,
		query: query,
	}
	return stmt, nil
}

func (sc *snowflakeConn) Prepare(query string) (driver.Stmt, error) {
	return sc.PrepareContext(context.TODO(), query)
}

func (sc *snowflakeConn) ExecContext(ctx context.Context, query string, args []driver.NamedValue) (driver.Result, error) {
	glog.V(2).Infof("Exec: %#v, %v", query, args)
	if sc.rest == nil {
		return nil, driver.ErrBadConn
	}
	// TODO: handle noResult and isInternal
	data, err := sc.exec(ctx, query, false, false, args)
	if err != nil {
		return nil, err
	}
	var updatedRows int64
	if sc.isDml(data.Data.StatementTypeID) {
		// collects all values from the returned row sets
		updatedRows = 0
		for i, n := 0, len(data.Data.RowType); i < n; i++ {
			v, err := strconv.ParseInt(*data.Data.RowSet[0][i], 10, 64)
			if err != nil {
				return nil, err
			}
			updatedRows += v
		}
		glog.V(2).Infof("number of updated rows: %#v", updatedRows)
		return &snowflakeResult{
			affectedRows: updatedRows,
			insertID:     -1}, nil // last insert id is not supported by Snowflake
	}
	glog.V(2).Info("DDL")
	return driver.ResultNoRows, nil
}

func (sc *snowflakeConn) QueryContext(ctx context.Context, query string, args []driver.NamedValue) (driver.Rows, error) {
	glog.V(2).Infof("Query: %#v, %v", query, args)
	if sc.rest == nil {
		return nil, driver.ErrBadConn
	}
	// TODO: handle noResult and isInternal
	data, err := sc.exec(ctx, query, false, false, args)
	if err != nil {
		glog.V(2).Infof("error: %v", err)
		return nil, err
	}

	rows := new(snowflakeRows)
	rows.sc = sc
	rows.RowType = data.Data.RowType
	rows.ChunkDownloader = &snowflakeChunkDownloader{
		sc:                 sc,
		ctx:                ctx,
		CurrentChunk:       data.Data.RowSet,
		ChunkMetas:         data.Data.Chunks,
		Total:              int64(data.Data.Total),
		TotalRowIndex:      int64(-1),
		CellCount:          len(data.Data.RowType),
		Qrmk:               data.Data.Qrmk,
		ChunkHeader:        data.Data.ChunkHeaders,
		FuncDownload:       downloadChunk,
		FuncDownloadHelper: downloadChunkHelper,
		FuncGet:            getChunk,
	}
	rows.ChunkDownloader.start()
	return rows, err
}

func (sc *snowflakeConn) Exec(
	query string,
	args []driver.Value) (
	driver.Result, error) {
	return sc.ExecContext(context.TODO(), query, toNamedValues(args))
}

func (sc *snowflakeConn) Query(
	query string,
	args []driver.Value) (
	driver.Rows, error) {
	return sc.QueryContext(context.TODO(), query, toNamedValues(args))
}

func (sc *snowflakeConn) Ping(ctx context.Context) error {
	glog.V(2).Infoln("Ping")
	if sc.rest == nil {
		return driver.ErrBadConn
	}
	// TODO: handle noResult and isInternal
	_, err := sc.exec(ctx, "SELECT 1", false, false, []driver.NamedValue{})
	return err
}

func (sc *snowflakeConn) populateSessionParameters(parameters []nameValueParameter) {
	// other session parameters (not all)
	glog.V(2).Infof("params: %#v", parameters)
	for _, param := range parameters {
		v := ""
		switch param.Value.(type) {
		case int64:
			if vv, ok := param.Value.(int64); ok {
				v = strconv.FormatInt(vv, 10)
			}
		case float64:
			if vv, ok := param.Value.(float64); ok {
				v = strconv.FormatFloat(vv, 'g', -1, 64)
			}
		case bool:
			if vv, ok := param.Value.(bool); ok {
				v = strconv.FormatBool(vv)
			}
		default:
			if vv, ok := param.Value.(string); ok {
				v = vv
			}
		}
		glog.V(3).Infof("parameter. name: %v, value: %v", param.Name, v)
		sc.cfg.Params[strings.ToLower(param.Name)] = &v
	}
}

func (sc *snowflakeConn) isClientSessionKeepAliveEnabled() bool {
	v, ok := sc.cfg.Params[sessionClientSessionKeepAlive]
	if !ok {
		return false
	}
	return strings.Compare(*v, "true") == 0
}

func (sc *snowflakeConn) startHeartBeat() {
	if !sc.isClientSessionKeepAliveEnabled() {
		return
	}
	sc.rest.HeartBeat = &heartbeat{
		restful: sc.rest,
	}
	sc.rest.HeartBeat.start()
}

func (sc *snowflakeConn) stopHeartBeat() {
	if !sc.isClientSessionKeepAliveEnabled() {
		return
	}
	sc.rest.HeartBeat.stop()
}
