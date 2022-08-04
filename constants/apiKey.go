package constants

import (
	"fmt"
	"github.com/joho/godotenv"
	"os"
)

func LoadEnv() {
	envErr := godotenv.Load(".env")
	if envErr != nil {
		fmt.Printf("Could not load .env file")
		os.Exit(1)
	}

}
func GetApiKey() string {
	apiKey, isSet := os.LookupEnv("RANDOM_API_KEY")
	if !isSet {
		fmt.Printf("RANDOM_API_KEY env value is not set")
		os.Exit(1)
	}
	return apiKey

}
