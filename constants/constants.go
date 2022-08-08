package constants

import (
	"fmt"
	"github.com/joho/godotenv"
	"os"
	"strconv"
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

func GetMin() int {
	minValue, isSet := os.LookupEnv("MIN_VALUE")
	if !isSet {
		fmt.Printf("MIN_VALUE env value is not set")
		os.Exit(1)
	}
	intMinValue, err := strconv.Atoi(minValue)
	if err != nil {
		fmt.Printf("Error converting MIN_VALUE to int")
	}
	return intMinValue
}

func GetMax() int {
	maxValue, isSet := os.LookupEnv("MAX_VALUE")
	if !isSet {
		fmt.Printf("MAX_VALUE env value is not set")
		os.Exit(1)
	}
	intMaxValue, err := strconv.Atoi(maxValue)
	if err != nil {
		fmt.Printf("Error converting MAX_VALUE to int")
	}
	return intMaxValue
}
