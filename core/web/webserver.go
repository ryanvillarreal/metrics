package web

import (
	"log"
)

func run() error {
	return nil
}

// Start used to init the db into mem as early as possible
func Start() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}
