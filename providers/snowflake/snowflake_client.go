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

type TfGrant struct {
	Name      string
	Privilege string
	Users     []string
	Roles     []string
	Shares    []string
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

type warehouse struct {
	Name            sql.NullString `db:"name"`
	State           sql.NullString `db:"state"`
	Type            sql.NullString `db:"type"`
	Size            sql.NullString `db:"size"`
	MinClusterCount sql.NullInt64  `db:"min_cluster_count"`
	MaxClusterCount sql.NullInt64  `db:"max_cluster_count"`
	StartedClusters sql.NullInt64  `db:"started_clusters"`
	Running         sql.NullInt64  `db:"running"`
	Queued          sql.NullInt64  `db:"queued"`
	IsDefault       sql.NullString `db:"is_default"`
	IsCurrent       sql.NullString `db:"is_current"`
	AutoSuspend     sql.NullInt64  `db:"auto_suspend"`
	AutoResume      sql.NullBool   `db:"auto_resume"`
	Available       sql.NullString `db:"available"`
	Provisioning    sql.NullString `db:"provisioning"`
	Quiescing       sql.NullString `db:"quiescing"`
	Other           sql.NullString `db:"other"`
	CreatedOn       sql.NullTime   `db:"created_on"`
	ResumedOn       sql.NullTime   `db:"resumed_on"`
	UpdatedOn       sql.NullTime   `db:"updated_on"`
	Owner           sql.NullString `db:"owner"`
	Comment         sql.NullString `db:"comment"`
	ResourceMonitor sql.NullString `db:"resource_monitor"`
	Actives         sql.NullInt64  `db:"actives"`
	Pendings        sql.NullInt64  `db:"pendings"`
	Failed          sql.NullInt64  `db:"failed"`
	Suspended       sql.NullInt64  `db:"suspended"`
	UUID            sql.NullString `db:"uuid"`
	ScalingPolicy   sql.NullString `db:"scaling_policy"`
}

type schema struct {
	CreatedOn     sql.NullTime   `db:"created_on"`
	Name          sql.NullString `db:"name"`
	IsDefault     sql.NullString `db:"is_default"`
	IsCurrent     sql.NullString `db:"is_current"`
	DatabaseName  sql.NullString `db:"database_name"`
	Owner         sql.NullString `db:"owner"`
	Comment       sql.NullString `db:"comment"`
	Options       sql.NullString `db:"options"`
	RetentionTime sql.NullString `db:"retention_time"`
}

type schemaGrant struct {
	CreatedOn   sql.NullString `db:"created_on"`
	Privilege   sql.NullString `db:"privilege"`
	GrantedOn   sql.NullString `db:"granted_on"`
	Name        sql.NullString `db:"name"`
	GrantedTo   sql.NullString `db:"granted_to"`
	GranteeName sql.NullString `db:"grantee_name"`
	GrantOption sql.NullString `db:"grant_option"`
	GrantedBy   sql.NullString `db:"granted_by"`
}

type roleGrant struct {
	CreatedOn   sql.NullString `db:"created_on"`
	Role        sql.NullString `db:"role"`
	GrantedTo   sql.NullString `db:"granted_to"`
	GranteeName sql.NullString `db:"grantee_name"`
	GrantedBy   sql.NullString `db:"granted_by"`
}

type warehouseGrant struct {
	CreatedOn   sql.NullString `db:"created_on"`
	Privilege   sql.NullString `db:"privilege"`
	GrantedOn   sql.NullString `db:"granted_on"`
	Name        sql.NullString `db:"name"`
	GrantedTo   sql.NullString `db:"granted_to"`
	GranteeName sql.NullString `db:"grantee_name"`
	GrantOption sql.NullString `db:"grant_option"`
	GrantedBy   sql.NullString `db:"granted_by"`
}

type view struct {
	CreatedOn      sql.NullString `db:"created_on"`
	Name           sql.NullString `db:"name"`
	Reserved       sql.NullString `db:"reserved"`
	DatabaseName   sql.NullString `db:"database_name"`
	SchemaName     sql.NullString `db:"schema_name"`
	Owner          sql.NullString `db:"owner"`
	Comment        sql.NullString `db:"comment"`
	Text           sql.NullString `db:"text"`
	IsSecure       sql.NullString `db:"is_secure"`
	IsMaterialized sql.NullString `db:"is_materialized"`
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
	stmt := fmt.Sprintf(`SHOW GRANTS ON DATABASE "%s"`, database.DBName.String)
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

func (sc *client) ListWarehouses() ([]warehouse, error) {
	sdb := sqlx.NewDb(sc.db, "snowflake")
	stmt := "SHOW WAREHOUSES"
	rows, err := sdb.Queryx(stmt)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	warehouse := []warehouse{}
	err = sqlx.StructScan(rows, &warehouse)
	if err == sql.ErrNoRows {
		log.Printf("[DEBUG] no WAREHOUSES found")
		return nil, nil
	}
	return warehouse, errors.Wrap(err, "unable to scan row for SHOW WAREHOUSES")
}

func (sc *client) ListSchemas(database *database) ([]schema, error) {
	sdb := sqlx.NewDb(sc.db, "snowflake")
	stmt := "SHOW SCHEMAS"
	if database != nil {
		stmt = fmt.Sprintf(`%s IN DATABASE "%s"`, stmt, database.DBName.String)
	}
	rows, err := sdb.Queryx(stmt)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	schema := []schema{}
	err = sqlx.StructScan(rows, &schema)
	if err == sql.ErrNoRows {
		log.Printf("[DEBUG] no SCHEMAS found")
		return nil, nil
	}
	return schema, errors.Wrap(err, "unable to scan row for SHOW SCHEMAS")
}

func (sc *client) ListSchemaGrants(database database, schema schema) ([]schemaGrant, error) {
	sdb := sqlx.NewDb(sc.db, "snowflake")
	stmt := fmt.Sprintf(`SHOW GRANTS ON SCHEMA "%s"."%s"`, database.DBName.String, schema.Name.String)
	rows, err := sdb.Queryx(stmt)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	schemaGrants := []schemaGrant{}
	err = sqlx.StructScan(rows, &schemaGrants)
	if err == sql.ErrNoRows {
		log.Printf("[DEBUG] no schema grants found")
		return nil, nil
	}
	return schemaGrants, errors.Wrap(err, "unable to scan row for SHOW GRANTS ON SCHEMA")
}

func (sc *client) ListRoleGrants(role role) ([]roleGrant, error) {
	sdb := sqlx.NewDb(sc.db, "snowflake")
	stmt := fmt.Sprintf(`SHOW GRANTS OF ROLE "%s"`, role.Name.String)
	rows, err := sdb.Queryx(stmt)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	roleGrants := []roleGrant{}
	err = sqlx.StructScan(rows, &roleGrants)
	if err == sql.ErrNoRows {
		log.Printf("[DEBUG] no role grants found")
		return nil, nil
	}
	return roleGrants, errors.Wrap(err, "unable to scan row for SHOW GRANTS OF ROLE")
}

func (sc *client) ListWarehouseGrants(warehouse warehouse) ([]warehouseGrant, error) {
	sdb := sqlx.NewDb(sc.db, "snowflake")
	stmt := fmt.Sprintf(`SHOW GRANTS ON WAREHOUSE "%s"`, warehouse.Name.String)
	rows, err := sdb.Queryx(stmt)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	warehouseGrants := []warehouseGrant{}
	err = sqlx.StructScan(rows, &warehouseGrants)
	if err == sql.ErrNoRows {
		log.Printf("[DEBUG] no warehouse grants found")
		return nil, nil
	}
	return warehouseGrants, errors.Wrap(err, "unable to scan row for SHOW GRANTS ON WAREHOUSE")
}

func (sc *client) ListViews() ([]view, error) {
	sdb := sqlx.NewDb(sc.db, "snowflake")
	stmt := "SHOW VIEWS"
	rows, err := sdb.Queryx(stmt)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	view := []view{}
	err = sqlx.StructScan(rows, &view)
	if err == sql.ErrNoRows {
		log.Printf("[DEBUG] no VIEWS found")
		return nil, nil
	}
	return view, errors.Wrap(err, "unable to scan row for SHOW VIEWS")
}
