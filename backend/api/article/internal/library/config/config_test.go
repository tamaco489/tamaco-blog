package config

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
)

// MockSecretsFetcher is a mock implementation of SecretsFetcher for testing
type MockSecretsFetcher struct {
	FetchFunc func(ctx context.Context, cfg aws.Config, secrets map[string]any) error
}

func (m *MockSecretsFetcher) FetchSecrets(ctx context.Context, cfg aws.Config, secrets map[string]any) error {
	if m.FetchFunc != nil {
		return m.FetchFunc(ctx, cfg, secrets)
	}
	// Default mock implementation
	for key := range secrets {
		if db, ok := secrets[key].(*DatabaseConfig); ok {
			db.Host = "test-host"
			db.Port = "5432"
			db.User = "test-user"
			db.Pass = "test-pass"
			db.Name = "test-db"
		}
	}
	return nil
}

func TestConfigLoader_Load(t *testing.T) {
	tests := []struct {
		name           string
		setupEnv       func()
		mockFetcher    *MockSecretsFetcher
		wantErr        bool
		expectedErrMsg string
	}{
		{
			name: "successful load with dev environment",
			setupEnv: func() {
				t.Setenv("API_ENV", "dev")
				t.Setenv("API_PORT", "8080")
				t.Setenv("API_SERVICE_NAME", "test-service")
			},
			mockFetcher: &MockSecretsFetcher{},
			wantErr:     false,
		},
		{
			name: "invalid port number",
			setupEnv: func() {
				t.Setenv("API_ENV", "dev")
				t.Setenv("API_PORT", "99999")
				t.Setenv("API_SERVICE_NAME", "test-service")
			},
			mockFetcher:    &MockSecretsFetcher{},
			wantErr:        true,
			expectedErrMsg: "validate API config",
		},
		{
			name: "invalid environment",
			setupEnv: func() {
				t.Setenv("API_ENV", "invalid")
				t.Setenv("API_PORT", "8080")
				t.Setenv("API_SERVICE_NAME", "test-service")
			},
			mockFetcher:    &MockSecretsFetcher{},
			wantErr:        true,
			expectedErrMsg: "validate API config",
		},
		{
			name: "secrets fetch error",
			setupEnv: func() {
				t.Setenv("API_ENV", "dev")
				t.Setenv("API_PORT", "8080")
				t.Setenv("API_SERVICE_NAME", "test-service")
			},
			mockFetcher: &MockSecretsFetcher{
				FetchFunc: func(ctx context.Context, cfg aws.Config, secrets map[string]any) error {
					return errors.New("fetch error")
				},
			},
			wantErr:        true,
			expectedErrMsg: "load secrets",
		},
		{
			name: "missing database host",
			setupEnv: func() {
				t.Setenv("API_ENV", "dev")
				t.Setenv("API_PORT", "8080")
				t.Setenv("API_SERVICE_NAME", "test-service")
			},
			mockFetcher: &MockSecretsFetcher{
				FetchFunc: func(ctx context.Context, cfg aws.Config, secrets map[string]any) error {
					for key := range secrets {
						if db, ok := secrets[key].(*DatabaseConfig); ok {
							db.Host = "" // Missing host
							db.Port = "5432"
							db.User = "test-user"
							db.Pass = "test-pass"
							db.Name = "test-db"
						}
					}
					return nil
				},
			},
			wantErr:        true,
			expectedErrMsg: "validate database config",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Clear environment
			t.Setenv("API_ENV", "")
			t.Setenv("API_PORT", "")
			t.Setenv("API_SERVICE_NAME", "")

			// Setup test environment
			tt.setupEnv()

			// Create loader with mock fetcher
			loader := NewConfigLoader(tt.mockFetcher)

			// Execute
			ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
			defer cancel()

			cfg, err := loader.Load(ctx)

			// Verify
			if tt.wantErr {
				if err == nil {
					t.Errorf("expected error but got none")
				} else if tt.expectedErrMsg != "" && !contains(err.Error(), tt.expectedErrMsg) {
					t.Errorf("expected error containing %q, got %q", tt.expectedErrMsg, err.Error())
				}
			} else {
				if err != nil {
					t.Errorf("unexpected error: %v", err)
				}
				if cfg == nil {
					t.Error("expected config but got nil")
				}
			}
		})
	}
}

func TestConfig_GetDatabaseDSN(t *testing.T) {
	cfg := &Config{
		CoreDB: DatabaseConfig{
			Host: "localhost",
			Port: "5432",
			User: "user",
			Pass: "pass",
			Name: "dbname",
		},
	}

	expected := "host=localhost port=5432 user=user password=pass dbname=dbname sslmode=require"
	got := cfg.GetDatabaseDSN()

	if got != expected {
		t.Errorf("GetDatabaseDSN() = %v, want %v", got, expected)
	}
}

func TestConfig_IsProduction(t *testing.T) {
	tests := []struct {
		name string
		env  Environment
		want bool
	}{
		{"production environment", EnvPrd, true},
		{"development environment", EnvDev, false},
		{"test environment", EnvTest, false},
		{"staging environment", EnvStg, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cfg := &Config{
				API: APIConfig{
					Env: tt.env,
				},
			}

			if got := cfg.IsProduction(); got != tt.want {
				t.Errorf("IsProduction() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestConfig_IsDevelopment(t *testing.T) {
	tests := []struct {
		name string
		env  Environment
		want bool
	}{
		{"development environment", EnvDev, true},
		{"test environment", EnvTest, true},
		{"production environment", EnvPrd, false},
		{"staging environment", EnvStg, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cfg := &Config{
				API: APIConfig{
					Env: tt.env,
				},
			}

			if got := cfg.IsDevelopment(); got != tt.want {
				t.Errorf("IsDevelopment() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestValidateAPIConfig(t *testing.T) {
	loader := &DefaultConfigLoader{}

	tests := []struct {
		name    string
		api     APIConfig
		wantErr bool
	}{
		{
			name: "valid config",
			api: APIConfig{
				Env:         EnvDev,
				Port:        "8080",
				ServiceName: "test-service",
			},
			wantErr: false,
		},
		{
			name: "missing environment",
			api: APIConfig{
				Env:         "",
				Port:        "8080",
				ServiceName: "test-service",
			},
			wantErr: true,
		},
		{
			name: "invalid port - not a number",
			api: APIConfig{
				Env:         EnvDev,
				Port:        "abc",
				ServiceName: "test-service",
			},
			wantErr: true,
		},
		{
			name: "invalid port - out of range",
			api: APIConfig{
				Env:         EnvDev,
				Port:        "99999",
				ServiceName: "test-service",
			},
			wantErr: true,
		},
		{
			name: "invalid environment value",
			api: APIConfig{
				Env:         Environment("invalid"),
				Port:        "8080",
				ServiceName: "test-service",
			},
			wantErr: true,
		},
		{
			name: "missing service name",
			api: APIConfig{
				Env:         EnvDev,
				Port:        "8080",
				ServiceName: "",
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := loader.validateAPIConfig(&tt.api)
			if (err != nil) != tt.wantErr {
				t.Errorf("validateAPIConfig() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestValidateDatabaseConfig(t *testing.T) {
	loader := &DefaultConfigLoader{}

	tests := []struct {
		name    string
		db      DatabaseConfig
		wantErr bool
	}{
		{
			name: "valid config",
			db: DatabaseConfig{
				Host: "localhost",
				Port: "5432",
				User: "user",
				Pass: "pass",
				Name: "dbname",
			},
			wantErr: false,
		},
		{
			name: "missing host",
			db: DatabaseConfig{
				Host: "",
				Port: "5432",
				User: "user",
				Pass: "pass",
				Name: "dbname",
			},
			wantErr: true,
		},
		{
			name: "invalid port",
			db: DatabaseConfig{
				Host: "localhost",
				Port: "invalid",
				User: "user",
				Pass: "pass",
				Name: "dbname",
			},
			wantErr: true,
		},
		{
			name: "missing user",
			db: DatabaseConfig{
				Host: "localhost",
				Port: "5432",
				User: "",
				Pass: "pass",
				Name: "dbname",
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := loader.validateDatabaseConfig(&tt.db)
			if (err != nil) != tt.wantErr {
				t.Errorf("validateDatabaseConfig() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

// Helper function
func contains(s, substr string) bool {
	return len(s) >= len(substr) && s[:len(substr)] == substr || len(s) > len(substr) && contains(s[1:], substr)
}
