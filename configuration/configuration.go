package configuration

import (
	"github.com/cjodra14/telemetry_commons/configuration/grpc"
	"github.com/cjodra14/telemetry_commons/configuration/mongo"
	http "github.com/cjodra14/telemetry_commons/configuration/rest"
)

type Configuration struct {
	Server     http.HTTPServerConfiguration `configuration:"server"`
	GRPCServer grpc.GRPCServerConfiguration `configuration:"grpc_server"`
	Log        LogConfiguration             `configuration:"log"`
	Storage    StorageConfiguration         `configuration:"storage"`
}

type StorageType string

const (
	MongoStorageType  StorageType = "mongo"
	MemoryStorageType StorageType = "memory"
)

type LogConfiguration struct {
	Name   string `configuration:"name" efault:"report-service" usage:"name showing in logger"`
	Level  string `configuration:"level" default:"info" usage:"level of logger (debug,info,error)" `
	Format string `configuration:"format" default:"text" usage:"format of log (text,json)"`
}

type StorageConfiguration struct {
	StorageType     StorageType              `configuration:"type" default:"mongo" usage:"storage type memory,mongo"`
	Mongo           mongo.MongoConfiguration `configuration:"mongo" usage:"mongo application service type"`
	MongoCollection string                   `configuration:"mongo_collection" default:"test" usage:"mongo documents collection"`
}
