package helper

import (
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
)

type responseAPI struct {
	Meta meta        `json:"meta"`
	Data interface{} `json:"data"`
}

type meta struct {
	Status   int    `json:"status"`
	Messsage string `json:"message"`
}

func ResponseAPI(message string, status int, data interface{}) *responseAPI {
	return &responseAPI{
		Meta: meta{
			Status:   status,
			Messsage: message,
		},
		Data: data,
	}
}

func GetENV(path string) (map[string]string, error) {
	env, err := godotenv.Read(path)
	if err != nil {
		return env, fmt.Errorf("error env : %v", err)
	}

	return env, err
}

func IsStatusValid(status string) bool {
	statusMap := map[string]string{
		"publish": "publish",
		"draft":   "draft",
		"thrash":  "thrash",
	}

	if _, isExist := statusMap[status]; !isExist {
		return false
	}

	return true
}

func GenerateErrorBinding(err error) (errorBinding []string) {
	for _, e := range err.(validator.ValidationErrors) {
		errorBinding = append(errorBinding, e.Error())
	}

	return errorBinding
}

func IsIDValid(ID int) (bool, error) {
	if ID < 1 {
		return false, fmt.Errorf("id harus lebih besar dari 0")
	}

	return true, nil
}
