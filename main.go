package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func getPassword() string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	port := os.Getenv("PORT")
	return port
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	connectToDB()
}
