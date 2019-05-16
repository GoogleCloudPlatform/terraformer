package openstack

import (
	"strconv"

	"github.com/hashicorp/terraform/helper/resource"

	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/openstack/db/v1/configurations"
)

func expandDatabaseConfigurationV1Datastore(rawDatastore []interface{}) configurations.DatastoreOpts {
	v := rawDatastore[0].(map[string]interface{})
	datastore := configurations.DatastoreOpts{
		Version: v["version"].(string),
		Type:    v["type"].(string),
	}

	return datastore
}

func expandDatabaseConfigurationV1Values(rawValues []interface{}) map[string]interface{} {
	values := make(map[string]interface{})

	for _, rawValue := range rawValues {
		v := rawValue.(map[string]interface{})
		name := v["name"].(string)
		value := v["value"].(interface{})

		// check if value can be converted into int
		if valueInt, err := strconv.Atoi(value.(string)); err == nil {
			value = valueInt
		}

		values[name] = value
	}

	return values
}

// databaseConfigurationV1StateRefreshFunc returns a resource.StateRefreshFunc that is used to watch
// an cloud database instance.
func databaseConfigurationV1StateRefreshFunc(client *gophercloud.ServiceClient, cgroupID string) resource.StateRefreshFunc {
	return func() (interface{}, string, error) {
		i, err := configurations.Get(client, cgroupID).Extract()
		if err != nil {
			if _, ok := err.(gophercloud.ErrDefault404); ok {
				return i, "DELETED", nil
			}
			return nil, "", err
		}

		return i, "ACTIVE", nil
	}
}
