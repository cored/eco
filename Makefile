compile-proto:
	protoc -I . eco-schema.proto --go_out=plugins=grpc:.
