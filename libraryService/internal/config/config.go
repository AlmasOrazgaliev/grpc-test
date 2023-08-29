package config

import (
	"os"
	"path/filepath"
	"time"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

const (
	defaultPort               = "80"
	defaultReadTimeout        = 15 * time.Second
	defaultWriteTimeout       = 15 * time.Second
	defaultIdleTimeout        = 60 * time.Second
	defaultMaxHeaderMegabytes = 1
)

type (
	Config struct {
		GRPC  GRPCConfig
		MONGO DatabaseConfig
	}

	GRPCConfig struct {
		Port               string
		Host               string
		ReadTimeout        time.Duration
		WriteTimeout       time.Duration
		IdleTimeout        time.Duration
		MaxHeaderMegabytes int
		Schema             string
	}

	ClientConfig struct {
		Endpoint string
		Username string
		Password string
	}

	DatabaseConfig struct {
		DB string
	}
)

// New populates Config struct with values from config file
// located at filepath and environment variables.
func New() (cfg Config, err error) {
	root, err := os.Getwd()
	if err != nil {
		return
	}

	grpcConfig := GRPCConfig{
		Port:               defaultPort,
		ReadTimeout:        defaultReadTimeout,
		WriteTimeout:       defaultWriteTimeout,
		IdleTimeout:        defaultIdleTimeout,
		MaxHeaderMegabytes: defaultMaxHeaderMegabytes,
	}
	cfg.GRPC = grpcConfig

	godotenv.Load(filepath.Join(root, ".env"))

	err = envconfig.Process("GRPC", &cfg.GRPC)
	if err != nil {
		return
	}

	err = envconfig.Process("MONGO", &cfg.MONGO)
	if err != nil {
		return
	}

	return
}
