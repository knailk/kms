package entity

import (
	"kms/app/errs"
	"kms/app/secure"

	"github.com/google/uuid"
)

// OrgKind is a way of classifying an organization. Examples are Genesis, Test, Standard
type OrgKind struct {
	// ID: The unique identifier
	ID uuid.UUID
	// External ID: The unique external identifier
	ExternalID string
	// Description: A longer description of the organization kind
	Description string
}

// Validate determines whether the Person has proper data to be considered valid
func (o OrgKind) Validate() error {
	const op errs.Op = "diygoapi/OrgKind.Validate"

	switch {
	case o.ID == uuid.Nil:
		return errs.E(op, errs.Validation, "OrgKind ID cannot be nil")
	case o.ExternalID == "":
		return errs.E(op, errs.Validation, "OrgKind ExternalID cannot be empty")
	case o.Description == "":
		return errs.E(op, errs.Validation, "OrgKind Description cannot be empty")
	}

	return nil
}

// Org represents an Organization (company, institution or any other
// organized body of people with a particular purpose)
type Org struct {
	// ID: The unique identifier
	ID uuid.UUID
	// External ID: The unique external identifier
	ExternalID secure.Identifier
	// Name: The organization name
	Name string
	// Description: A longer description of the organization
	Description string
	// Kind: a way of classifying organizations
	Kind *OrgKind
}

// Validate determines whether the Org has proper data to be considered valid
func (o Org) Validate() (err error) {
	const op errs.Op = "diygoapi/Org.Validate"

	switch {
	case o.ID == uuid.Nil:
		return errs.E(op, errs.Validation, "Org ID cannot be nil")
	case o.ExternalID.String() == "":
		return errs.E(op, errs.Validation, "Org ExternalID cannot be empty")
	case o.Name == "":
		return errs.E(op, errs.Validation, "Org Name cannot be empty")
	case o.Description == "":
		return errs.E(op, errs.Validation, "Org Description cannot be empty")
	}

	if err = o.Kind.Validate(); err != nil {
		return errs.E(op, err)
	}

	return nil
}
