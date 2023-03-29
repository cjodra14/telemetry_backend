package grpc

import (
	"context"

	"github.com/cjodra14/telemetry_backend/api/models"
	grpcConfig "github.com/cjodra14/telemetry_commons/configuration/grpc"
	commonsGRPC "github.com/cjodra14/telemetry_commons/services/grpc"
)

type grpcTelemetryClient struct {
	client TelemetryServiceClient
}

// var (
// 	_ client.CellClient = (*grpcTelemetryClient)(nil)
// )

func NewCellClient(ctx context.Context, config grpcConfig.GRPCClientConfiguration) (*grpcTelemetryClient, error) {
	conn, err := commonsGRPC.Client(ctx, config)
	if err != nil {
		return &grpcTelemetryClient{}, err
	}

	return &grpcTelemetryClient{
		client: NewTelemetryServiceClient(conn),
	}, nil
}

func (gRPCClient *grpcTelemetryClient) GetUserTelemetries(ctx context.Context, userID string) ([]models.Telemetry, error) {
	response, err := gRPCClient.client.GetUserTelemetry(ctx, &UserTelemetryRequest{
		UserId: userID,
	})
	if err != nil {
		return []models.Telemetry{}, err
	}

	userTelemetries := []models.Telemetry{}
	for _, telemetry := range response.Telemetries {
		userTelemetries = append(userTelemetries, models.Telemetry{
			Timestamp: telemetry.Timestamp,
			UserID:    telemetry.UserId,
			GPS: models.GPS{
				Latitude:            telemetry.Gps.Latitude,
				Longitude:           telemetry.Gps.Longitude,
				Speed:               telemetry.Gps.Speed,
				Direction:           telemetry.Gps.Direction,
				ConnectedSatellites: telemetry.Gps.ConnectedSatellites,
			},
			Accelerometer: models.Accelerometer{
				AngleX: telemetry.Accelerometer.AngleX,
				AngleY: telemetry.Accelerometer.AngleY,
				AngleZ: telemetry.Accelerometer.AngleZ,
				GForce: telemetry.Accelerometer.GForce,
			},
		})
	}

	return userTelemetries, nil
}

func (gRPCClient *grpcTelemetryClient) SaveTelemetry(ctx context.Context, telemetry models.Telemetry) error {
	_, err := gRPCClient.client.SaveTelemetry(ctx, &Telemetry{
		Timestamp: telemetry.Timestamp,
		UserId:    telemetry.UserID,
		Gps: &GPS{
			Latitude:            telemetry.GPS.Latitude,
			Longitude:           telemetry.GPS.Longitude,
			Speed:               telemetry.GPS.Speed,
			Direction:           telemetry.GPS.Direction,
			ConnectedSatellites: telemetry.GPS.ConnectedSatellites,
		},
		Accelerometer: &Accelerometer{
			AngleX: telemetry.Accelerometer.AngleX,
			AngleY: telemetry.Accelerometer.AngleY,
			AngleZ: telemetry.Accelerometer.AngleZ,
			GForce: telemetry.Accelerometer.GForce,
		},
	})
	if err != nil {
		return err
	}

	return nil
}
