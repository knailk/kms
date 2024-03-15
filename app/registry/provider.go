package registry

import (
	"kms/app/config"
	"kms/app/external/infra/redis"
	"kms/app/external/infra/ristretto"
	"kms/pkg/mailer"

	"github.com/google/wire"
	"gorm.io/gorm"
)

// Dependency Injection: All repository set for wire generate
var (
	BuilderSet = wire.NewSet(
		// Repository

		// Service
		// userService.NewUserService,
		// userService.NewUserTokenService,

		// Connection
		wire.FieldsOf(new(*Provider), "Config"),
		wire.FieldsOf(new(*Provider), "DB"),
		wire.FieldsOf(new(*Provider), "RedisClient"),
		wire.FieldsOf(new(*Provider), "RistrettoClient"),
		wire.FieldsOf(new(*Provider), "MailClient"),
	)
)

// ConnectionSet is wire set related to connection

//var ConnectionSet = wire.NewSet()

// Provider is the object to bring interface in initialization process
type Provider struct {
	Config          *config.Config
	DB              *gorm.DB
	RedisClient     redis.RedisClient
	RistrettoClient ristretto.RistrettoCache
	MailClient      mailer.OTPMailer
}
