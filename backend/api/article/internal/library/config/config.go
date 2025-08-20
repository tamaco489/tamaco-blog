package config

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
)

const (
	defaultTimeout   = 5 * time.Second
	defaultAWSRegion = "ap-northeast-1"
	localstackURL    = "http://localhost.localstack.cloud:4566"
)

// Config holds all application configuration
type Config struct {
	AWSConfig aws.Config
	API       APIConfig
	CoreDB    DatabaseConfig
	mu        *sync.RWMutex // Protect concurrent access (pointer to avoid copy issues)
}

// APIConfig holds API-specific configuration
type APIConfig struct {
	Env         Environment `envconfig:"API_ENV" default:"dev"`
	Port        string      `envconfig:"API_PORT" default:"8080"`
	ServiceName string      `envconfig:"API_SERVICE_NAME" default:"article-api"`
}

// DatabaseConfig holds database connection configuration
type DatabaseConfig struct {
	Host string `json:"host" validate:"required"`
	Port string `json:"port" validate:"required"`
	User string `json:"username" validate:"required"`
	Pass string `json:"password" validate:"required"`
	Name string `json:"dbname" validate:"required"`
}

// Singleton pattern for config management
var (
	instance *Config
	once     sync.Once
	initErr  error
)

// GetInstance returns the singleton config instance
func GetInstance(ctx context.Context) (*Config, error) {
	once.Do(func() {
		loader := NewConfigLoader(nil)
		instance, initErr = loader.Load(ctx)
	})
	return instance, initErr
}

// GetDatabaseDSN returns a formatted database connection string
func (c *Config) GetDatabaseDSN() string {
	if c.mu != nil {
		c.mu.RLock()
		defer c.mu.RUnlock()
	}

	return fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=require",
		c.CoreDB.Host,
		c.CoreDB.Port,
		c.CoreDB.User,
		c.CoreDB.Pass,
		c.CoreDB.Name,
	)
}

// IsProduction returns true if the environment is production
func (c *Config) IsProduction() bool {
	if c.mu != nil {
		c.mu.RLock()
		defer c.mu.RUnlock()
	}

	return c.API.Env.IsProduction()
}

// IsDevelopment returns true if the environment is development or test
func (c *Config) IsDevelopment() bool {
	if c.mu != nil {
		c.mu.RLock()
		defer c.mu.RUnlock()
	}

	return c.API.Env.IsDevelopment()
}

// GetAPIAddr returns the API server address
func (c *Config) GetAPIAddr() string {
	if c.mu != nil {
		c.mu.RLock()
		defer c.mu.RUnlock()
	}

	return fmt.Sprintf(":%s", c.API.Port)
}
