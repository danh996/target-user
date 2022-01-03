gen-proto:
	protoc -I . --grpc-gateway_out ./ --proto_path=proto --go_out=. --go-grpc_out=:. --validate_out=lang=go:. proto/*.proto

start-api:
	go run apis/main.go

start-server:
	go run cmd/main.go