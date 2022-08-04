package utils

import (
	"github.com/joho/godotenv"
	"github.com/pkg/errors"
)

func LoadEnv() (err error) {
	err = godotenv.Load("deployments/.env")
	if err != nil {
		err = errors.Wrap(err, "load .env file")
	}
	return
}
