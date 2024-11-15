gen-user:
	@protoc --go_out=. --go-grpc_out=. proto/user.proto 
get:
	@go mod tidy 
run:
	@go run cmd/main.go