package entity

import (
	"kms/app/errs"
	"time"

	"github.com/google/uuid"
)

// User - from Wikipedia: "A user is a person who utilizes a computer or network service." In the context of this
// project, given that we allow Persons to authenticate with multiple providers, a User is akin to a persona
// (Wikipedia - "The word persona derives from Latin, where it originally referred to a theatrical mask. On the
// social web, users develop virtual personas as online identities.") and as such, a Person can have one or many
// Users (for instance, I can have a GitHub user and a Google user, but I am just one Person).
//
// As a general, practical matter, most operations are considered at the User level. For instance, roles are
// assigned at the user level instead of the Person level, which allows for more fine-grained access control.
type User struct {
	// ID: The unique identifier for the Person's profile
	ID uuid.UUID

	// NamePrefix: The name prefix for the Profile (e.g. Mx., Ms., Mr., etc.)
	NamePrefix string

	// FirstName: The person's first name.
	FirstName string

	// MiddleName: The person's middle name.
	MiddleName string

	// LastName: The person's last name.
	LastName string

	// FullName: The person's full name.
	FullName string

	// NameSuffix: The name suffix for the person's name (e.g. "PhD", "CCNA", "OBE").
	// Other examples include generational designations like "Sr." and "Jr." and "I", "II", "III", etc.
	NameSuffix string

	// Nickname: The person's nickname
	Nickname string

	// Gender: The user's gender. TODO - setup Gender properly. not binary.
	Gender string

	// Email: The primary email for the User
	Email string

	// BirthDate: The full birthdate of a person (e.g. Dec 18, 1953)
	BirthDate time.Time

	// PictureURL: URL of the person's picture image for the profile.
	PictureURL string
}

// Validate determines whether the Person has proper data to be considered valid
func (u User) Validate() error {
	const op errs.Op = "entity/User.Validate"

	switch {
	case u.ID == uuid.Nil:
		return errs.E(op, errs.Validation, "User ID cannot be nil")
	case u.LastName == "":
		return errs.E(op, errs.Validation, "User LastName cannot be empty")
	case u.FirstName == "":
		return errs.E(op, errs.Validation, "User FirstName cannot be empty")
	}

	return nil
}

// NullUUID returns ID as uuid.NullUUID
func (u User) NullUUID() uuid.NullUUID {
	if u.ID == uuid.Nil {
		return uuid.NullUUID{}
	}
	return uuid.NullUUID{
		UUID:  u.ID,
		Valid: true,
	}
}
