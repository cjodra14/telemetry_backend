doc:
	swag init --parseInternal --parseDepth 1 --instanceName telemetry_backend
	swagger generate markdown -f$(shell pwd)/docs/telemetry_backend_swagger.yaml --output api-doc.md
proto:
	protoc --proto_path=./api/models/proto/ --go_out=./api/client/grpc/ --go_opt=paths=source_relative --go-grpc_out=./api/client/grpc/ --go-grpc_opt=paths=source_relative ./api/models/proto/telemetry.proto

