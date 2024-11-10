package dish

import (
	"kms/app/domain/entity"
)

func toDishesResponse(dishes []*entity.Dish) []DishResponse {
	var dishResponses []DishResponse

	for _, dish := range dishes {
		dishResponses = append(dishResponses, DishResponse{
			ID:             dish.ID,
			DayOfWeek:      dish.DayOfWeek,
			Date:           dish.Date,
			Breakfast:      dish.Breakfast,
			EatLightly:     dish.EatLightly,
			Lunch:          dish.Lunch,
			AfternoonSnack: dish.AfternoonSnack,
			Dinner:         dish.Dinner,
		})
	}

	return dishResponses
}
