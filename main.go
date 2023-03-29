package main

import (
	"context"
	"flag"
	"log"

	"github.com/cjodra14/telemetry_backend/configuration"
	"github.com/cjodra14/telemetry_backend/server/http"
	"github.com/cjodra14/telemetry_backend/services/telemetry_service"
	"github.com/cjodra14/telemetry_backend/storage"
	telemetryMongo "github.com/cjodra14/telemetry_backend/storage/mongo"
	configurationController "github.com/cjodra14/telemetry_commons/configuration"
	commonsMongo "github.com/cjodra14/telemetry_commons/services/mongo"
	"github.com/oklog/run"
	"github.com/sirupsen/logrus"
)


// @title           Telemetry IoT
// @version         1.0
// @description     Telemetry REST API of IoT telemetry device

// @contact.name   Christian Jodra
// @contact.email c.jodra14@gmail.com

// @host      localhost:8080
// @BasePath  /
func main() {

	conf := &configuration.Configuration{}

	err := configurationController.Init(conf)
	if err != nil {
		logrus.Fatal(err)
	}

	var configFile string

	flag.StringVar(&configFile, "config", "", "config file")

	flag.Parse()

	if configFile != "" {
		err = configurationController.FromFile(conf, configFile)
		if err != nil {
			logrus.Fatal(err)
		}
	}

	logrus.Infof("conf:%+v", conf)

	logrus.SetLevel(logrus.Level(5))

	var telemetryDAO storage.TelemetryStorage

	ctx, cancelCtx := context.WithCancel(context.Background())
	defer cancelCtx()

	switch conf.Storage.StorageType {
	case configuration.MongoStorageType:
		mongoService, err := commonsMongo.NewMongoService(ctx, conf.Storage.Mongo)
		if err != nil {
			logrus.Fatal(err)
		}

		telemetryDAO, err = telemetryMongo.NewMongoTelemetryStorage(mongoService, conf.Storage)
		if err != nil {
			logrus.Fatal(err)
		}
	default:
		logrus.Fatalf("you have selected the '%s' storage type which is not recognized", conf.Storage.StorageType)
	}

	telemetryService := telemetry_service.New(telemetryDAO)

	var group run.Group

	group.Add(func() error {

		return http.Init(conf.Server, telemetryService)
	}, func(e error) {
		log.Fatal(e)
	})

	// group.Add(func() error {
	// 	return serverGrpc.InitTelemetryServiceServer(conf.GRPCServer, telemetryService)
	// }, func(e error) {
	// 	log.Fatal(e)
	// })

	if err := group.Run(); err != nil {
		log.Fatal(err)
	}
}
