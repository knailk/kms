package class

import (
	"kms/app/domain/entity"
	"kms/app/errs"
	"time"

	"github.com/google/uuid"
)

type CreateClassRequest struct {
	TeacherID   string            `json:"teacherID"`
	DriverID    string            `json:"driverID"`
	FromDate    int64             `json:"fromDate"`
	ToDate      int64             `json:"toDate"`
	ClassName   string            `json:"className"`
	AgeGroup    int               `json:"ageGroup"`
	Schedules   []ScheduleRequest `json:"schedules"`
	Price       float64           `json:"price"`
	Currency    string            `json:"currency"`
	Description string            `json:"description"`

	UserRequested string `json:"-"`
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
	StudentID string    `form:"-"`
	ID        uuid.UUID `form:"-"`
	FromDate  int64     `form:"fromDate"`
	ToDate    int64     `form:"toDate"`
}

func (g *GetClassRequest) Validate() errs.Kind {
	if g.ID == uuid.Nil && g.TeacherID == "" && g.DriverID == "" && g.StudentID == "" {
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

func (l *ListMembersInClassRequest) Validate() errs.Kind {
	if l.ClassID == uuid.Nil {
		return errs.InvalidRequest
	}
	return errs.Other
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
	ClassID   uuid.UUID `json:"-"`
	Usernames []string  `json:"usernames"`
}

func (r *RemoveMembersFromClassRequest) Validate() errs.Kind {
	if r.ClassID == uuid.Nil || len(r.Usernames) == 0 {
		return errs.InvalidRequest
	}
	return errs.Other
}

type CheckInOutRequest struct {
	Action      entity.CheckInOutAction `json:"action"`
	CheckInOuts []struct {
		UserClassID uuid.UUID `json:"userClassID"`
		Date        int64     `json:"date"`
	} `json:"checkInOuts"`
}

type CheckInOutHistoriesRequest struct {
	ClassID  uuid.UUID `form:"-"`
	FromDate int64     `form:"fromDate"`
	ToDate   int64     `form:"toDate"`
}

func (c *CheckInOutHistoriesRequest) Validate() errs.Kind {
	if c.ClassID == uuid.Nil || (c.FromDate == 0 && c.ToDate == 0) {
		return errs.InvalidRequest
	}
	return errs.Other
}

type CreateScheduleRequest struct {
	ClassID uuid.UUID `json:"classID"`

	Action   string    `json:"action"`
	FromTime time.Time `json:"fromTime"`
	ToTime   time.Time `json:"toTime"`
	Date     int64     `json:"date"`
}

func (c *CreateScheduleRequest) Validate() errs.Kind {
	if c.ClassID == uuid.Nil || c.Action == "" || c.FromTime.IsZero() || c.ToTime.IsZero() || c.Date == 0 {
		return errs.InvalidRequest
	}
	return errs.Other
}

type UpdateScheduleRequest struct {
	ID uuid.UUID `json:"-"`

	Action string `json:"action"`
}

func (u *UpdateScheduleRequest) Validate() errs.Kind {
	if u.ID == uuid.Nil || u.Action == "" {
		return errs.InvalidRequest
	}
	return errs.Other
}

type DeleteScheduleRequest struct {
	ID uuid.UUID `json:"-"`
}

func (d *DeleteScheduleRequest) Validate() errs.Kind {
	if d.ID == uuid.Nil {
		return errs.InvalidRequest
	}
	return errs.Other
}
