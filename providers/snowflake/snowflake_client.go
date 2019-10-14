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
	if err != nil {
		if err == sql.ErrNoRows {
			log.Printf("[WARN] no databases found")
			return nil, nil
		}
		return nil, errors.Wrap(err, "unable to scan row for SHOW DATABASES")
	}
	return db, nil
}
