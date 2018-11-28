# eco

Experiment building an echo server with Go

# What is it?

Eco is a simple GRPC [echo protocol](https://en.wikipedia.org/wiki/Echo_Protocol) implementation. It will be use for experimentation with[GRPC]( https://grpc.io)

# Usage

Require Golang 1.9 or above. Installation instructions [here](https://golang.org/doc/install)

## Install

```
$ go get -u github.com/cored/eco
```

## Server

The easiest way to start using the service is by running:

```
$ make start-server
```

or

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

## Using grpcurl

# Development


