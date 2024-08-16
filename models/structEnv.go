package models

import (
	"fmt"
	"github.com/caarlos0/env/v9"

	"github.com/joho/godotenv"
)

type EnVariatiles struct {
	PostgresUser     string `env:"POSTGRES_USER,required,notEmpty"`
	PostgresPassword string `env:"POSTGRES_PASSWORD,required,notEmpty"`
}

const envFilePath = ".env.local"

func LoadEnv() *EnVariatiles {
	environmentVariables := EnVariatiles{}
	if err := godotenv.Load(envFilePath); err != nil {
		panic(fmt.Sprintf("Failed to load .env.local file: %s", err.Error()))
	}
	if err := env.Parse(&environmentVariables); err != nil {
		panic(fmt.Sprintf("Failed to parse environment variables: %s", err.Error()))
	}

	return &environmentVariables
}
