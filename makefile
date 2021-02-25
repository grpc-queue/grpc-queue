test: #just in case we forget which action we want
	go test ./... -v

rm-data: #we can quickly reset data for testing porpuses
	rm -rf $(GRPCQUEUE_DATA_PATH)/queuedata
	
generate-grpc:
	protoc --go_out=. --go-grpc_out=. api/proto/v1/queue.proto