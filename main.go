package main

import (
	"flag"
	"fmt"

	"github.com/cored/eco/server"
)

func main() {
	port := flag.Int("p", 8080, "port to listen to")
	flag.Parse()

	server.Start(fmt.Sprintf(":%d", *port))
}
