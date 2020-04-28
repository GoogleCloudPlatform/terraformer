package keycloak

import (
	"fmt"
	"strings"
)

type Group struct {
	Id          string              `json:"id,omitempty"`
	RealmId     string              `json:"-"`
	ParentId    string              `json:"-"`
	Name        string              `json:"name"`
	Path        string              `json:"path,omitempty"`
	SubGroups   []*Group            `json:"subGroups,omitempty"`
	RealmRoles  []string            `json:"realmRoles,omitempty"`
	ClientRoles map[string][]string `json:"clientRoles,omitempty"`
	Attributes  map[string][]string `json:"attributes"`
}

/*
 * There is no way to get a subgroup's parent ID using the Keycloak API (that I know of, PRs are welcome)
 * The best we can do is use the group's path to figure out what its parents' names are and iterate over all subgroups
 * until we find it.
 */
func (keycloakClient *KeycloakClient) groupParentId(group *Group) (string, error) {
	// Check the path of the group being passed in.
	// If there is only one group in the path, then this is a top-level group with no parentId
	parts := strings.Split(strings.TrimPrefix(group.Path, "/"), "/")

	if len(parts) == 1 {
		return "", nil
	}

	groups, err := keycloakClient.ListGroupsWithName(group.RealmId, group.Name)
	if err != nil {
		return "", err
	}

	currentGroups := &groups

	for index, groupName := range parts {
		for _, group := range *currentGroups {
			if group.Name == groupName {
				// if we're on the second to last index for the path, then this group must contain the passed in group as a child
				// thus, this group is the parent
				if index == len(parts)-2 {
					return group.Id, nil
				}

				currentGroups = &(group.SubGroups)

				break
			}
		}
	}

	// maybe panic here?  this should never happen
	return "", fmt.Errorf("unable to determine parent ID for group with path %s", group.Path)
}

func (keycloakClient *KeycloakClient) ValidateGroupMembers(usernames []interface{}) error {
	for _, username := range usernames {
		if username.(string) != strings.ToLower(username.(string)) {
			return fmt.Errorf("expected all usernames within group membership to be lowercase")
		}
	}

	return nil
}

/*
 * Top level groups are created via POST /realms/${realm_id}/groups
 * Child groups are created via POST /realms/${realm_id}/groups/${parent_id}/children
 */
func (keycloakClient *KeycloakClient) NewGroup(group *Group) error {
	var createGroupUrl string

	if group.ParentId == "" {
		createGroupUrl = fmt.Sprintf("/realms/%s/groups", group.RealmId)
	} else {
		createGroupUrl = fmt.Sprintf("/realms/%s/groups/%s/children", group.RealmId, group.ParentId)
	}

	_, location, err := keycloakClient.post(createGroupUrl, group)
	if err != nil {
		return err
	}

	group.Id = getIdFromLocationHeader(location)

	return nil
}

func (keycloakClient *KeycloakClient) GetGroups(realmId string) ([]*Group, error) {
	var groups []*Group

	err := keycloakClient.get(fmt.Sprintf("/realms/%s/groups", realmId), &groups, nil)
	if err != nil {
		return nil, err
	}

	for _, group := range groups {
		group.RealmId = realmId
	}

	return groups, nil
}

func (keycloakClient *KeycloakClient) GetGroup(realmId, id string) (*Group, error) {
	var group Group

	err := keycloakClient.get(fmt.Sprintf("/realms/%s/groups/%s", realmId, id), &group, nil)
	if err != nil {
		return nil, err
	}

	group.RealmId = realmId // it's important to set RealmId here because fetching the ParentId depends on it

	parentId, err := keycloakClient.groupParentId(&group)
	if err != nil {
		return nil, err
	}

	group.ParentId = parentId

	return &group, nil
}

func (keycloakClient *KeycloakClient) GetGroupByName(realmId, name string) (*Group, error) {
	var groups []Group

	// We can't get a group by name, so we have to search for it
	params := map[string]string{
		"search": name,
	}

	err := keycloakClient.get(fmt.Sprintf("/realms/%s/groups", realmId), &groups, params)
	if err != nil {
		return nil, err
	}

	if len(groups) == 0 {
		return nil, fmt.Errorf("no group with name " + name + " found")
	}

	// The search may return more than 1 result even if there is a group exactly matching the search string
	for _, group := range groups {
		if group.Name == name {

			group.RealmId = realmId // it's important to set RealmId here because fetching the ParentId depends on it

			parentId, err := keycloakClient.groupParentId(&group)
			if err != nil {
				return nil, err
			}

			group.ParentId = parentId

			return &group, nil
		}
	}

	return nil, fmt.Errorf("no group with name " + name + " found")
}

func (keycloakClient *KeycloakClient) UpdateGroup(group *Group) error {
	return keycloakClient.put(fmt.Sprintf("/realms/%s/groups/%s", group.RealmId, group.Id), group)
}

func (keycloakClient *KeycloakClient) DeleteGroup(realmId, id string) error {
	return keycloakClient.delete(fmt.Sprintf("/realms/%s/groups/%s", realmId, id), nil)
}

func (keycloakClient *KeycloakClient) ListGroupsWithName(realmId, name string) ([]*Group, error) {
	var groups []*Group

	params := map[string]string{
		"search": name,
	}

	err := keycloakClient.get(fmt.Sprintf("/realms/%s/groups", realmId), &groups, params)
	if err != nil {
		return nil, err
	}

	return groups, nil
}

func (keycloakClient *KeycloakClient) GetGroupMembers(realmId, groupId string) ([]*User, error) {
	var users []*User

	err := keycloakClient.get(fmt.Sprintf("/realms/%s/groups/%s/members?max=-1", realmId, groupId), &users, nil)
	if err != nil {
		return nil, err
	}

	for _, user := range users {
		user.RealmId = realmId
	}

	return users, nil
}

func defaultGroupURL(realmName, groupId string) string {
	return fmt.Sprintf("/realms/%s/default-groups/%s", realmName, groupId)
}

// PutDefaultGroup will PUT a new group ID to the realm default groups. This is effectively
// an "upsert".
func (keycloakClient *KeycloakClient) PutDefaultGroup(realmName, groupId string) error {
	url := defaultGroupURL(realmName, groupId)
	return keycloakClient.put(url, nil)
}

// DeleteDefaultGroup deletes a group ID from the realm default groups.
func (keycloakClient *KeycloakClient) DeleteDefaultGroup(realmName, groupId string) error {
	url := defaultGroupURL(realmName, groupId)
	return keycloakClient.delete(url, nil)
}

// GetDefaultGroups returns all the default groups for a realm.
func (keycloakClient *KeycloakClient) GetDefaultGroups(realmName string) ([]Group, error) {
	url := fmt.Sprintf("/realms/%s/default-groups", realmName)

	var defaultGroups []Group
	err := keycloakClient.get(url, &defaultGroups, nil)

	return defaultGroups, err
}

func (keycloakClient *KeycloakClient) AddRealmRolesToGroup(realmId, groupId string, roles []*Role) error {
	_, _, err := keycloakClient.post(fmt.Sprintf("/realms/%s/groups/%s/role-mappings/realm", realmId, groupId), roles)

	return err
}

func (keycloakClient *KeycloakClient) AddClientRolesToGroup(realmId, groupId, clientId string, roles []*Role) error {
	_, _, err := keycloakClient.post(fmt.Sprintf("/realms/%s/groups/%s/role-mappings/clients/%s", realmId, groupId, clientId), roles)

	return err
}

func (keycloakClient *KeycloakClient) RemoveRealmRolesFromGroup(realmId, groupId string, roles []*Role) error {
	err := keycloakClient.delete(fmt.Sprintf("/realms/%s/groups/%s/role-mappings/realm", realmId, groupId), roles)

	return err
}

func (keycloakClient *KeycloakClient) RemoveClientRolesFromGroup(realmId, groupId, clientId string, roles []*Role) error {
	err := keycloakClient.delete(fmt.Sprintf("/realms/%s/groups/%s/role-mappings/clients/%s", realmId, groupId, clientId), roles)

	return err
}
