package class

import (
	"time"

	"github.com/google/uuid"
)

type CreateClassResponse struct{}

type GetClassResponse struct {
	ID        uuid.UUID          `json:"id"`
	TeacherID string             `json:"teacherID"`
	DriverID  string             `json:"driverID"`
	FromDate  int64              `json:"fromDate"`
	ToDate    int64              `json:"toDate"`
	ClassName string             `json:"className"`
	AgeGroup  int                `json:"ageGroup"`
	Price     float64            `json:"price"`
	Currency  string             `json:"currency"`
	Schedules []ScheduleResponse `json:"schedules"`
}

type ListClassesResponse struct {
	Classes []*GetClassResponse `json:"classes"`
}

type UpdateClassResponse struct{}

type DeleteClassResponse struct{}

type ScheduleResponse struct {
	ID       uuid.UUID `json:"id"`
	ClassID  uuid.UUID `json:"classID"`
	FromTime time.Time `json:"fromTime"`
	ToTime   time.Time `json:"toTime"`
	Action   string    `json:"action"`
	Date     int64     `json:"date"`
}
