package helper

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetENV(t *testing.T) {
	testCases := []struct {
		Name      string
		Path      string
		WantError bool
	}{
		{
			Name:      "success",
			Path:      "../.env",
			WantError: false,
		}, {
			Name:      "failed",
			Path:      "../../.env",
			WantError: true,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.Name, func(t *testing.T) {
			if testCase.WantError {
				_, err := GetENV(testCase.Path)
				assert.NotNil(t, err)
			} else {
				env, _ := GetENV(testCase.Path)
				assert.NotNil(t, env)
			}
		})
	}
}

func TestIsStatusValid(t *testing.T) {
	testCases := []struct {
		Name        string
		Status      string
		Expectation bool
	}{
		{
			Name:        "success",
			Status:      "publish",
			Expectation: true,
		}, {
			Name:        "success",
			Status:      "draft",
			Expectation: true,
		}, {
			Name:        "success",
			Status:      "thrash",
			Expectation: true,
		}, {
			Name:        "failed",
			Status:      "failed status",
			Expectation: false,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.Name, func(t *testing.T) {
			actual := IsStatusValid(testCase.Status)
			assert.Equal(t, testCase.Expectation, actual)
		})
	}
}
