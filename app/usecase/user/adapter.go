package user

import "kms/app/domain/entity"

func toUsers(users []entity.User) []*GetUserResponse {
	var res []*GetUserResponse
	for _, user := range users {
		res = append(res, toUser(user))
	}
	return res
}

func toUser(user entity.User) GetUserResponse {
	return GetUserResponse{
		Username:    user.Username,
		FullName:    user.FullName,
		Email:       user.Email,
		Role:        string(user.Role),
		Gender:      user.Gender,
		BirthDate:   user.BirthDate,
		PhoneNumber: user.PhoneNumber,
		PictureURL:  user.PictureURL,
		Address:     user.Address,
	}
}
