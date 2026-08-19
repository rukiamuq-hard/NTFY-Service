package main

import (
	"Service/internal/app"
	"log"
)

func main() {
	a := app.New()

	if err := a.Start(); err != nil {
		log.Fatal(err)
	}

	defer a.Close()
}
