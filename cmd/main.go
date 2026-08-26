package main

import (
	"Service/internal/app"
	"github.com/joho/godotenv"
	"log"
)

func init() {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal("error load .env")
	}
	log.Println("loaded .env")
}

func main() {
	a := app.New()

	if err := a.Start(); err != nil {
		log.Fatal(err)
	}

	defer a.Close()
}
