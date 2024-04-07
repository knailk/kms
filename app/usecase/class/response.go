package class

import "github.com/google/uuid"

type CreateClassResponse struct{}

type GetClassResponse struct {
	ID        uuid.UUID `json:"id"`
	TeacherID string    `json:"teacherID"`
	DriverID  string    `json:"driverID"`
	FromDate  int64     `json:"fromDate"`
	ToDate    int64     `json:"toDate"`
	ClassName string    `json:"className"`
	AgeGroup  int       `json:"ageGroup"`
	Price     float64   `json:"price"`
	Currency  string    `json:"currency"`
}

type ListClassesResponse struct{}

type UpdateClassResponse struct{}

type DeleteClassResponse struct{}
