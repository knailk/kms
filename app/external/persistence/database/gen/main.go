package main

import (
	"kms/app/domain/entity"

	"gorm.io/gen"
)

func main() {
	// Initialize the generator with configuration
	g := gen.NewGenerator(gen.Config{
		OutPath:       "./app/external/persistence/database/repository",
		Mode:          gen.WithoutContext | gen.WithDefaultQuery | gen.WithQueryInterface,
		FieldNullable: true,
	})

	// Generate default DAO interface for those specified structs
	g.ApplyBasic(
		entity.User{},
	)

	// Execute the generator
	g.Execute()
}
