package entity

import (
	"kms/app/errs"
	"kms/app/secure"
	"strings"
	"time"

	"github.com/google/uuid"
	"golang.org/x/oauth2"
)

// BearerTokenType is used in authorization to access a resource
const BearerTokenType string = "Bearer"

// Provider defines the provider of authorization (Google, Github, Apple, auth0, etc.).
//
// Only Google is used currently.
type Provider uint8

// Provider of authorization
//
// The app uses Oauth2 to authorize users with one of the following Providers
const (
	UnknownProvider Provider = iota
	Google                   // Google
)

func (p Provider) String() string {
	switch p {
	case Google:
		return "google"
	}
	return "unknown_provider"
}

// ParseProvider initializes a Provider given a case-insensitive string
func ParseProvider(s string) Provider {
	switch strings.ToLower(s) {
	case "google":
		return Google
	}
	return UnknownProvider
}

// ProviderInfo contains information returned from an authorization provider
type ProviderInfo struct {
	Provider  Provider
	TokenInfo *ProviderTokenInfo
	UserInfo  *ProviderUserInfo
}

// ProviderTokenInfo contains non-user information gleaned from the
// Oauth2 provider's access token and subsequent calls to get information
// about a person using it. See ProviderUserInfo for user information.
type ProviderTokenInfo struct {

	// Token is the Oauth2 token. For inbound requests, only the
	// Access Token is given in the Authorization header, so the
	// other details (Refresh Token, Token Type, Expiry) must be
	// retrieved from a 3rd party service. The token's Expiry is
	// a calculated time of expiration (estimated). This is a moving
	// target as some providers send the actual time of expiration,
	// others just send seconds until expiration, which means it's
	// a calculation and won't have perfect precision.
	Token *oauth2.Token

	// Client ID: External ID representing the Oauth2 client which
	// authenticated the user.
	ClientID string

	// Scope: The space separated list of scopes granted to this token.
	Scope string

	// Audience: Who is the intended audience for this token. In general the
	// same as issued_to.
	Audience string `json:"audience,omitempty"`

	// IssuedTo: To whom was the token issued to. In general the same as
	// audience.
	IssuedTo string `json:"issued_to,omitempty"`
}

// ProviderUserInfo contains common fields from the various Oauth2 providers.
// Currently only using Google, so looks a lot like Google's.
type ProviderUserInfo struct {
	// ID: The obfuscated ID of the user assigned by the authentication provider.
	ExternalID string

	// Email: The user's email address.
	Email string

	// VerifiedEmail: Boolean flag which is true if the email address is
	// verified. Present only if the email scope is present in the request.
	VerifiedEmail bool

	// NamePrefix: The name prefix for the Profile (e.g. Mx., Ms., Mr., etc.)
	NamePrefix string

	// MiddleName: The person's middle name.
	MiddleName string

	// FirstName: The user's first name.
	FirstName string

	// FamilyName: The user's last name.
	LastName string

	// FullName: The user's full name.
	FullName string

	// NameSuffix: The name suffix for the person's name (e.g. "PhD", "CCNA", "OBE").
	// Other examples include generational designations like "Sr." and "Jr." and "I", "II", "III", etc.
	NameSuffix string

	// Nickname: The person's nickname
	Nickname string

	// Gender: The user's gender. TODO - setup Gender properly. not binary.
	Gender string

	// BirthDate: The full birthdate of a person (e.g. Dec 18, 1953)
	BirthDate time.Time

	// Hd: The hosted domain e.g. example.com if the user is Google apps
	// user.
	HostedDomain string

	// Link: URL of the profile page.
	ProfileLink string

	// Locale: The user's preferred locale.
	Locale string

	// Picture: URL of the user's picture image.
	Picture string
}

// Auth represents a user's authorization in the database. It captures
// the provider Oauth2 credentials. Users are linked to a Person.
// A single Person could authenticate through multiple providers.
type Auth struct {
	// ID is the unique identifier for authorization record in database
	ID uuid.UUID

	// User is the unique user associated to the authorization record.
	//
	// A Person can have one or more methods of authentication, however,
	// only one per authorization provider is allowed per User.
	User *User

	// Provider is the authentication provider
	Provider Provider

	// ProviderClientID is the external ID representing the Oauth2 client which
	// authenticated the user.
	ProviderClientID string

	// ProviderPersonID is the authentication provider's unique person/user ID.
	ProviderPersonID string

	// Provider Access Token
	ProviderAccessToken string

	// Provider Access Token Expiration Date/Time
	ProviderAccessTokenExpiry time.Time

	// Provider Refresh Token
	ProviderRefreshToken string
}

// Permission stores an approval of a mode of access to a resource.
type Permission struct {
	// ID is the unique ID for the Permission.
	ID uuid.UUID
	// ExternalID is the unique External ID to be given to outside callers.
	ExternalID secure.Identifier
	// Resource is a human-readable string which represents a resource (e.g. an HTTP route or document, etc.).
	Resource string
	// Operation represents the action taken on the resource (e.g. POST, GET, edit, etc.)
	Operation string
	// Description is what the permission is granting, e.g. "grants ability to edit a billing document".
	Description string
	// Active is a boolean denoting whether the permission is active (true) or not (false).
	Active bool
}

// Validate determines if the Permission is valid
func (p Permission) Validate() error {
	const op errs.Op = "entity/Permission.Validate"

	switch {
	case p.ID == uuid.Nil:
		return errs.E(op, errs.Validation, "ID is required")
	case p.ExternalID.String() == "":
		return errs.E(op, errs.Validation, "External ID is required")
	case p.Resource == "":
		return errs.E(op, errs.Validation, "Resource is required")
	case p.Description == "":
		return errs.E(op, errs.Validation, "Description is required")
	}
	return nil
}

// Role is a job function or title which defines an authority level.
type Role struct {
	// The unique ID for the Role.
	ID uuid.UUID
	// Unique External ID to be given to outside callers.
	ExternalID secure.Identifier
	// A human-readable code which represents the role.
	Code string
	// A longer description of the role.
	Description string
	// A boolean denoting whether the role is active (true) or not (false).
	Active bool
	// Permissions is the list of permissions allowed for the role.
	Permissions []*Permission
}

// Validate determines if the Role is valid.
func (r Role) Validate() error {
	const op errs.Op = "entity/Role.Validate"

	switch {
	case r.ID == uuid.Nil:
		return errs.E(op, errs.Validation, "ID is required")
	case r.ExternalID.String() == "":
		return errs.E(op, errs.Validation, "External ID is required")
	case r.Code == "":
		return errs.E(op, errs.Validation, "Code is required")
	case r.Description == "":
		return errs.E(op, errs.Validation, "Description is required")
	}

	return nil
}

// AuthenticationParams is the parameters needed for authenticating a User.
type AuthenticationParams struct {
	// Realm is a description of a protected area, used in the WWW-Authenticate header.
	Realm string
	// Provider is the authentication provider.
	Provider Provider
	// Token is the authentication token sent as part of Oauth2.
	Token *oauth2.Token
}
