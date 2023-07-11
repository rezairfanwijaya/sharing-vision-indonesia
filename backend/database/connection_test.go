package database

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetConnection(t *testing.T) {
	testCases := []struct {
		Name      string
		ENVPath      string
		WantError bool
	}{
		{
			Name:      "success",
			ENVPath:      "../.env",
			WantError: false,
		}, {
			Name:      "failed",
			ENVPath:      "../../.env",
			WantError: true,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.Name, func(t *testing.T) {
			if testCase.WantError {
				_, err := NewConnection(testCase.ENVPath)
				assert.NotNil(t, err)
			} else {
				env, _ := NewConnection(testCase.ENVPath)
				assert.NotNil(t, env)
			}
		})
	}
}
