package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoadConfig(t *testing.T) {
	tests := []struct {
		name      string
		envVars   map[string]string
		want      *Config
		wantError bool
	}{
		{
			name: "success",
			envVars: map[string]string{
				"JWT_SECRET": "test-secret",
				"DB_URL":     "postgres://user:password@localhost:5432/dbname",
			},
			want: &Config{
				JwtSecret: []byte("test-secret"),
				DbURL:     "postgres://user:password@localhost:5432/dbname",
			},
			wantError: false,
		},
		{
			name:      "missing required env var",
			envVars:   map[string]string{},
			want:      nil,
			wantError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			// Set test environment variables
			for k, v := range tt.envVars {
				t.Setenv(k, v)
			}

			// Run test
			got, err := LoadConfig()

			if tt.wantError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.want, got)
			}
		})
	}
}
