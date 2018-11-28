build-protos:
	protoc -I ./protos eco-schema.proto --go_out=plugins=grpc:./protos/.

clean:
	@rm -rf eco

build:
	go build

tests:
	go test server/server_test.go
