package check_in_out

import "kms/app/domain/entity"

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
