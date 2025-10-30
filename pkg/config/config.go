package config

import (
	"os"
	"strings"

	"github.com/go-playground/validator/v10"
	_ "github.com/joho/godotenv/autoload"
	"github.com/knadh/koanf/providers/env"
	"github.com/knadh/koanf/v2"
	"github.com/rs/zerolog"
)

type Config struct {
	Primary     Primary           `koanf:"primary" validate:"required"`
	Server      ServerConfig      `koanf:"server" validate:"required"`
	Database    DatabaseConfig    `koanf:"database" validate:"required"`
	Auth        AuthConfig        `koanf:"auth" validate:"required"`
	Redis       RedisConfig       `koanf:"redis" validate:"required"`
	Integration IntegrationConfig `koanf:"integration" validate:"required"`
	AWS         AWSConfig         `koanf:"aws"`
}

type Primary struct {
	Env string `koanf:"env" validate:"required"`
}

type ServerConfig struct {
	Port               string   `koanf:"port" validate:"required"`
	ReadTimeout        int      `koanf:"read_timeout" validate:"required"`
	WriteTimeout       int      `koanf:"write_timeout" validate:"required"`
	IdleTimeout        int      `koanf:"idle_timeout" validate:"required"`
	CORSAllowedOrigins []string `koanf:"cors_allowed_origins" validate:"required"`
}

type DatabaseConfig struct {
	Host            string `koanf:"host" validate:"required"`
	Port            int    `koanf:"port" validate:"required"`
	User            string `koanf:"user" validate:"required"`
	Password        string `koanf:"password"`
	Name            string `koanf:"name" validate:"required"`
	SSLMode         string `koanf:"ssl_mode" validate:"required"`
	MaxOpenConns    int    `koanf:"max_open_conns" validate:"required"`
	MaxIdleConns    int    `koanf:"max_idle_conns" validate:"required"`
	ConnMaxLifetime int    `koanf:"conn_max_lifetime" validate:"required"`
	ConnMaxIdleTime int    `koanf:"conn_max_idle_time" validate:"required"`
	DSN             string `koanf:"dsn" validate:"required"`
	// Enable running GORM AutoMigrate on startup (use with caution in production)
	AutoMigrate bool `koanf:"auto_migrate"`
	// Create PostgreSQL extensions on startup before GORM migrations
	CreateExtensions bool `koanf:"create_extensions"`
	// Seed topics and subjects data on startup (required for app functionality)
	SeedData bool `koanf:"seed_data"`
}

type AuthConfig struct {
	ClerkSecretKey     string `koanf:"clerk_secret_key" validate:"required"`
	ClerkWebHookSecret string `koanf:"clerk_webhook_secret" validate:"required"`
}

type RedisConfig struct {
	Address      string `koanf:"address" validate:"required"`
	Username     string `koanf:"username"`
	Password     string `koanf:"password"`
	DB           int    `koanf:"db"`
	UserCacheTTL int    `koanf:"user_cache_ttl" validate:"required"` // TTL in minutes for user cache
}

type IntegrationConfig struct {
	ResendAPIKey string `koanf:"resend_api_key" validate:"required"`
}

type AWSConfig struct {
	Region              string `koanf:"region" validate:"required"`
	AccessKeyID         string `koanf:"access_key_id" validate:"required"`
	SecretAccessKey     string `koanf:"secret_access_key" validate:"required"`
	UploadBucket        string `koanf:"upload_bucket" validate:"required"`
	ParsedContentBucket string `koanf:"parsed_content_bucket" validate:"required"`
	EndpointURL         string `koanf:"endpoint_url"`
}

// ...existing code...

func LoadConfig() (*Config, error) {
	logger := zerolog.New(zerolog.ConsoleWriter{Out: os.Stderr}).With().Timestamp().Logger()

	k := koanf.New(".")

	// Load environment variables with proper delimiter handling
	err := k.Load(env.Provider("", ".", func(s string) string {
		// Convert to lowercase and replace underscores with dots for nested structure
		// e.g., PRIMARY_ENV -> primary.env, SERVER_PORT -> server.port
		return strings.Replace(strings.ToLower(s), "_", ".", 1)
	}), nil)

	if err != nil {
		logger.Fatal().Err(err).Msg("could not load initial env variables")
	}

	mainConfig := &Config{}

	err = k.Unmarshal("", mainConfig)
	if err != nil {
		logger.Fatal().Err(err).Msg("could not unmarshal main config")
	}

	validate := validator.New()

	err = validate.Struct(mainConfig)
	if err != nil {
		logger.Fatal().Err(err).Msg("config validation failed")
	}

	return mainConfig, nil
}
