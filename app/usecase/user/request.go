package user

import (
	"kms/app/errs"
	"time"
)

type GetUserRequest struct {
	Username string `json:"-"`
}

type UpdateUserRequest struct {
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
	Longitude   *float64  `json:"longitude"`
	Latitude    *float64  `json:"latitude"`
	IsDeleted   bool      `json:"isDeleted"`
}

func (r *UpdateUserRequest) Validate() errs.Kind {
	if r.Username == "" {
		return errs.InvalidRequest
	}

	if r.Password == "" && r.FullName == "" && r.Gender == "" && r.PhoneNumber == "" && r.Email == "" && r.BirthDate.IsZero() && r.PictureURL == "" && r.Address == "" && !r.IsDeleted && r.Longitude == nil && r.Latitude == nil {
		return errs.InvalidRequest
	}

	return errs.Other
}

type SearchUserRequest struct {
	Keyword string `form:"keyword"`
}
