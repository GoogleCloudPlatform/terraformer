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

	"database/sql"
)

type client struct {
	db   *sql.DB
	Name string
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

type databaseGrant struct {
	CreatedOn   sql.NullString `db:"created_on"`
	Privilege   sql.NullString `db:"privilege"`
	GrantedOn   sql.NullString `db:"granted_on"`
	Name        sql.NullString `db:"name"`
	GrantedTo   sql.NullString `db:"granted_to"`
	GranteeName sql.NullString `db:"grantee_name"`
	GrantOption sql.NullString `db:"grant_option"`
	GrantedBy   sql.NullString `db:"granted_by"`
}

func (sc *client) ListDatabases() ([]database, error) {
	sdb := sqlx.NewDb(sc.db, "snowflake")
	stmt := "SHOW DATABASES"
	rows, err := sdb.Queryx(stmt)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	db := []database{}
	err = sqlx.StructScan(rows, &db)
	if err == sql.ErrNoRows {
		log.Printf("[DEBUG] no databases found")
		return nil, nil
	}
	return db, errors.Wrap(err, "unable to scan row for SHOW DATABASES")
}

func (sc *client) ListDatabaseGrants(database database) ([]databaseGrant, error) {
	sdb := sqlx.NewDb(sc.db, "snowflake")
	stmt := fmt.Sprintf(`SHOW GRANTS ON DATABASE "%v"`, database.DBName.String)
	rows, err := sdb.Queryx(stmt)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	dbGrants := []databaseGrant{}
	err = sqlx.StructScan(rows, &dbGrants)
	if err == sql.ErrNoRows {
		log.Printf("[DEBUG] no database grants found")
		return nil, nil
	}
	return dbGrants, errors.Wrap(err, "unable to scan row for SHOW DATABASES ON DATABASE")
}
