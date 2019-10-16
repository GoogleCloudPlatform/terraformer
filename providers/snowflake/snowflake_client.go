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

type role struct {
	CreatedOn       sql.NullString `db:"created_on"`
	Name            sql.NullString `db:"name"`
	IsDefault       sql.NullString `db:"is_default"`
	IsCurrent       sql.NullString `db:"is_current"`
	IsInherited     sql.NullString `db:"is_inherited"`
	AssignedToUsers sql.NullInt32  `db:"assigned_to_users"`
	GrantedToRoles  sql.NullInt32  `db:"granted_to_roles"`
	GrantedRoles    sql.NullInt32  `db:"granted_roles"`
	Owner           sql.NullString `db:"owner"`
	Comment         sql.NullString `db:"comment"`
}

type user struct {
	Name               sql.NullString `db:"name"`
	CreatedOn          sql.NullString `db:"created_on"`
	LoginName          sql.NullString `db:"login_name"`
	DisplayName        sql.NullString `db:"display_name"`
	FirstName          sql.NullString `db:"first_name"`
	LastName           sql.NullString `db:"last_name"`
	Email              sql.NullString `db:"email"`
	MinsToUnlock       sql.NullString `db:"mins_to_unlock"`
	DaysToExpiry       sql.NullString `db:"days_to_expiry"`
	Comment            sql.NullString `db:"comment"`
	Disabled           sql.NullString `db:"disabled"`
	MustChangePassword sql.NullString `db:"must_change_password"`
	SnowflakeLock      sql.NullString `db:"snowflake_lock"`
	DefaultWarehouse   sql.NullString `db:"default_warehouse"`
	DefaultNamespace   sql.NullString `db:"default_namespace"`
	DefaultRole        sql.NullString `db:"default_role"`
	ExtAuthnDuo        sql.NullString `db:"ext_authn_duo"`
	ExtAuthnUid        sql.NullString `db:"ext_authn_uid"`
	MinsToBypassMfa    sql.NullString `db:"mins_to_bypass_mfa"`
	Owner              sql.NullString `db:"owner"`
	LastSuccessLogin   sql.NullString `db:"last_success_login"`
	ExpiresAtTime      sql.NullString `db:"expires_at_time"`
	LockedUntilTime    sql.NullString `db:"locked_until_time"`
	HasPassword        sql.NullString `db:"has_password"`
	HasRsaPublicKey    sql.NullString `db:"has_rsa_public_key"`
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

func (sc *client) ListRoles() ([]role, error) {
	sdb := sqlx.NewDb(sc.db, "snowflake")
	stmt := "SHOW ROLES"
	rows, err := sdb.Queryx(stmt)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	role := []role{}
	err = sqlx.StructScan(rows, &role)
	if err == sql.ErrNoRows {
		log.Printf("[DEBUG] no roles found")
		return nil, nil
	}
	return role, errors.Wrap(err, "unable to scan row for SHOW ROLES")
}

func (sc *client) ListUsers() ([]user, error) {
	sdb := sqlx.NewDb(sc.db, "snowflake")
	stmt := "SHOW USERS"
	rows, err := sdb.Queryx(stmt)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	user := []user{}
	err = sqlx.StructScan(rows, &user)
	if err == sql.ErrNoRows {
		log.Printf("[DEBUG] no users found")
		return nil, nil
	}
	return user, errors.Wrap(err, "unable to scan row for SHOW USERS")
}
