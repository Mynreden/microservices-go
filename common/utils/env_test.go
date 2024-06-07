package utils

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEnvString(t *testing.T) {
	// Save the original environment variable value
	originalValue, originalExists := os.LookupEnv("TEST_KEY")

	// Restore the original environment variable value after the test
	defer func() {
		if originalExists {
			err := os.Setenv("TEST_KEY", originalValue)
			if err != nil {
				return
			}
		} else {
			err := os.Unsetenv("TEST_KEY")
			if err != nil {
				return
			}
		}
	}()

	t.Run("Environment variable exists", func(t *testing.T) {
		err := os.Setenv("TEST_KEY", "test_value")
		if err != nil {
			return
		}
		result := EnvString("TEST_KEY", "fallback_value")
		assert.Equal(t, "test_value", result)
	})

	t.Run("Environment variable does not exist", func(t *testing.T) {
		err := os.Unsetenv("TEST_KEY")
		if err != nil {
			return
		}
		result := EnvString("TEST_KEY", "fallback_value")
		assert.Equal(t, "fallback_value", result)
	})
}
