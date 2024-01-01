proto:
	protoc \
		--go_out=. \
		--go_opt=paths=source_relative \
		--go-grpc_out=. \
		--go-grpc_opt=paths=source_relative \
		pkg/pb/*.proto

docs:
	swag init --instanceName $(name)

test: proto
	go test ./pkg/tests -v

server: proto docs
	go run main.go