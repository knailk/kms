package class

import "kms/app/domain/entity"

func toListClassesResponse(classes []*entity.Class) *ListClassesResponse {
	response := make([]*GetClassResponse, len(classes))

	for i, c := range classes {
		response[i] = toGetClassResponse(c)
	}

	return &ListClassesResponse{Classes: response}
}

func toGetClassResponse(class *entity.Class) *GetClassResponse {
	schedules := make([]ScheduleResponse, 0)

	for _, s := range class.Schedules {
		schedules = append(schedules, toScheduleResponse(s))
	}

	return &GetClassResponse{
		ID:          class.ID,
		TeacherID:   class.TeacherID,
		DriverID:    class.DriverID,
		FromDate:    class.FromDate,
		ToDate:      class.ToDate,
		Status:      class.Status,
		Description: class.Description,
		ClassName:   class.ClassName,
		AgeGroup:    class.AgeGroup,
		Price:       class.Price,
		Currency:    class.Currency,
		Schedules:   schedules,
	}
}

func toScheduleResponse(s entity.Schedule) ScheduleResponse {
	return ScheduleResponse{
		ID:       s.ID,
		ClassID:  s.ClassID,
		FromTime: s.FromTime,
		ToTime:   s.ToTime,
		Date:     s.Date,
		Action:   s.Action,
	}
}

func toCheckInOutHistoriesResponse(userClasses []*entity.UserClass) []*CheckInOutHistoryResponse {
	var responses []*CheckInOutHistoryResponse
	for _, userClass := range userClasses {
		responses = append(responses, &CheckInOutHistoryResponse{
			Username:    userClass.Username,
			FullName:    userClass.User.FullName,
			PictureURL:  userClass.User.PictureURL,
			UserClassID: userClass.ID,
			Status:      userClass.Status,
			CheckInOuts: toCheckInOuts(userClass.CheckInOuts),
		})
	}
	return responses
}

func toCheckInOuts(checkInOuts []*entity.CheckInOut) []*CheckInOut {
	var responses []*CheckInOut
	for _, checkInOut := range checkInOuts {
		responses = append(responses, &CheckInOut{
			ID:     checkInOut.ID,
			Action: checkInOut.Action,
			Date:   checkInOut.Date,
		})
	}

	return responses
}
