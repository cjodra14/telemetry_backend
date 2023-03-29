package services

import (
	"context"
	"errors"

	"github.com/cjodra14/telemetry-iot/telemetry_backend/api/models"
)

var (
	ErrInternalTelemetryService = errors.New("telmetry service internal error")
	ErrTelemetryNotFound        = errors.New("telmetry not found error")
)

type TelemetryService interface {
	SaveTelemetry(ctx context.Context, telemetry models.Telemetry) error
	GetuserTelemetries(ctx context.Context, userID string) ([]models.Telemetry, error)
}
