package mongo

import (
	"context"
	"errors"

	"github.com/cjodra14/telemetry_backend/api/models"
	"github.com/cjodra14/telemetry_backend/configuration"
	commonsMongo "github.com/cjodra14/telemetry_commons/services/mongo"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type MongoTelemetryStorage struct {
	service commonsMongo.Service
	dao     *commonsMongo.DAO[models.Telemetry]
}

func NewMongoTelemetryStorage(mongoService *commonsMongo.Service, configuration configuration.StorageConfiguration) (*MongoTelemetryStorage, error) {
	collectionDAO, err := commonsMongo.NewMongoDAO[models.Telemetry](mongoService, configuration.MongoCollection)
	if err != nil {
		switch {
		case errors.Is(err, commonsMongo.MongoServiceErrVoidCollection):
			return &MongoTelemetryStorage{}, err
		default:
			return &MongoTelemetryStorage{}, err
		}
	}

	return &MongoTelemetryStorage{
		service: *mongoService,
		dao:     collectionDAO,
	}, nil
}

func (mongoTelemetryStorage *MongoTelemetryStorage) SaveTelemetry(ctx context.Context, telemetry models.Telemetry) error {
	if _, err := mongoTelemetryStorage.dao.Insert(ctx, telemetry); err != nil {
		switch {
		case errors.Is(err, commonsMongo.ErrAlreadyExists):
			return err
		default:
			return err
		}
	}

	return nil
}

func (mongoTelemetryStorage *MongoTelemetryStorage) GetUserTelemetries(ctx context.Context, userID string) ([]models.Telemetry, error) {
	userTelemetries, err := mongoTelemetryStorage.dao.Find(ctx, primitive.D{{Key: "user_id", Value: userID}})
	if err != nil {
		return []models.Telemetry{}, err
	}

	return userTelemetries, nil
}
