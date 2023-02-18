package main

import (
	"log"

	"github.com/o1vanova/calculator/cmd/calculator/internal/pkg/app"
)

func main() {
	app, err := app.New()
	if err != nil {
		log.Fatal(err)
	}

	err = app.Start()
	if err != nil {
		log.Fatal(err)
	}
}
