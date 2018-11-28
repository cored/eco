# eco

Experiment building an echo server with Go

# What is it?

Eco is a simple GRPC [echo protocol](https://en.wikipedia.org/wiki/Echo_Protocol) implementation. It will be use for experimentation with[ GRPC]( https://grpc.io)

# Usage

Require Golang 1.9 or above. Installation instructions [here](https://golang.org/doc/install)

## Install

```
$ go get -u github.com/cored/eco
```

## Server

The easiest way to start using the service is by running:

```
$ ./eco
```

That will start the server on port `8080`. The port can be configured by using the `-p` option:

```
$ ./eco -p <port>
```

## Client

The purpose of the client is to provide an easy way to interact with the
server.

You can use it by specifying a port for the server and the text to send to the
server. If you run it without specifying anything it will send `PING` to
the default server listening port.

```
$ ./client/client -s <server:port> -t <text>
```

# Development

First you want to generate the stubs based on the protobuf schema. Note that
this requires the Go gRPC runtime and plug-in installed on your machine,
including protoc in v3 set up, see grpc.io for the steps.

```
$ make build-protos
```

Executing above command results in the auto-generated file protos/eco-schema.pb.go. Do not manually edit this file,
or put in other words: if you add a new message or service to the schema defined in eco-schema.proto
just run the make task again and you'll get an updated version of eco-schema.pb.go in the protos/ directory
as a result.

## Tests

You can run the tests using the following command:

```
$ make tests
```

## Why protobuf?

As part of the experiment I wanted to find a way to enforce a contract between
the client and the server. JSON schema is too loose in terms of enforcing type
checks and also slower than protobuf. For more information check the protobuf
site - https://developers.google.com/protocol-buffers/

## TODOs

- [ ] Add authentication
- [ ] Add manifests files for different deployment backend infrastructure
- [ ] Dockerized
