package telemetry_service

import (
	"context"

	"github.com/cjodra14/telemetry-iot/telemetry_backend/api/models"
	"github.com/cjodra14/telemetry-iot/telemetry_backend/storage"
)

type TelemetryService struct {
	telemetryStorage storage.TelemetryStorage
}

func New(telemetryStorage storage.TelemetryStorage) *TelemetryService {
	return &TelemetryService{
		telemetryStorage: telemetryStorage,
	}
}

func (telmetryService *TelemetryService) SaveTelemetry(ctx context.Context, telemetry models.Telemetry) error {
	if err := telmetryService.telemetryStorage.SaveTelemetry(ctx, telemetry); err != nil {
		return err
	}

	return nil
}

func (telmetryService *TelemetryService) GetuserTelemetries(ctx context.Context, userID string) ([]models.Telemetry, error) {
	userTelemetries, err := telmetryService.telemetryStorage.GetUserTelemetries(ctx, userID)
	if err != nil {
		return []models.Telemetry{}, err
	}

	return userTelemetries, nil
}
