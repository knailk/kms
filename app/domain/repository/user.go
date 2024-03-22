package repository

import "gorm.io/gen"

type IUserRepository interface {
	// SELECT * FROM @@table WHERE username=@username
	GetByUsername(username string) (*gen.T, error)
}
