package memory

import (
	"context"
	"errors"
	"strconv"

	"github.com/cjodra14/telemetry_backend/api/models"
)

type MemoryTelemetryStorage struct {
	memory map[string]models.Telemetry
}

func NewMongoTelemetryStorage() (*MemoryTelemetryStorage, error) {

	return &MemoryTelemetryStorage{
		memory: map[string]models.Telemetry{},
	}, nil
}

func (memoryTelemetryStorage *MemoryTelemetryStorage) SaveTelemetry(ctx context.Context, telemetry models.Telemetry) error {
	memoryTelemetryStorage.memory[telemetry.UserID+strconv.Itoa(int(telemetry.Timestamp))] = telemetry

	return nil
}

func (memoryTelemetryStorage *MemoryTelemetryStorage) GetUserTelemetries(ctx context.Context, userID string) ([]models.Telemetry, error) {
	userTelemetries := []models.Telemetry{}

	for _, telemetry := range memoryTelemetryStorage.memory {
		if telemetry.UserID == userID {
			userTelemetries = append(userTelemetries, telemetry)
		}
	}

	if len(userTelemetries) == 0 {
		return []models.Telemetry{}, errors.New("user telemetries not found")
	}

	return userTelemetries, nil
}
