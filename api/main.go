package main

import (
	"github.com/p9c/cms/connection"
)

const (
	Address = ":9999"

	addr = "localhost:4242"

	message = "foobar"
)

func main() {
	connection.StartServerModule()
}
