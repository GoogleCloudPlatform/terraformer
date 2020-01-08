package keycloak

import (
	"github.com/hashicorp/errwrap"
	"net/http"
)

type ApiError struct {
	Code    int
	Message string
}

func (e *ApiError) Error() string {
	return e.Message
}

func ErrorIs404(err error) bool {
	keycloakError, ok := errwrap.GetType(err, &ApiError{}).(*ApiError)

	return ok && keycloakError != nil && keycloakError.Code == http.StatusNotFound
}
