package dish

import (
	"kms/app/errs"

	"github.com/google/uuid"
)

type CreateDishRequest struct {
	DayOfWeek      string `json:"dayOfWeek" binding:"required"`
	Date           int    `json:"date" binding:"required"`
	Breakfast      string `json:"breakfast" binding:"required"`
	EatLightly     string `json:"eatLightly" binding:"required"`
	Lunch          string `json:"lunch" binding:"required"`
	AfternoonSnack string `json:"afternoonSnack" binding:"required"`
	Dinner         string `json:"dinner" binding:"required"`
}

func (c *CreateDishRequest) Validate() errs.Kind {
	return errs.Other
}

type CreateDishResponse struct{}

type UpdateDishRequest struct {
	DishID         uuid.UUID `json:"-"`
	Breakfast      string    `json:"breakfast"`
	EatLightly     string    `json:"eatLightly"`
	Lunch          string    `json:"lunch"`
	AfternoonSnack string    `json:"afternoonSnack"`
	Dinner         string    `json:"dinner"`
}

func (c *UpdateDishRequest) Validate() errs.Kind {
	return errs.Other
}

type UpdateDishResponse struct{}

type DeleteDishRequest struct {
	DishID uuid.UUID
}

type GetDishesForWeekRequest struct {
	FromDate int `form:"fromDate"`
	ToDate   int `form:"toDate"`
}

type GetDishesForWeekResponse struct {
	Dishes []DishResponse `json:"dishes"`
}

type DishResponse struct {
	ID             uuid.UUID `json:"id"`
	DayOfWeek      string    `json:"dayOfWeek"`
	Date           int       `json:"date"`
	Breakfast      string    `json:"breakfast"`
	EatLightly     string    `json:"eatLightly"`
	Lunch          string    `json:"lunch"`
	AfternoonSnack string    `json:"afternoonSnack"`
	Dinner         string    `json:"dinner"`
}
