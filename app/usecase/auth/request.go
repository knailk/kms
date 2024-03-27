package auth

import (
	"kms/app/errs"
	"time"
)

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type GetInfoRequest struct {
	Username string `json:"-"`
}

type RefreshRequest struct {
	RefreshToken string `json:"refreshToken"`
}

type UpdateProfileRequest struct {
	Username    string    `json:"-"`
	OldPassword string    `json:"oldPassword"`
	Password    string    `json:"password"`
	FullName    string    `json:"fullName"`
	Gender      string    `json:"gender"`
	PhoneNumber string    `json:"phoneNumber"`
	Email       string    `json:"email"`
	BirthDate   time.Time `json:"birthDate"`
	PictureURL  string    `json:"pictureURL"`
	Address     string    `json:"address"`
	IsDeleted   bool      `json:"isDeleted"`
}

func (r *UpdateProfileRequest) Validate() errs.Kind {
	if r.Username == "" {
		return errs.InvalidRequest
	}

	if r.Password == "" && r.FullName == "" && r.Gender == "" && r.PhoneNumber == "" && r.Email == "" && r.BirthDate.IsZero() && r.PictureURL == "" && r.Address == "" && !r.IsDeleted {
		return errs.InvalidRequest
	}

	return errs.Other
}
