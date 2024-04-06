package class

type CreateClassRequest struct {
	TeacherID string     `json:"teacherID"`
	DriverID  string     `json:"driverID"`
	FromDate  string     `json:"fromDate"`
	ToDate    string     `json:"toDate"`
	ClassName string     `json:"className"`
	AgeGroup  int        `json:"ageGroup"`
	Schedules []Schedule `json:"schedules"`
	Price     float64    `json:"price"`
	Currency  string     `json:"currency"`
}

type Schedule struct {
	
}

type GetClassRequest struct{}

type ListClassesRequest struct{}

type UpdateClassRequest struct{}

type DeleteClassRequest struct{}
