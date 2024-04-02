package check_in_out

type CheckInOutResponse struct {
	Status string `json:"status"`
}

type ListUsersInClass struct{}

type GetUserInClass struct {
	Username string `json:"username"`
	FullName string `json:"fullName"`
}

type ListUsersInClassResponse struct {
	Users []*GetUserInClass `json:"users"`
}
