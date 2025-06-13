package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv" // Import the godotenv package
)

func main() {
	// Load environment variables from .env file
	err := godotenv.Load("../.env.api_key")
	if err != nil {
		fmt.Println("Error loading .env file")
		return
	}
	apiKey := os.Getenv("API_KEY")

	fmt.Println("API Key:", apiKey)

}
