package storage

import (
	"context"

	"github.com/cjodra14/telemetry-iot/telemetry_backend/api/models"
)

type TelemetryStorage interface {
	SaveTelemetry(ctx context.Context, telemetry models.Telemetry) error
	GetUserTelemetries(ctx context.Context, userID string) ([]models.Telemetry, error)
}
