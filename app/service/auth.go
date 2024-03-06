package service

import (
	"context"
	"net/http"

	"github.com/rs/zerolog"
	"golang.org/x/oauth2"
	"kms/app/entity"
)

// PermissionServicer allows for creating, updating, reading and deleting a Permission
type PermissionServicer interface {
	Create(ctx context.Context, r *CreatePermissionRequest, adt entity.Audit) (*PermissionResponse, error)
	FindAll(ctx context.Context) ([]*PermissionResponse, error)
	Delete(ctx context.Context, extlID string) (DeleteResponse, error)
}

// RoleServicer allows for creating, updating, reading and deleting a Role
// as well as assigning permissions and users to it.
type RoleServicer interface {
	Create(ctx context.Context, r *CreateRoleRequest, adt entity.Audit) (*RoleResponse, error)
}

// AuthenticationServicer represents a service for managing authentication.
//
// For this project, Oauth2 is used for user authentication. It is assumed
// that the actual user interaction is being orchestrated externally and
// the server endpoints are being called after an access token has already
// been retrieved from an authentication provider.
//
// In addition, this project provides for a custom application authentication.
// If an endpoint request is sent using application credentials, then those
// will be used. If none are sent, then the client id from the access token
// must be registered in the system and that is used as the calling application.
// The latter is likely the more common use case.
type AuthenticationServicer interface {

	// SelfRegister is used for first-time registration of a Person/User
	// in the system (associated with an Organization). This is "self
	// registration" as opposed to one person registering another person.
	SelfRegister(ctx context.Context, params *entity.AuthenticationParams) (ur *UserResponse, err error)

	// FindExistingAuth looks up a User given a Provider and Access Token.
	// If a User is not found, an error is returned.
	FindExistingAuth(r *http.Request, realm string) (entity.Auth, error)

	// FindAppByProviderClientID Finds an App given a Provider Client ID as part
	// of an Auth object.
	FindAppByProviderClientID(ctx context.Context, realm string, auth entity.Auth) (a *entity.App, err error)

	// DetermineAppContext checks to see if the request already has an app as part of
	// if it does, use that app as the app for session, if it does not, determine the
	// app based on the user's provider client ID. In either case, return a new context
	// with an app. If there is no app to be found for either, return an error.
	DetermineAppContext(ctx context.Context, auth entity.Auth, realm string) (context.Context, error)

	// FindAppByAPIKey finds an app given its External ID and determines
	// if the given API key is a valid key for it. It is used as part of
	// app authentication.
	FindAppByAPIKey(r *http.Request, realm string) (*entity.App, error)

	// AuthenticationParamExchange returns a ProviderInfo struct
	// after calling remote Oauth2 provider.
	AuthenticationParamExchange(ctx context.Context, params *entity.AuthenticationParams) (*entity.ProviderInfo, error)

	// NewAuthenticationParams parses the provider and authorization
	// headers and returns AuthenticationParams based on the results
	NewAuthenticationParams(r *http.Request, realm string) (*entity.AuthenticationParams, error)
}

// AuthorizationServicer represents a service for managing authorization.
type AuthorizationServicer interface {
	Authorize(r *http.Request, lgr zerolog.Logger, adt entity.Audit) error
}

// TokenExchanger exchanges an oauth2.Token for a ProviderUserInfo
// struct populated with information retrieved from an authentication provider.
type TokenExchanger interface {
	Exchange(ctx context.Context, realm string, provider entity.Provider, token *oauth2.Token) (*entity.ProviderInfo, error)
}

// CreatePermissionRequest is the request struct for creating a permission
type CreatePermissionRequest struct {
	// A human-readable string which represents a resource (e.g. an HTTP route or document, etc.).
	Resource string `json:"resource"`
	// A string representing the action taken on the resource (e.g. POST, GET, edit, etc.)
	Operation string `json:"operation"`
	// A description of what the permission is granting, e.g. "grants ability to edit a billing document".
	Description string `json:"description"`
	// A boolean denoting whether the permission is active (true) or not (false).
	Active bool `json:"active"`
}

// FindPermissionRequest is the response struct for finding a permission
type FindPermissionRequest struct {
	// Unique External ID to be given to outside callers.
	ExternalID string `json:"external_id"`
	// A human-readable string which represents a resource (e.g. an HTTP route or document, etc.).
	Resource string `json:"resource"`
	// A string representing the action taken on the resource (e.g. POST, GET, edit, etc.)
	Operation string `json:"operation"`
}

// PermissionResponse is the response struct for a permission
type PermissionResponse struct {
	// Unique External ID to be given to outside callers.
	ExternalID string `json:"external_id"`
	// A human-readable string which represents a resource (e.g. an HTTP route or document, etc.).
	Resource string `json:"resource"`
	// A string representing the action taken on the resource (e.g. POST, GET, edit, etc.)
	Operation string `json:"operation"`
	// A description of what the permission is granting, e.g. "grants ability to edit a billing document".
	Description string `json:"description"`
	// A boolean denoting whether the permission is active (true) or not (false).
	Active bool `json:"active"`
}

// CreateRoleRequest is the request struct for creating a role
type CreateRoleRequest struct {
	// A human-readable code which represents the role.
	Code string `json:"role_cd"`
	// A longer description of the role.
	Description string `json:"role_description"`
	// A boolean denoting whether the role is active (true) or not (false).
	Active bool `json:"active"`
	// The list of permissions to be given to the role
	Permissions []*FindPermissionRequest
}

// RoleResponse is the response struct for a Role.
type RoleResponse struct {
	// Unique External ID to be given to outside callers.
	ExternalID string `json:"external_id"`
	// A human-readable code which represents the role.
	Code string `json:"role_cd"`
	// A longer description of the role.
	Description string `json:"role_description"`
	// A boolean denoting whether the role is active (true) or not (false).
	Active bool `json:"active"`
	// Permissions is the list of permissions allowed for the role.
	Permissions []*entity.Permission
}
