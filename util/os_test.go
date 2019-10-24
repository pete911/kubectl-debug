package util

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func Test_GetStringEnvVarThatIsNotSet(t *testing.T) {

	actual := GetStringEnvVar("TEST_STRING_VALUE", "xyz")
	assert.Equal(t, "xyz", actual)
}

func Test_GetStringEnvVar(t *testing.T) {

	cases := []struct {
		envValue     string
		defaultValue string
		expected     string
	}{
		{"abc", "def", "abc"},
		{"", "def", "def"},
	}

	for _, c := range cases {
		rollback := SetEnv("TEST_STRING_VALUE", c.envValue)
		actual := GetStringEnvVar("TEST_STRING_VALUE", c.defaultValue)
		rollback()
		assert.Equal(t, c.expected, actual)
	}
}

func Test_GetBoolEnvVar(t *testing.T) {

	cases := []struct {
		envValue     string
		defaultValue bool
		expected     bool
	}{
		{"true", false, true},
		{"TrUe", false, false},
		{"True", false, true},
		{"TRUE", false, true},
		{"1", false, true},
		{"", false, false},
		{"abc", false, false},
	}

	for _, c := range cases {
		rollback := SetEnv("TEST_BOOL_VALUE", c.envValue)
		actual := GetBoolEnvVar("TEST_BOOL_VALUE", c.defaultValue)
		rollback()
		assert.Equal(t, c.expected, actual, fmt.Sprintf("env: %s, default: %t", c.envValue, c.defaultValue))
	}
}

// --- helper functions ---

func SetEnv(key, val string) (rollback func()) {

	prev := os.Getenv(key)
	os.Setenv(key, val)
	return func() { os.Setenv(key, prev) }
}
