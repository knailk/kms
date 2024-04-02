package check_in_out

import (
	"kms/app/domain/entity"

	"github.com/google/uuid"
)

type CheckInOutRequest struct {
	Usernames []string                `json:"usernames"`
	Date      int64                   `json:"date"`
	Action    entity.CheckInOutAction `json:"action"`
}

type ListUsersInClassRequest struct {
	ClassID   uuid.UUID `json:"classID"`
	TeacherID string    `json:"teacherID"`
}
