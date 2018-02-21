build:
			protoc -I proto/consignment consignment.proto --go_out=plugins=micro:$(GOPATH)/src/github.com/giancarlopetrini/microservices/consignment-service/proto/consignment/
			docker build -t consignment-service .
run:
			docker run -p 50051:50051 \
			-e MICRO_SERVER_ADDRESS=:50051 \
			-e MICRO_REGISTRY=mdns consignment-service
