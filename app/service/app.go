package service

import (
	"context"
	"kms/app/entity"
	"kms/app/errs"
)

// AppServicer manages the retrieval and manipulation of an App
type AppServicer interface {
	Create(ctx context.Context, r *CreateAppRequest, adt entity.Audit) (*AppResponse, error)
	Update(ctx context.Context, r *UpdateAppRequest, adt entity.Audit) (*AppResponse, error)
}

// CreateAppRequest is the request struct for Creating an App
type CreateAppRequest struct {
	Name                   string `json:"name"`
	Description            string `json:"description"`
	Oauth2Provider         string `json:"oauth2_provider"`
	Oauth2ProviderClientID string `json:"oauth2_provider_client_id"`
}

// Validate determines whether the CreateAppRequest has proper data to be considered valid
func (r CreateAppRequest) Validate() error {
	const op errs.Op = "diygoapi/CreateAppRequest.Validate"

	switch {
	case r.Name == "":
		return errs.E(op, errs.Validation, "app name is required")
	case r.Description == "":
		return errs.E(op, errs.Validation, "app description is required")
	case r.Oauth2Provider != "" && r.Oauth2ProviderClientID == "":
		return errs.E(op, errs.Validation, "oAuth2 provider client ID is required when Oauth2 provider is given")
	case r.Oauth2Provider == "" && r.Oauth2ProviderClientID != "":
		return errs.E(op, errs.Validation, "oAuth2 provider is required when Oauth2 provider client ID is given")
	}
	return nil
}

// UpdateAppRequest is the request struct for Updating an App
type UpdateAppRequest struct {
	ExternalID  string
	Name        string `json:"name"`
	Description string `json:"description"`
}

// AppResponse is the response struct for an App
type AppResponse struct {
	ExternalID          string           `json:"external_id"`
	Name                string           `json:"name"`
	Description         string           `json:"description"`
	CreateAppExtlID     string           `json:"create_app_extl_id"`
	CreateUserFirstName string           `json:"create_user_first_name"`
	CreateUserLastName  string           `json:"create_user_last_name"`
	CreateDateTime      string           `json:"create_date_time"`
	UpdateAppExtlID     string           `json:"update_app_extl_id"`
	UpdateUserFirstName string           `json:"update_user_first_name"`
	UpdateUserLastName  string           `json:"update_user_last_name"`
	UpdateDateTime      string           `json:"update_date_time"`
	APIKeys             []APIKeyResponse `json:"api_keys"`
}

// APIKeyResponse is the response fields for an API key
type APIKeyResponse struct {
	Key              string `json:"key"`
	DeactivationDate string `json:"deactivation_date"`
}
