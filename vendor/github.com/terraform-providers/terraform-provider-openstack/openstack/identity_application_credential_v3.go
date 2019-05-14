package openstack

import (
	"github.com/gophercloud/gophercloud/openstack/identity/v3/applicationcredentials"
)

func flattenIdentityApplicationCredentialRolesV3(roles []applicationcredentials.Role) []string {
	var res []string
	for _, role := range roles {
		res = append(res, role.Name)
	}
	return res
}

func expandIdentityApplicationCredentialRolesV3(roles []interface{}) []applicationcredentials.Role {
	var res []applicationcredentials.Role
	for _, role := range roles {
		res = append(res, applicationcredentials.Role{Name: role.(string)})
	}
	return res
}
