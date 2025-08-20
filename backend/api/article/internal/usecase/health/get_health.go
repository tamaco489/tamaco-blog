// Package health provides health check business logic.
package health

import (
	"context"

	"github.com/tamaco489/tamaco-blog/backend/api/article/internal/gen"
)

// HealthcheckUseCase handles health check business logic
type HealthcheckUseCase interface {
	Healthcheck(ctx context.Context) (*gen.HealthCheck, error)
}

type healthcheckUseCase struct{}

// NewHealthcheckUseCase creates a new health usecase
func NewHealthcheckUseCase() HealthcheckUseCase {
	return &healthcheckUseCase{}
}

// Healthcheck implements HealthcheckUseCase
func (u *healthcheckUseCase) Healthcheck(ctx context.Context) (*gen.HealthCheck, error) {
	return &gen.HealthCheck{
		Message: "OK",
	}, nil
}
