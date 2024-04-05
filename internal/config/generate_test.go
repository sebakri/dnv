package config_test

import (
	"path/filepath"
	"testing"

	"github.com/sebakri/dnv/internal/config"
	"github.com/stretchr/testify/assert"
)

func TestLoad(t *testing.T) {

	tests := []struct {
		name     string
		config   string
		expected *config.Config
	}{
		{
			name: "simple config",
			config: `environment_variables: {
						foo: "bar"
						pi: "3.14159"
						"UNDER_SCORE": "true"
						FOO_BAR: "baz"
					}`,
			expected: &config.Config{
				EnvironmentVariables: map[string]string{
					"foo":         "bar",
					"pi":          "3.14159",
					"UNDER_SCORE": "true",
					"FOO_BAR":     "baz",
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cfg, err := config.Parse(tt.config)

			assert.Nil(t, err)
			assert.Equal(t, tt.expected, cfg)
		})
	}
}

func TestLookup(t *testing.T) {
	tests := []struct {
		name     string
		path     string
		expected []string
	}{
		{
			name:     "simple lookup",
			path:     filepath.Join("testdata", "lookup"),
			expected: []string{filepath.Join("testdata", "lookup", ".dnv")},
		},
		{
			name:     "lookup with subdirectories",
			path:     filepath.Join("testdata", "lookup", "a", "aa"),
			expected: []string{filepath.Join("testdata", "lookup", "a", "aa", ".dnv"), filepath.Join("testdata", "lookup", ".dnv")},
		},
		{
			name:     "no .dnv file",
			path:     filepath.Join("testdata", "lookup", "b"),
			expected: []string{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rcs := config.LookupConfigs(tt.path)

			assert.Equal(t, tt.expected, rcs)
		})
	}
}
