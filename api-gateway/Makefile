gen-user:
	@protoc --go_out=. --go-grpc_out=. protos/user-service/user.proto 
gen-device:
	@protoc --go_out=. --go-grpc_out=. protos/device-service/device.proto 
gen-command:
	@protoc --go_out=. --go-grpc_out=. protos/command-service/command.proto 
get:
	@go mod tidy  
run:
	@go run cmd/main.go

SWAGGER := ~/go/bin/swag
SWAGGER_DOCS := docs
SWAGGER_INIT := $(SWAGGER) init -g ./api/api.go -o $(SWAGGER_DOCS)

swag-gen:
	$(SWAGGER_INIT)
