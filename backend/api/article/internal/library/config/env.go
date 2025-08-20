// Package config provides configuration management for the application,
// including environment variables, AWS configuration, and secrets management.
package config

// Environment represents the application environment
type Environment string

// Environment constants
const (
	EnvDev  Environment = "dev"
	EnvTest Environment = "test"
	EnvStg  Environment = "stg"
	EnvPrd  Environment = "prd"
)

// String returns the string representation of the environment
func (e Environment) String() string {
	return string(e)
}

// IsValid checks if the environment value is valid
func (e Environment) IsValid() bool {
	switch e {
	case EnvDev, EnvTest, EnvStg, EnvPrd:
		return true
	default:
		return false
	}
}

// IsProduction returns true if the environment is production
func (e Environment) IsProduction() bool {
	return e == EnvPrd
}

// IsDevelopment returns true if the environment is development or test
func (e Environment) IsDevelopment() bool {
	return e == EnvDev || e == EnvTest
}
