package middleware

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

//	func GetEnv(key, defaultValue string) string {
//		value, exists := os.LookupEnv(key)
//		if !exists {
//			return defaultValue
//		}
//		return value
//	}
func GetEnv(key string) string {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Environment variable %s not set", key)
	}
	return os.Getenv(key)
}
