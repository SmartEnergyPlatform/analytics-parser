package main

import (
	"log"
	"github.com/joho/godotenv"
	"fmt"
	"parsing-service/api"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Print("Error loading .env file")
	}
	fmt.Println("GO")
	api.CreateServer()
}