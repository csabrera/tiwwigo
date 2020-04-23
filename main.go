package main

import (
	"log"

	"github.com/csabrera/tiwwigo/bd"
	"github.com/csabrera/tiwwigo/handlers"
)

func main() {
	if bd.ChequeoConnection() == 0 {
		log.Fatal("Sin conexión a la BD")
		return
	}
	handlers.Manejadores()
}
