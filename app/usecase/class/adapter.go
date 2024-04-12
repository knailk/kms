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
		ID:        class.ID,
		TeacherID: class.TeacherID,
		DriverID:  class.DriverID,
		FromDate:  class.FromDate,
		ToDate:    class.ToDate,
		ClassName: class.ClassName,
		AgeGroup:  class.AgeGroup,
		Price:     class.Price,
		Currency:  class.Currency,
		Schedules: schedules,
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

func toUsersInClass(users []*entity.User) []*GetUserInClass {
	var usersInClass []*GetUserInClass
	for _, user := range users {
		usersInClass = append(usersInClass, toUserInClass(user))
	}
	return usersInClass
}

func toUserInClass(user *entity.User) *GetUserInClass {
	return &GetUserInClass{
		Username: user.Username,
		FullName: user.FullName,
	}
}

func toCheckInOutHistoriesResponse(checkInOuts []*entity.CheckInOut) []*CheckInOutHistoryResponse {
	var responses []*CheckInOutHistoryResponse
	for _, checkInOut := range checkInOuts {
		responses = append(responses, &CheckInOutHistoryResponse{
			Username: checkInOut.Username,
			Action:   checkInOut.Action.String(),
			Date:     checkInOut.Date,
		})
	}
	return responses
}
