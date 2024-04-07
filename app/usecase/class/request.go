package class

import (
	"kms/app/errs"
	"time"

	"github.com/google/uuid"
)

type CreateClassRequest struct {
	TeacherID string     `json:"teacherID"`
	DriverID  string     `json:"driverID"`
	FromDate  int64      `json:"fromDate"`
	ToDate    int64      `json:"toDate"`
	ClassName string     `json:"className"`
	AgeGroup  int        `json:"ageGroup"`
	Schedules []Schedule `json:"schedules"`
	Price     float64    `json:"price"`
	Currency  string     `json:"currency"`
}

func (c *CreateClassRequest) Validate() errs.Kind {
	if c.TeacherID == "" || c.DriverID == "" || c.FromDate == 0 || c.ToDate == 0 || c.ClassName == "" || c.AgeGroup == 0 || c.Price == 0 || c.Currency == "" {
		return errs.InvalidRequest
	}
	return errs.Other
}

type Schedule struct {
	FromTime time.Time `json:"fromTime"`
	ToTime   time.Time `json:"toTime"`
	Action   string    `json:"action"`
	Date     int64     `json:"date"`
}

type GetClassRequest struct {
	ID        uuid.UUID `form:"id"`
	ClassName string    `form:"className"`
}

func (g *GetClassRequest) Validate() errs.Kind {
	if g.ID == uuid.Nil && g.ClassName == "" {
		return errs.InvalidRequest
	}
	return errs.Other
}

type ListClassesRequest struct{}

type UpdateClassRequest struct{}

type DeleteClassRequest struct{}
