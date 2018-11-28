compile-proto:
	protoc -I ./protos eco-schema.proto --go_out=plugins=grpc:./protos/.
