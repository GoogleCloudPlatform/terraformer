package keycloak

import (
	"fmt"
)

type FederatedIdentity struct {
	IdentityProvider string `json:"identityProvider"`
	UserId           string `json:"userId"`
	UserName         string `json:"userName"`
}

type FederatedIdentities []FederatedIdentity

type User struct {
	Id      string `json:"id,omitempty"`
	RealmId string `json:"-"`

	Username            string              `json:"username"`
	Email               string              `json:"email"`
	FirstName           string              `json:"firstName"`
	LastName            string              `json:"lastName"`
	Enabled             bool                `json:"enabled"`
	Attributes          map[string][]string `json:"attributes"`
	FederatedIdentities FederatedIdentities `json:"federatedIdentities"`
}

type PasswordCredentials struct {
	Value     string `json:"value"`
	Type      string `json:"type"`
	Temporary bool   `json:"temporary"`
}

func (keycloakClient *KeycloakClient) NewUser(user *User) error {
	_, location, err := keycloakClient.post(fmt.Sprintf("/realms/%s/users", user.RealmId), user)
	if err != nil {
		return err
	}

	user.Id = getIdFromLocationHeader(location)

	return nil
}

func (keycloakClient *KeycloakClient) ResetUserPassword(realmId, userId string, newPassword string, isTemporary bool) error {
	resetCredentials := &PasswordCredentials{
		Value:     newPassword,
		Type:      "password",
		Temporary: isTemporary,
	}

	err := keycloakClient.put(fmt.Sprintf("/realms/%s/users/%s/reset-password", realmId, userId), resetCredentials)
	if err != nil {
		return err
	}
	return nil
}

func (keycloakClient *KeycloakClient) GetUsers(realmId string) ([]*User, error) {
	var users []*User

	err := keycloakClient.get(fmt.Sprintf("/realms/%s/users", realmId), &users, nil)
	if err != nil {
		return nil, err
	}

	for _, user := range users {
		user.RealmId = realmId
	}

	return users, nil
}

func (keycloakClient *KeycloakClient) GetUsersRoles(realmId string) ([]*Role, error) {
	var roles []*Role
	var users []*User

	err := keycloakClient.get(fmt.Sprintf("/realms/%s/users", realmId), &users, nil)
	if err != nil {
		return nil, err
	}

	for _, user := range users {
		var roles_user []*Role
		err = keycloakClient.get(fmt.Sprintf("/realms/%s/users/%s/roles", realmId, user.Id), &roles_user, nil)
		if err != nil {
			return nil, err
		}
		roles = append(roles, roles_user...)
	}

	for _, role := range roles {
		role.RealmId = realmId
	}

	return roles, nil
}

func (keycloakClient *KeycloakClient) GetUser(realmId, id string) (*User, error) {
	var user User

	err := keycloakClient.get(fmt.Sprintf("/realms/%s/users/%s", realmId, id), &user, nil)
	if err != nil {
		return nil, err
	}

	user.RealmId = realmId

	return &user, nil
}

func (keycloakClient *KeycloakClient) UpdateUser(user *User) error {
	return keycloakClient.put(fmt.Sprintf("/realms/%s/users/%s", user.RealmId, user.Id), user)
}

func (keycloakClient *KeycloakClient) DeleteUser(realmId, id string) error {
	return keycloakClient.delete(fmt.Sprintf("/realms/%s/users/%s", realmId, id), nil)
}

func (keycloakClient *KeycloakClient) GetUserByUsername(realmId, username string) (*User, error) {
	var users []*User

	params := map[string]string{
		"username": username,
	}

	err := keycloakClient.get(fmt.Sprintf("/realms/%s/users", realmId), &users, params)
	if err != nil {
		return nil, err
	}

	// more than one user could be returned so we need to search through all of the results and return the correct one
	// ex: foo and foo-user could both exist, but searching for "foo" will return both
	for _, user := range users {
		if user.Username == username {
			user.RealmId = realmId

			return user, nil
		}
	}

	// the requested user does not exist
	// we shouldn't raise an error here since it will be difficult to differentiate between a non-existent user and a network error
	return nil, nil
}

func (keycloakClient *KeycloakClient) addUserToGroup(user *User, groupId string) error {
	return keycloakClient.put(fmt.Sprintf("/realms/%s/users/%s/groups/%s", user.RealmId, user.Id, groupId), nil)
}

func (keycloakClient *KeycloakClient) AddUsersToGroup(realmId, groupId string, users []interface{}) error {
	for _, username := range users {
		user, err := keycloakClient.GetUserByUsername(realmId, username.(string)) // we need the user's id in order to add them to a group
		if err != nil {
			return err
		}
		if user == nil {
			return fmt.Errorf("user with username %s does not exist", username.(string))
		}

		err = keycloakClient.addUserToGroup(user, groupId)
		if err != nil {
			return err
		}
	}

	return nil
}

func (keycloakClient *KeycloakClient) RemoveUserFromGroup(user *User, groupId string) error {
	return keycloakClient.delete(fmt.Sprintf("/realms/%s/users/%s/groups/%s", user.RealmId, user.Id, groupId), nil)
}

func (keycloakClient *KeycloakClient) RemoveUsersFromGroup(realmId, groupId string, usernames []interface{}) error {
	for _, username := range usernames {
		user, err := keycloakClient.GetUserByUsername(realmId, username.(string)) // we need the user's id in order to remove them from a group
		if err != nil {
			return err
		}
		if user == nil {
			return fmt.Errorf("user with username %s does not exist", username.(string))
		}

		err = keycloakClient.RemoveUserFromGroup(user, groupId)
		if err != nil {
			return err
		}
	}

	return nil
}
