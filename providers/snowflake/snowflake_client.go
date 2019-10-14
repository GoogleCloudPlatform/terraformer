// Copyright 2019 The Terraformer Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package snowflake

import (
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	"github.com/snowflakedb/gosnowflake"

	"database/sql"
)

type client struct {
	db   *sql.DB
	Name string
}

func init() {
	fmt.Println("snowflake_client init")
	sql.Register("snowflake", &gosnowflake.SnowflakeDriver{})
}

type database struct {
	CreatedOn     sql.NullString `db:"created_on"`
	DBName        sql.NullString `db:"name"`
	IsDefault     sql.NullString `db:"is_default"`
	IsCurrent     sql.NullString `db:"is_current"`
	Origin        sql.NullString `db:"origin"`
	Owner         sql.NullString `db:"owner"`
	Comment       sql.NullString `db:"comment"`
	Options       sql.NullString `db:"options"`
	RetentionTime sql.NullString `db:"retention_time"`
}

// sc = Snowflake Client
func (sc *client) ListDatabases() ([]database, error) {
	fmt.Println("listdatabases")
	sdb := sqlx.NewDb(sc.db, "snowflake")
	stmt := "SHOW DATABASES"
	fmt.Println("queryx")
	rows, err := sdb.Queryx(stmt)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	fmt.Println("got some rows")
	db := []database{}
	fmt.Println("structscan in listdatabases")
	err = sqlx.StructScan(rows, &db)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Printf("[WARN] databases not found, removing from state file")
			return nil, nil
		}
		return nil, errors.Wrap(err, "unable to scan row for SHOW DATABASES")
	}
	return db, nil
}
