package class

import (
	"kms/app/domain/entity"
	"kms/app/errs"
	"time"

	"github.com/google/uuid"
)

type CreateClassRequest struct {
	TeacherID string            `json:"teacherID"`
	DriverID  string            `json:"driverID"`
	FromDate  int64             `json:"fromDate"`
	ToDate    int64             `json:"toDate"`
	ClassName string            `json:"className"`
	AgeGroup  int               `json:"ageGroup"`
	Schedules []ScheduleRequest `json:"schedules"`
	Price     float64           `json:"price"`
	Currency  string            `json:"currency"`
}

func (c *CreateClassRequest) Validate() errs.Kind {
	if c.TeacherID == "" || c.DriverID == "" || c.FromDate == 0 || c.ToDate == 0 || c.ClassName == "" || c.AgeGroup == 0 || c.Price == 0 || c.Currency == "" {
		return errs.InvalidRequest
	}
	return errs.Other
}

type ScheduleRequest struct {
	FromTime time.Time `json:"fromTime"`
	ToTime   time.Time `json:"toTime"`
	Action   string    `json:"action"`
	Date     int64     `json:"date"`
}

type GetClassRequest struct {
	TeacherID string    `form:"-"`
	DriverID  string    `form:"-"`
	ID        uuid.UUID `form:"id"`
	FromDate  int64     `form:"fromDate"`
	ToDate    int64     `form:"toDate"`
}

func (g *GetClassRequest) Validate() errs.Kind {
	if g.ID == uuid.Nil && g.TeacherID == "" && g.DriverID == "" {
		return errs.InvalidRequest
	}
	return errs.Other
}

type ListClassesRequest struct {
	Limit    int   `form:"limit"`
	Page     int   `form:"page"`
	FromDate int64 `form:"fromDate"`
	ToDate   int64 `form:"toDate"`
	AgeGroup int   `form:"ageGroup"`
}

func (l *ListClassesRequest) Validate() errs.Kind {
	if l.Limit == 0 || l.Page == 0 {
		return errs.InvalidRequest
	}
	return errs.Other
}

type UpdateClassRequest struct {
	ID uuid.UUID `json:"-"`
}

type DeleteClassRequest struct {
	ID uuid.UUID `json:"-"`
}

func (d *DeleteClassRequest) Validate() errs.Kind {
	if d.ID == uuid.Nil {
		return errs.InvalidRequest
	}
	return errs.Other
}

type ListMembersInClassRequest struct {
	ClassID uuid.UUID `form:"-"`
}

type AddMembersToClassRequest struct {
	ClassID   uuid.UUID `json:"-"`
	Usernames []string  `json:"usernames"`
}

func (a *AddMembersToClassRequest) Validate() errs.Kind {
	if a.ClassID == uuid.Nil || len(a.Usernames) == 0 {
		return errs.InvalidRequest
	}
	return errs.Other
}

type RemoveMembersFromClassRequest struct {
	ClassID   uuid.UUID `form:"-"`
	Usernames []string  `form:"usernames"`
}

func (r *RemoveMembersFromClassRequest) Validate() errs.Kind {
	if r.ClassID == uuid.Nil || len(r.Usernames) == 0 {
		return errs.InvalidRequest
	}
	return errs.Other
}

type CheckInOutRequest struct {
	ClassID   uuid.UUID               `json:"-"`
	Usernames []string                `json:"usernames"`
	Action    entity.CheckInOutAction `json:"action"`
}

type CheckInOutHistoriesRequest struct {
	ClassID uuid.UUID `form:"-"`
	Date    int64     `form:"date"`
}

func (c *CheckInOutHistoriesRequest) Validate() errs.Kind {
	if c.ClassID == uuid.Nil || c.Date == 0 {
		return errs.InvalidRequest
	}
	return errs.Other
}
