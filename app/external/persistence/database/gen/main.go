package main

import (
	"kms/app/domain/entity"
	"kms/app/domain/repository"

	"gorm.io/gen"
)

func main() {
	// Initialize the generator with configuration
	g := gen.NewGenerator(gen.Config{
		OutPath:       "./app/external/persistence/database/repository",
		Mode:          gen.WithoutContext | gen.WithDefaultQuery | gen.WithQueryInterface,
		FieldNullable: true,
	})

	g.ApplyInterface(func(repository.IUserRepository) {}, entity.User{})

	// Generate default DAO interface for those specified structs
	g.ApplyBasic(
		entity.User{},
		entity.UserRequested{},
		entity.ChatMessage{},
		entity.ChatParticipant{},
		entity.ChatSession{},
		entity.CheckInOut{},
		entity.Class{},
		entity.Schedule{},
		entity.UserClass{},
	)

	// Execute the generator
	g.Execute()
}
