package infrastructure

import (
	"errors"
	"io/fs"
	"log"
	"os"

	"github.com/joho/godotenv"
)

const envFile = "internal/infrastructure/.env"

// LoadEnvVariables reads settings from the local .env file when it exists
// (copy .env.example to create it); otherwise it uses the process environment.
func LoadEnvVariables() {
	if err := godotenv.Load(envFile); err != nil && !errors.Is(err, fs.ErrNotExist) {
		log.Fatalf("Error loading %s: %v", envFile, err)
	}

	for _, key := range []string{"DBURL", "JWTSECRET"} {
		if os.Getenv(key) == "" {
			log.Fatalf("%s is not set; see internal/infrastructure/.env.example", key)
		}
	}
}
