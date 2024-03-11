package cmd

import (
	"context"
	"fmt"
	"kms/app/errs"
	"kms/app/secure"
	"kms/database/sqldb"
	"kms/internal/config"
	"kms/internal/httpserver"
	"kms/internal/localtime"
	"kms/internal/shutdown"
	"kms/pkg/logger"

	"golang.org/x/text/language"
)

const (
	// log level environment variable name
	loglevelEnv string = "LOG_LEVEL"
	// minimum accepted log level environment variable name
	logLevelMinEnv string = "LOG_LEVEL_MIN"
	// log error stack environment variable name
	logErrorStackEnv string = "LOG_ERROR_STACK"
	// server port environment variable name
	portEnv string = "PORT"
	// encryption key environment variable name
	encryptKeyEnv string = "ENCRYPT_KEY"
)

// Run parses command line flags and starts the server
func Run() (err error) {
	const op errs.Op = "cmd.Run"

	cfg, err := config.Init()
	if err != nil {
		return errs.E(op, err)
	}

	// validate port in acceptable range
	err = portRange(cfg.HTTPServer.Port)
	if err != nil {
		logger.Fatal().Err(err).Msg("portRange() error")
	}

	ctx := context.Background()

	tasks := shutdown.NewShutdownTasks()
	defer tasks.ExecuteAll(ctx)

	err = localtime.Init()
	if err != nil {
		lgr.Error().Err(err).Msgf("initialize timezone")
		return
	}

	// initialize HTTP Server enfolding a http.Server with default timeouts
	// a Gorilla mux router with /api subroute and a zerolog.Logger
	s := httpserver.New(ctx, cfg, tasks)

	if cfg.EncryptionKey == "" {
		lgr.Fatal().Msg("no encryption key found")
	}

	// decode and retrieve encryption key
	var ek *[32]byte
	ek, err = secure.ParseEncryptionKey(cfg.EncryptionKey)
	if err != nil {
		lgr.Fatal().Err(err).Msg("secure.ParseEncryptionKey() error")
	}

	// initialize PostgreSQL database
	db, err := sqldb.NewPostgreSQLPool(ctx, lgr, config.NewPostgreSQLDSN(cfg))
	if err != nil {
		lgr.Fatal().Err(err).Msg("sqldb.NewPostgreSQLPool error")
	}

	var supportedLangs = []language.Tag{
		language.AmericanEnglish,
	}

	matcher := language.NewMatcher(supportedLangs)

	s.Services = httpserver.Services{
		OrgServicer: &service.OrgService{
			Datastorer:      db,
			APIKeyGenerator: secure.RandomGenerator{},
			EncryptionKey:   ek},
		AppServicer: &service.AppService{
			Datastorer:      db,
			APIKeyGenerator: secure.RandomGenerator{},
			EncryptionKey:   ek},
		PingService:   &service.PingService{Datastorer: db},
		LoggerService: &service.LoggerService{Logger: lgr},
		GenesisServicer: &service.GenesisService{
			Datastorer:      db,
			APIKeyGenerator: secure.RandomGenerator{},
			EncryptionKey:   ek,
			TokenExchanger:  gateway.Oauth2TokenExchange{},
			LanguageMatcher: matcher,
		},
		AuthenticationServicer: service.DBAuthenticationService{
			Datastorer:      db,
			TokenExchanger:  gateway.Oauth2TokenExchange{},
			EncryptionKey:   ek,
			LanguageMatcher: matcher,
		},
		AuthorizationServicer: &service.DBAuthorizationService{Datastorer: db},
		PermissionServicer:    &service.PermissionService{Datastorer: db},
		RoleServicer:          &service.RoleService{Datastorer: db},
		MovieServicer:         &service.MovieService{Datastorer: db},
	}

	return s.ListenAndServe()
}

// portRange validates the port be in an acceptable range
func portRange(port int) error {
	const op errs.Op = "cmd/portRange"

	if port < 0 || port > 65535 {
		return errs.E(op, fmt.Sprintf("port %d is not within valid port range (0 to 65535)", port))
	}
	return nil
}
