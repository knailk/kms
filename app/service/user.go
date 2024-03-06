package service

import (
	"context"
	"kms/app/entity"
	"kms/app/secure"
	"time"

	"github.com/google/uuid"
	"golang.org/x/text/language"
)

// RegisterUserServicer registers a new user
type RegisterUserServicer interface {
	SelfRegister(ctx context.Context, adt entity.Audit) error
}

// UserResponse - from Wikipedia: "A user is a person who utilizes a computer or network service." In the context of this
// project, given that we allow Persons to authenticate with multiple providers, a User is akin to a persona
// (Wikipedia - "The word persona derives from Latin, where it originally referred to a theatrical mask. On the
// social web, users develop virtual personas as online identities.") and as such, a Person can have one or many
// Users (for instance, I can have a GitHub user and a Google user, but I am just one Person).
//
// As a general, practical matter, most operations are considered at the User level. For instance, roles are
// assigned at the user level instead of the Person level, which allows for more fine-grained access control.
type UserResponse struct {
	// ID: The unique identifier for the Person's profile
	ID uuid.UUID

	// ExternalID: unique external identifier of the User
	ExternalID secure.Identifier `json:"external_id"`

	// NamePrefix: The name prefix for the Profile (e.g. Mx., Ms., Mr., etc.)
	NamePrefix string `json:"name_prefix"`

	// FirstName: The person's first name.
	FirstName string `json:"first_name"`

	// MiddleName: The person's middle name.
	MiddleName string `json:"middle_name"`

	// LastName: The person's last name.
	LastName string `json:"last_name"`

	// FullName: The person's full name.
	FullName string `json:"full_name"`

	// NameSuffix: The name suffix for the person's name (e.g. "PhD", "CCNA", "OBE").
	// Other examples include generational designations like "Sr." and "Jr." and "I", "II", "III", etc.
	NameSuffix string `json:"name_suffix"`

	// Nickname: The person's nickname
	Nickname string `json:"nickname"`

	// Email: The primary email for the User
	Email string `json:"email"`

	// CompanyName: The Company Name that the person works at
	CompanyName string `json:"company_name"`

	// CompanyDepartment: is the department at the company that the person works at
	CompanyDepartment string `json:"company_department"`

	// JobTitle: The person's Job Title
	JobTitle string `json:"job_title"`

	// BirthDate: The full birthdate of a person (e.g. Dec 18, 1953)
	BirthDate time.Time `json:"birth_date"`

	// LanguagePreferences is the user's language tag preferences.
	LanguagePreferences []language.Tag `json:"language_preferences"`

	// HostedDomain: The hosted domain e.g. example.com.
	HostedDomain string `json:"hosted_domain"`

	// PictureURL: URL of the person's picture image for the profile.
	PictureURL string `json:"picture_url"`

	// ProfileLink: URL of the profile page.
	ProfileLink string `json:"profile_link"`

	// Source: The origin of the User (e.g. Google Oauth2, Apple Oauth2, etc.)
	Source string `json:"source"`
}
