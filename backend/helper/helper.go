package helper

import (
	"fmt"

	"github.com/joho/godotenv"
)

func GetENV(path string) (map[string]string, error) {
	env, err := godotenv.Read(path)
	if err != nil {
		return env, fmt.Errorf("error env : %v", err)
	}

	return env, err
}
