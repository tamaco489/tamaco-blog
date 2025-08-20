package config

import (
	"context"
	"errors"
	"fmt"
	"strconv"
	"sync"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/kelseyhightower/envconfig"
)

var (
	ErrInvalidConfig = errors.New("invalid configuration")
	ErrMissingField  = errors.New("missing required field")
)

// DefaultConfigLoader implements the ConfigLoader interface
type DefaultConfigLoader struct {
	secretsFetcher SecretsFetcher
}

// NewConfigLoader creates a new ConfigLoader instance
func NewConfigLoader(fetcher SecretsFetcher) ConfigLoader {
	if fetcher == nil {
		fetcher = &DefaultSecretsFetcher{}
	}
	return &DefaultConfigLoader{
		secretsFetcher: fetcher,
	}
}

// Load loads the configuration from environment variables and secrets
func (l *DefaultConfigLoader) Load(ctx context.Context) (*Config, error) {
	cfg := &Config{
		mu: &sync.RWMutex{},
	}

	// Load environment variables
	if err := l.loadEnvConfig(cfg); err != nil {
		return nil, fmt.Errorf("load env config: %w", err)
	}

	// Validate API configuration
	if err := l.validateAPIConfig(&cfg.API); err != nil {
		return nil, fmt.Errorf("validate API config: %w", err)
	}

	// Load AWS configuration with timeout
	timeoutCtx, cancel := context.WithTimeout(ctx, defaultTimeout)
	defer cancel()

	if err := l.loadAWSConfig(timeoutCtx, cfg); err != nil {
		return nil, fmt.Errorf("load AWS config: %w", err)
	}

	// Load secrets from AWS Secrets Manager
	if err := l.loadSecrets(timeoutCtx, cfg); err != nil {
		return nil, fmt.Errorf("load secrets: %w", err)
	}

	// Validate database configuration
	if err := l.validateDatabaseConfig(&cfg.CoreDB); err != nil {
		return nil, fmt.Errorf("validate database config: %w", err)
	}

	return cfg, nil
}

// loadEnvConfig loads configuration from environment variables
func (l *DefaultConfigLoader) loadEnvConfig(cfg *Config) error {
	if err := envconfig.Process("", &cfg.API); err != nil {
		return fmt.Errorf("process API env config: %w", err)
	}
	return nil
}

// validateAPIConfig validates the API configuration
func (l *DefaultConfigLoader) validateAPIConfig(api *APIConfig) error {
	if api.Env == "" {
		return fmt.Errorf("%w: API.Env", ErrMissingField)
	}

	// Validate port is a valid number
	port, err := strconv.Atoi(api.Port)
	if err != nil || port < 1 || port > 65535 {
		return fmt.Errorf("%w: API.Port must be a valid port number (1-65535)", ErrInvalidConfig)
	}

	// Validate environment value
	if !api.Env.IsValid() {
		return fmt.Errorf("%w: API.Env must be one of: %s, %s, %s, %s",
			ErrInvalidConfig, EnvDev, EnvTest, EnvStg, EnvPrd)
	}

	if api.ServiceName == "" {
		return fmt.Errorf("%w: API.ServiceName", ErrMissingField)
	}

	return nil
}

// loadAWSConfig loads AWS configuration
func (l *DefaultConfigLoader) loadAWSConfig(ctx context.Context, cfg *Config) error {
	awsCfg, err := config.LoadDefaultConfig(ctx, config.WithRegion(defaultAWSRegion))
	if err != nil {
		return fmt.Errorf("load default AWS config: %w", err)
	}

	// Use LocalStack for development and test environments
	if cfg.API.Env.IsDevelopment() {
		awsCfg.BaseEndpoint = aws.String(localstackURL)
	}

	cfg.AWSConfig = awsCfg
	return nil
}

// loadSecrets loads secrets from AWS Secrets Manager
func (l *DefaultConfigLoader) loadSecrets(ctx context.Context, cfg *Config) error {
	secretKey := fmt.Sprintf("%s/tamaco-blog/article/core/rds", cfg.API.Env)
	secrets := map[string]any{
		secretKey: &cfg.CoreDB,
	}

	return l.secretsFetcher.FetchSecrets(ctx, cfg.AWSConfig, secrets)
}

// validateDatabaseConfig validates the database configuration
func (l *DefaultConfigLoader) validateDatabaseConfig(db *DatabaseConfig) error {
	if db.Host == "" {
		return fmt.Errorf("%w: CoreDB.Host", ErrMissingField)
	}
	if db.Port == "" {
		return fmt.Errorf("%w: CoreDB.Port", ErrMissingField)
	}
	if db.User == "" {
		return fmt.Errorf("%w: CoreDB.User", ErrMissingField)
	}
	if db.Pass == "" {
		return fmt.Errorf("%w: CoreDB.Pass", ErrMissingField)
	}
	if db.Name == "" {
		return fmt.Errorf("%w: CoreDB.Name", ErrMissingField)
	}

	// Validate port is a valid number
	port, err := strconv.Atoi(db.Port)
	if err != nil || port < 1 || port > 65535 {
		return fmt.Errorf("%w: CoreDB.Port must be a valid port number", ErrInvalidConfig)
	}

	return nil
}

// DefaultSecretsFetcher is the default implementation of SecretsFetcher
type DefaultSecretsFetcher struct{}

// FetchSecrets fetches secrets from AWS Secrets Manager
func (f *DefaultSecretsFetcher) FetchSecrets(ctx context.Context, cfg aws.Config, secretsMap map[string]any) error {
	secretManager := NewSecretManager(cfg)
	return secretManager.GetSecrets(ctx, secretsMap)
}
