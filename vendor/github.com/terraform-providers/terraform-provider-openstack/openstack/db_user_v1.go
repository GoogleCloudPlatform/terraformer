package openstack

import (
	"fmt"

	"github.com/hashicorp/terraform/helper/resource"

	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/openstack/db/v1/databases"
	"github.com/gophercloud/gophercloud/openstack/db/v1/users"
)

func expandDatabaseUserV1Databases(rawDatabases []interface{}) databases.BatchCreateOpts {
	var dbs databases.BatchCreateOpts

	for _, db := range rawDatabases {
		dbs = append(dbs, databases.CreateOpts{
			Name: db.(string),
		})
	}

	return dbs
}

func flattenDatabaseUserV1Databases(dbs []databases.Database) []string {
	var databases []string
	for _, db := range dbs {
		databases = append(databases, db.Name)
	}

	return databases
}

// databaseUserV1StateRefreshFunc returns a resource.StateRefreshFunc that is used to watch db user.
func databaseUserV1StateRefreshFunc(client *gophercloud.ServiceClient, instanceID string, userName string) resource.StateRefreshFunc {
	return func() (interface{}, string, error) {
		pages, err := users.List(client, instanceID).AllPages()
		if err != nil {
			return nil, "", fmt.Errorf("Unable to retrieve OpenStack database users: %s", err)
		}

		allUsers, err := users.ExtractUsers(pages)
		if err != nil {
			return nil, "", fmt.Errorf("Unable to extract OpenStack database users: %s", err)
		}

		for _, v := range allUsers {
			if v.Name == userName {
				return v, "ACTIVE", nil
			}
		}

		return nil, "BUILD", nil
	}
}

// databaseUserV1Exists is used to check whether user exists on particular database instance
func databaseUserV1Exists(client *gophercloud.ServiceClient, instanceID string, userName string) (bool, users.User, error) {
	var exists bool
	var err error
	var userObj users.User

	pages, err := users.List(client, instanceID).AllPages()
	if err != nil {
		return exists, userObj, err
	}

	allUsers, err := users.ExtractUsers(pages)
	if err != nil {
		return exists, userObj, err
	}

	for _, v := range allUsers {
		if v.Name == userName {
			exists = true
			return exists, v, nil
		}
	}

	return false, userObj, err
}
