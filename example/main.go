package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv" // Import the godotenv package
)

func getApiKey() string {
	// check for .env.api_key file in the parent directory
	if _, err := os.Stat("../.env.api_key"); os.IsNotExist(err) {
		fmt.Println("No .env.api_key file found in the parent directory")

		// TODO: getApiKeyConsole()

	} else {
		fmt.Println("Found .env.api_key file in the parent directory")
	}

	// Load environment variables from .env file
	err := godotenv.Load("../.env.api_key")
	if err != nil {
		fmt.Println("Error loading .env file")
		return ""
	}
	return os.Getenv("API_KEY")
}

func main() {
	apiKey := getApiKey()
	if apiKey == "" {
		return
	}

	fmt.Println("API Key:", apiKey)

	// Use the API key for your application logic
}
