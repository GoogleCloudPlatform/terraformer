package openstack

import (
	"fmt"

	"github.com/hashicorp/terraform/helper/resource"

	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/openstack/db/v1/databases"
)

// databaseDatabaseV1StateRefreshFunc returns a resource.StateRefreshFunc
// that is used to watch a database.
func databaseDatabaseV1StateRefreshFunc(client *gophercloud.ServiceClient, instanceID string, dbName string) resource.StateRefreshFunc {
	return func() (interface{}, string, error) {
		pages, err := databases.List(client, instanceID).AllPages()
		if err != nil {
			return nil, "", fmt.Errorf("Unable to retrieve OpenStack databases: %s", err)
		}

		allDatabases, err := databases.ExtractDBs(pages)
		if err != nil {
			return nil, "", fmt.Errorf("Unable to extract OpenStack databases: %s", err)
		}

		for _, v := range allDatabases {
			if v.Name == dbName {
				return v, "ACTIVE", nil
			}
		}

		return nil, "BUILD", nil
	}
}

func databaseDatabaseV1Exists(client *gophercloud.ServiceClient, instanceID string, dbName string) (bool, error) {
	var exists bool
	var err error

	pages, err := databases.List(client, instanceID).AllPages()
	if err != nil {
		return exists, err
	}

	allDatabases, err := databases.ExtractDBs(pages)
	if err != nil {
		return exists, err
	}

	for _, v := range allDatabases {
		if v.Name == dbName {
			exists = true
			return exists, err
		}
	}

	return false, nil
}
