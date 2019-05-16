package openstack

import (
	"fmt"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/helper/schema"

	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/openstack/db/v1/databases"
	"github.com/gophercloud/gophercloud/openstack/db/v1/instances"
	"github.com/gophercloud/gophercloud/openstack/db/v1/users"
)

func expandDatabaseInstanceV1Datastore(rawDatastore []interface{}) instances.DatastoreOpts {
	v := rawDatastore[0].(map[string]interface{})
	datastore := instances.DatastoreOpts{
		Version: v["version"].(string),
		Type:    v["type"].(string),
	}

	return datastore
}

func expandDatabaseInstanceV1Networks(rawNetworks []interface{}) []instances.NetworkOpts {
	var networks []instances.NetworkOpts
	for _, v := range rawNetworks {
		network := v.(map[string]interface{})
		networks = append(networks, instances.NetworkOpts{
			UUID:      network["uuid"].(string),
			Port:      network["port"].(string),
			V4FixedIP: network["fixed_ip_v4"].(string),
			V6FixedIP: network["fixed_ip_v6"].(string),
		})
	}

	return networks
}

func expandDatabaseInstanceV1Databases(rawDatabases []interface{}) databases.BatchCreateOpts {
	var dbs databases.BatchCreateOpts
	for _, v := range rawDatabases {
		db := v.(map[string]interface{})
		dbs = append(dbs, databases.CreateOpts{
			Name:    db["name"].(string),
			CharSet: db["charset"].(string),
			Collate: db["collate"].(string),
		})
	}

	return dbs
}

func expandDatabaseInstanceV1Users(rawUsers []interface{}) users.BatchCreateOpts {
	var userList users.BatchCreateOpts
	for _, v := range rawUsers {
		user := v.(map[string]interface{})
		userList = append(userList, users.CreateOpts{
			Name:      user["name"].(string),
			Password:  user["password"].(string),
			Databases: expandInstanceV1UserDatabases(user["databases"].(*schema.Set).List()),
			Host:      user["host"].(string),
		})
	}

	return userList
}

// databaseInstanceV1StateRefreshFunc returns a resource.StateRefreshFunc
// that is used to watch a database instance.
func databaseInstanceV1StateRefreshFunc(client *gophercloud.ServiceClient, instanceID string) resource.StateRefreshFunc {
	return func() (interface{}, string, error) {
		i, err := instances.Get(client, instanceID).Extract()
		if err != nil {
			if _, ok := err.(gophercloud.ErrDefault404); ok {
				return i, "DELETED", nil
			}
			return nil, "", err
		}

		if i.Status == "error" {
			return i, i.Status, fmt.Errorf("There was an error creating the database instance.")
		}

		return i, i.Status, nil
	}
}

func expandInstanceV1UserDatabases(v []interface{}) databases.BatchCreateOpts {
	var dbs databases.BatchCreateOpts

	for _, db := range v {
		dbs = append(dbs, databases.CreateOpts{
			Name: db.(string),
		})
	}

	return dbs
}
