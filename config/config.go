package config

import (
	"fmt"
	"github.com/ilyakaznacheev/cleanenv"
	"log"
	"time"
)

type (
	Config struct {
		App        AppConfig        `yaml:"app"`
		GRPC       GRPCConfig       `yaml:"grpc"`
		Log        LogConfig        `yaml:"logger"`
		Token      TokenConfig      `yaml:"token"`
		PG         PGConfig         `yaml:"postgres"`
		Migrations MigrationsConfig `yaml:"migrations"`
	}
	AppConfig struct {
		Name    string `yaml:"name"`
		Version string `yaml:"version"`
	}
	GRPCConfig struct {
		Port    int           `yaml:"port"`
		Timeout time.Duration `yaml:"timeout"`
	}

	LogConfig struct {
		Level string `yaml:"level"`
	}

	TokenConfig struct {
		Secret string `yaml:"secret"`
	}

	PGConfig struct {
		Port        int           `yaml:"port"`
		User        string        `yaml:"pg_user"`
		Password    string        `yaml:"pg_password"`
		Host        string        `yaml:"pg_host"`
		Name        string        `yaml:"pg_db_name"`
		MaxConns    int32         `yaml:"db_max_connections"`
		ConnTimeout time.Duration `yaml:"db_connection_timeout"`
	}

	MigrationsConfig struct {
		Path string `yaml:"path"`
	}
)

func (pc PGConfig) Url() string {
	return fmt.Sprintf("postgresql://%s:%s@%s:%d/%s?sslmode=disable",
		pc.User, pc.Password, pc.Host, pc.Port, pc.Name)
}

func (pc PGConfig) MigrationsUrl() string {
	return fmt.Sprintf("pgx5://%s:%s@%s:%d/%s?sslmode=disable",
		pc.User, pc.Password, pc.Host, pc.Port, pc.Name)
}

func NewConfig() (*Config, error) {
	cfg := &Config{}
	err := cleanenv.ReadConfig("./config/config.yaml", cfg)
	if err != nil {
		log.Fatalf("Failed to read config file: %v", err)
		return nil, err
	}
	return cfg, err
}
