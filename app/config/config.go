package config

import (
	"errors"
	"fmt"
	"os"

	"github.com/jinzhu/configor"
)

const (
	EnvLocal       = "local"
	EnvDevelopment = "development"
	EnvStaging     = "staging"
	EnvProduction  = "production"
)

const (
	PlatformLocal = "local"
	PlatformAWS   = "aws"
)

// // Env defines the environment
// type Env uint8

// const (
// 	// Existing environment - current environment is not overridden
// 	Existing Env = iota
// 	// Local environment (Local machine)
// 	Local
// 	// Staging environment (GCP)
// 	Staging
// 	// Production environment (GCP)
// 	Production

// 	// Invalid defines an invalid environment option
// 	Invalid Env = 99
// )

// func (e Env) String() string {
// 	switch e {
// 	case Existing:
// 		return "existing"
// 	case Local:
// 		return "local"
// 	case Staging:
// 		return "staging"
// 	case Production:
// 		return "production"
// 	case Invalid:
// 		return "invalid"
// 	}
// 	return "unknown_env_config"
// }

// // ParseEnv converts an env string into an Env value.
// // returns Invalid if the input string does not match known values.
// func ParseEnv(envStr string) Env {
// 	switch envStr {
// 	case "existing":
// 		return Existing
// 	case "local":
// 		return Local
// 	case "staging":
// 		return Staging
// 	case "prod":
// 		return Production
// 	default:
// 		return Invalid
// 	}
// }

// Config defines the configuration file. It is the superset of
// fields for the various environments/builds.
type Config struct {
	Env                 string        `yaml:"env" env:"ENV" default:"development"`
	Platform            string        `yaml:"platform" env:"PLATFORM" default:"local"`
	App                 AppCnf        `yaml:"app"`
	Log                 LoggerCnf     `yaml:"logger"`
	DB                  DBPostgresCnf `yaml:"db"`
	EncryptionKey       string        `yaml:"encryptionKey" env:"ENCRYPTION_KEY"`
	GCP                 GCPCnf        `yaml:"gcp"`
	Cache               CacheCnf      `yaml:"cache"`
	Session             SessionCnf    `yaml:"session"`
	Email               EmailCnf      `yaml:"email"`
	WebSocket           WebSocketCnf  `yaml:"websocket"`
	WebHook             WebhookCnf    `yaml:"webhook"`
	HTTPServer          HTTPServerCnf `yaml:"httpServer"`
	BusinessAdminEmails []string      `yaml:"business_admin_emails"`
}

type AppCnf struct {
	Name             string `yaml:"name" env:"APP_NAME" default:"appname"`
	Version          string `yaml:"version" env:"APP_VERSION"`
	ProfilingEnabled bool   `yaml:"profiling_enabled" env:"APP_PROFILING_ENABLED" default:"false"`
}

type LoggerCnf struct {
	LogLevel      string `yaml:"logLevel" env:"LOG_LEVEL"`
	MinLogLevel   string `yaml:"minLogLevel" env:"MIN_LOG_LEVEL"`
	LogErrorStack bool   `yaml:"logErrorStack" env:"LOG_ERROR_STACK"`
}

type GCPCnf struct {
	ProjectID        string              `yaml:"projectID" env:"GCP_PROJECT_ID"`
	ArtifactRegistry ArtifactRegistryCnf `yaml:"artifactRegistry"`
	CloudSQL         CloudSQLCnf         `yaml:"cloudSQL"`
	CloudRun         CloudRunCnf         `yaml:"cloudRun"`
}

type ArtifactRegistryCnf struct {
	RepoLocation string `yaml:"repoLocation" env:"GCP_ARTIFACT_REGISTRY_REPO_LOCATION"`
	RepoName     string `yaml:"repoName" env:"GCP_ARTIFACT_REGISTRY_REPO_NAME"`
	ImageID      string `yaml:"imageID" env:"GCP_ARTIFACT_REGISTRY_IMAGE_ID"`
	Tag          string `yaml:"tag" env:"GCP_ARTIFACT_REGISTRY_TAG"`
}

type CloudSQLCnf struct {
	InstanceName           string `yaml:"instanceName" env:"GCP_CLOUD_SQL_INSTANCE_NAME"`
	InstanceConnectionName string `yaml:"instanceConnectionName" env:"GCP_CLOUD_SQL_INSTANCE_CONNECTION_NAME"`
}

type CloudRunCnf struct {
	ServiceName string `yaml:"serviceName" env:"GCP_CLOUD_RUN_SERVICE_NAME"`
}

type DBPostgresCnf struct {
	Host       string `yaml:"host" env:"DB_PG_HOST"`
	User       string `yaml:"user" env:"DB_PG_USER"`
	Password   string `yaml:"password" env:"DB_PG_PASSWORD"`
	DBName     string `yaml:"dbName" env:"DB_PG_NAME"`
	Port       int    `yaml:"port" env:"DB_PG_PORT"`
	SearchPath string `yaml:"searchPath" env:"DB_PG_SEARCH_PATH"`
}

type CacheCnf struct {
	URL          string `yaml:"url" env:"CACHE_URL"`
	PoolSize     int    `yaml:"pool_size" env:"CACHE_POOL_SIZE" default:"10"`
	IdleTimeout  int    `yaml:"idle_timeout" env:"CACHE_IDLE_TIMEOUT"`
	ReadTimeout  int    `yaml:"read_timeout" env:"CACHE_READ_TIMEOUT"`
	WriteTimeout int    `yaml:"write_timeout" env:"CACHE_WRITE_TIMEOUT"`
	MinIdleConns int    `yaml:"min_idle_conns" env:"CACHE_MIN_IDLE_CONNS"`
	UseTLS       bool   `yaml:"use_tls" env:"CACHE_USE_TLS" default:"false"`
}

type SessionCnf struct {
	BasicAuth BasicAuthCnf `yaml:"basicAuth"`
	JWT       JWTCnf       `yaml:"jwt"`
}

type EmailCnf struct {
	Contact   string   `yaml:"contact"`
	NoReply   string   `yaml:"no_reply"`
	Ceo       string   `yaml:"ceo"`
	EAPAdmins []string `yaml:"eap_admins"`
}

type WebSocketCnf struct {
	URL string `yaml:"url" env:"WEBSOCKET_URL"`
}

type BasicAuthCnf struct {
	UserName string `yaml:"userName" env:"BASIC_AUTH_USERNAME"`
	Password string `yaml:"password" env:"BASIC_AUTH_PASSWORD"`
}

type JWTCnf struct {
	Secret          string `yaml:"secret" env:"SESSION_JWT_SECRET" default:"dummy_for_local"`
	AccessTokenExp  int64  `yaml:"accessTokenExp" env:"SESSION_JWT_ACCESS_TOKEN_EXP" default:"14400"`
	RefreshTokenExp int64  `yaml:"refreshTokenExp" env:"SESSION_JWT_REFRESH_TOKEN_EXP" default:"86400"`
}

type HTTPServerCnf struct {
	Port int     `yaml:"port" env:"HTTP_SERVER_PORT" default:"20000"`
	CORS CorsCnf `yaml:"cors"`
}

type CorsCnf struct {
	AllowOrigins []string `yaml:"allowOrigins" env:"HTTP_SERVER_CORS_ALLOW_ORIGINS"`
}

type WebhookCnf struct {
	Stripe StripeCnf `yaml:"stripe"`
	Slack  SlackCnf  `yaml:"slack"`
}

type SlackCnf struct {
	TransactionNotificationWebhookURL string `yaml:"transaction_notification_slack_url" env:"TRANSACTION_PAYMENT_NOTIFICATION_SLACK_URL"`
	PaymentNotificationWebhookURL     string `yaml:"payment_notification_slack_url" env:"SLACK_PAYMENT_NOTIFICATION_SLACK_URL"`
	CustomerSignUpWebhookURL          string `yaml:"customer_sign_up_slack_url" env:"SLACK_CUSTOMER_SIGN_UP_SLACK_URL"`
	TherapistSignUpWebhookURL         string `yaml:"therapist_sign_up_slack_url" env:"SLACK_THERAPIST_SIGN_UP_SLACK_URL"`
	FinishTestWebhookURL              string `yaml:"finish_test_slack_url" env:"SLACK_FINISH_TEST_SLACK_URL"`
	UpdateTherapistScheduleURL        string `yaml:"update_therapist_schedule_url" env:"SLACK_UPDATE_THERAPIST_SCHEDULE_URL"`
}

type StripeCnf struct {
	StripeEndpointKey string `yaml:"stripe_endpoint_key" env:"STRIPE_ENDPOINT_KEY"`
	StripeSecretKey   string `yaml:"stripe_secret_key" env:"STRIPE_SECRET_KEY"`
}

// loadConfigAs to load config of specified struct
func loadConfigAs(config interface{}, path string) error {
	// configor doesn't check file existence
	if _, err := os.Stat(path); errors.Is(err, os.ErrNotExist) {
		return fmt.Errorf("config file not exist %s: %w", path, err)
	}

	err := configor.Load(config, path)
	if err != nil {
		return fmt.Errorf("failed to load config file %s: %w", path, err)
	}
	return nil
}

// LoadConfig to load config of type Config
func LoadConfig(path string) (*Config, error) {
	config := Config{}
	err := loadConfigAs(&config, path)
	if err != nil {
		return nil, err
	}
	return &config, nil
}
