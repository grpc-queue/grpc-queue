test: #just in case we forget which action we want
	go test ./... -v
	
generate-grpc:
	protoc --go_out=. --go-grpc_out=. api/proto/v1/queue.proto