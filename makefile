test: #just in case we forget which action we want
	go test ./... -v  -count=1

benchmark:
	go test ./... -bench=.

rm-data: #we can quickly reset data for testing porpuses
	rm -rf $(GRPCQUEUE_DATA_PATH)/queuedata

run: 
	go run cmd/server/main.go

restart: rm-data run

generate-grpc:
	protoc --go_out=. --go-grpc_out=. api/proto/v1/queue.proto