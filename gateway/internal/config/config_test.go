package configs

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestGetConfig(t *testing.T) {
	// Set environment variables for the test
	os.Setenv("ADDR", ":9090")
	os.Setenv("USER_ADDR", "localhost:60061")
	os.Setenv("POST_ADDR", "localhost:60062")
	os.Setenv("DSN", "pgsql:host=localhost;port=5432;dbname=testdb;user=testuser;password=testpassword")
	os.Setenv("STATIC_DIR", "/test_static")

	// Get the config
	cfg, err := GetConfig()

	// Ensure no errors
	assert.NoError(t, err)

	// Check the config values
	assert.Equal(t, ":9090", cfg.Addr)
	assert.Equal(t, "localhost:60061", cfg.UsersAddr)
	assert.Equal(t, "localhost:60062", cfg.PostsAddr)
	assert.Equal(t, "pgsql:host=localhost;port=5432;dbname=testdb;user=testuser;password=testpassword", cfg.DB.DSN)
	assert.Equal(t, "/test_static", cfg.StaticDir)

	// Clean up environment variables
	os.Unsetenv("ADDR")
	os.Unsetenv("USER_ADDR")
	os.Unsetenv("POST_ADDR")
	os.Unsetenv("DSN")
	os.Unsetenv("STATIC_DIR")
}
