package main

import (
	"log"

	"github.com/markovidakovic/surl/server/internal/rest"
)

func main() {
	srvr := rest.NewServer()
	log.Fatal(srvr.Start())
}
