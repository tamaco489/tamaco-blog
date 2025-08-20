package config

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
)

// ConfigLoader defines the interface for loading configuration
type ConfigLoader interface {
	Load(ctx context.Context) (*Config, error)
}

// SecretsFetcher defines the interface for fetching secrets
type SecretsFetcher interface {
	FetchSecrets(ctx context.Context, cfg aws.Config, secrets map[string]any) error
}

// SecretManager provides an interface for managing secrets
type SecretManager interface {
	GetSecret(ctx context.Context, secretID string, target any) error
	GetSecrets(ctx context.Context, secrets map[string]any) error
	RefreshSecret(ctx context.Context, secretID string, target any) error
}
