package config

import (
	"context"
	"encoding/json"
	"fmt"
	"maps"
	"slices"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/secretsmanager"
)

// SecretError represents an error when fetching secrets
type SecretError struct {
	SecretName string
	Message    string
	Err        error
}

func (e *SecretError) Error() string {
	if e.Err != nil {
		return fmt.Sprintf("secret %q: %s: %v", e.SecretName, e.Message, e.Err)
	}
	return fmt.Sprintf("secret %q: %s", e.SecretName, e.Message)
}

func (e *SecretError) Unwrap() error {
	return e.Err
}

// batchGetSecretsImproved fetches secrets with improved error handling
func batchGetSecretsImproved(ctx context.Context, cfg aws.Config, secrets map[string]any) error {
	if len(secrets) == 0 {
		return nil
	}

	svc := secretsmanager.NewFromConfig(cfg)

	// Validate input
	secretIDs := slices.Collect(maps.Keys(secrets))
	if len(secretIDs) == 0 {
		return fmt.Errorf("no secret IDs provided")
	}

	// Create paginator
	paginator := secretsmanager.NewBatchGetSecretValuePaginator(
		svc,
		&secretsmanager.BatchGetSecretValueInput{
			SecretIdList: secretIDs,
		},
	)

	processedSecrets := make(map[string]bool)

	// Process all pages
	for paginator.HasMorePages() {
		select {
		case <-ctx.Done():
			return fmt.Errorf("context cancelled while fetching secrets: %w", ctx.Err())
		default:
		}

		output, err := paginator.NextPage(ctx)
		if err != nil {
			return fmt.Errorf("batch get secrets: %w", err)
		}

		// Handle API errors
		if len(output.Errors) > 0 {
			var errMsgs []string
			for _, apiErr := range output.Errors {
				if apiErr.SecretId != nil && apiErr.Message != nil {
					errMsgs = append(errMsgs, fmt.Sprintf("%s: %s", *apiErr.SecretId, *apiErr.Message))
				}
			}
			return fmt.Errorf("failed to fetch secrets: %v", errMsgs)
		}

		// Process secret values
		for _, secretValue := range output.SecretValues {
			if secretValue.Name == nil || secretValue.SecretString == nil {
				continue
			}

			secretName := *secretValue.Name
			target, exists := secrets[secretName]
			if !exists {
				continue
			}

			// Mark as processed
			processedSecrets[secretName] = true

			// Unmarshal based on target type
			if err := unmarshalSecret(secretName, *secretValue.SecretString, target); err != nil {
				return &SecretError{
					SecretName: secretName,
					Message:    "failed to unmarshal",
					Err:        err,
				}
			}
		}
	}

	// Check if all secrets were processed
	for secretID := range secrets {
		if !processedSecrets[secretID] {
			return &SecretError{
				SecretName: secretID,
				Message:    "secret not found in response",
			}
		}
	}

	return nil
}

// unmarshalSecret unmarshals a secret value into the target based on its type
func unmarshalSecret(name string, secretString string, target any) error {
	if target == nil {
		return fmt.Errorf("target is nil")
	}

	switch t := target.(type) {
	case *string:
		if t == nil {
			return fmt.Errorf("target string pointer is nil")
		}
		*t = secretString
		return nil

	case *[]byte:
		if t == nil {
			return fmt.Errorf("target byte slice pointer is nil")
		}
		*t = []byte(secretString)
		return nil

	default:
		// For other types, unmarshal as JSON
		if err := json.Unmarshal([]byte(secretString), target); err != nil {
			return fmt.Errorf("unmarshal JSON: %w", err)
		}
		return nil
	}
}

// DefaultSecretManager implements SecretManager
type DefaultSecretManager struct {
	awsConfig aws.Config
}

// NewSecretManager creates a new SecretManager
func NewSecretManager(cfg aws.Config) SecretManager {
	return &DefaultSecretManager{
		awsConfig: cfg,
	}
}

// GetSecret fetches a single secret
func (m *DefaultSecretManager) GetSecret(ctx context.Context, secretID string, target any) error {
	secrets := map[string]any{
		secretID: target,
	}
	return m.GetSecrets(ctx, secrets)
}

// GetSecrets fetches multiple secrets
func (m *DefaultSecretManager) GetSecrets(ctx context.Context, secrets map[string]any) error {
	return batchGetSecretsImproved(ctx, m.awsConfig, secrets)
}

// RefreshSecret refreshes a single secret value
func (m *DefaultSecretManager) RefreshSecret(ctx context.Context, secretID string, target any) error {
	// Force refresh by bypassing any potential cache
	svc := secretsmanager.NewFromConfig(m.awsConfig)

	output, err := svc.GetSecretValue(ctx, &secretsmanager.GetSecretValueInput{
		SecretId:     aws.String(secretID),
		VersionStage: aws.String("AWSCURRENT"),
	})
	if err != nil {
		return fmt.Errorf("get secret value %q: %w", secretID, err)
	}

	if output.SecretString == nil {
		return fmt.Errorf("secret %q has no string value", secretID)
	}

	return unmarshalSecret(secretID, *output.SecretString, target)
}
