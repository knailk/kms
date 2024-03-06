package service

import (
	"context"
	"kms/app/entity"
	"kms/app/errs"
)

// OrgServicer manages the retrieval and manipulation of an Org
type OrgServicer interface {
	// Create manages the creation of an Org (and optional app)
	Create(ctx context.Context, r *CreateOrgRequest, adt entity.Audit) (*OrgResponse, error)
	Update(ctx context.Context, r *UpdateOrgRequest, adt entity.Audit) (*OrgResponse, error)
	Delete(ctx context.Context, extlID string) (DeleteResponse, error)
	FindAll(ctx context.Context) ([]*OrgResponse, error)
	FindByExternalID(ctx context.Context, extlID string) (*OrgResponse, error)
}

// CreateOrgRequest is the request struct for Creating an Org
type CreateOrgRequest struct {
	Name             string            `json:"name"`
	Description      string            `json:"description"`
	Kind             string            `json:"kind"`
	CreateAppRequest *CreateAppRequest `json:"app"`
}

// Validate determines whether the CreateOrgRequest has proper data to be considered valid
func (r CreateOrgRequest) Validate() error {
	const op errs.Op = "diygoapi/CreateOrgRequest.Validate"

	switch {
	case r.Name == "":
		return errs.E(op, errs.Validation, "org name is required")
	case r.Description == "":
		return errs.E(op, errs.Validation, "org description is required")
	case r.Kind == "":
		return errs.E(op, errs.Validation, "org kind is required")
	}
	return nil
}

// UpdateOrgRequest is the request struct for Updating an Org
type UpdateOrgRequest struct {
	ExternalID  string
	Name        string `json:"name"`
	Description string `json:"description"`
}

// OrgResponse is the response struct for an Org.
// It contains only one app (even though an org can have many apps).
// This app is only present in the response when creating an org and
// accompanying app. I may change this later to be different response
// structs for different purposes, but for now, this works.
type OrgResponse struct {
	ExternalID          string       `json:"external_id"`
	Name                string       `json:"name"`
	KindExternalID      string       `json:"kind_description"`
	Description         string       `json:"description"`
	CreateAppExtlID     string       `json:"create_app_extl_id"`
	CreateUserFirstName string       `json:"create_user_first_name"`
	CreateUserLastName  string       `json:"create_user_last_name"`
	CreateDateTime      string       `json:"create_date_time"`
	UpdateAppExtlID     string       `json:"update_app_extl_id"`
	UpdateUserFirstName string       `json:"update_user_first_name"`
	UpdateUserLastName  string       `json:"update_user_last_name"`
	UpdateDateTime      string       `json:"update_date_time"`
	App                 *AppResponse `json:"app,omitempty"`
}
