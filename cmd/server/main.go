package main

import (
	"context"
	"fmt"
	"kms/app/errs"
	"kms/app/registry"
	"kms/database/sqldb"
	"kms/internal/config"
	"kms/internal/cron"
	"kms/internal/httpserver"
	"kms/internal/jwt"
	"kms/internal/localtime"
	"kms/internal/shutdown"
	"kms/pkg/logger"
	"kms/pkg/mailer"
	"os"

	"kms/internal/cache"
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

func main() {
	if err := Run(); err != nil {
		fmt.Fprintf(os.Stderr, "error from commands.Run(): %s\n", err)
		os.Exit(1)
	}
}

// Run parses command line flags and starts the server
func Run() (err error) {
	const op errs.Op = "cmd.Run"

	cfg, err := config.Init()
	if err != nil {
		return errs.E(op, err)
	}

	provider := &registry.Provider{Config: cfg}

	// validate port in acceptable range
	err = portRange(cfg.HTTPServer.Port)
	if err != nil {
		logger.Fatal("portRange() error: ", err)
	}

	ctx := context.Background()

	tasks := shutdown.NewShutdownTasks()
	defer tasks.ExecuteAll(ctx)

	err = localtime.Init()
	if err != nil {
		logger.Error("initialize timezone error: ", err)
		return
	}

	// init jwt service
	err = jwt.InitJWTService(ctx, tasks, cfg)
	if err != nil {
		return errs.E(op, err)
	}

	// // cache
	// provider.RedisClient, err = cache.InitRedis(ctx, tasks, cfg)
	// if err != nil {
	// 	return errs.E(op, err)
	// }
	provider.RistrettoClient, err = cache.InitRistretto(ctx, tasks)
	if err != nil {
		return errs.E(op, err)
	}

	// mailer
	provider.MailClient = mailer.NewOTPMailer("config.AppConfig.OTPEmail", "config.AppConfig.OTPPassword")

	// initialize PostgreSQL database
	provider.DB, err = sqldb.DBInit(ctx, config.NewPostgreSQLDSN(cfg), tasks)
	if err != nil {
		logger.Fatal("sqldb.DBInit error: ", err)
	}

	if err = cron.StartCron(ctx, provider, tasks); err != nil {
		logger.WithError(err).Error(err, "start cron job error")
		return
	}

	// initialize HTTP Server enfolding a http.Server with default timeouts, a Gin router with /api subroute
	httpserver.Init(ctx, cfg, tasks, provider)

	tasks.WaitForServerStop(ctx)

	return nil
}

// portRange validates the port be in an acceptable range
func portRange(port int) error {
	const op errs.Op = "cmd/portRange"

	if port < 0 || port > 65535 {
		return errs.E(op, fmt.Sprintf("port %d is not within valid port range (0 to 65535)", port))
	}
	return nil
}
