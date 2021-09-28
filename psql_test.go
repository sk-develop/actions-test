package main

import (
	"testing"

	_ "github.com/lib/pq"
	"github.com/stretchr/testify/assert"
)

func Test_psqlClient_ping(t *testing.T) {
	tests := []struct {
		name       string
		psqlClient *psqlClient
		wantErr    string
	}{
		{
			name:       "success to ping",
			psqlClient: connPsql(),
			wantErr:    "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.psqlClient.ping()
			if tt.wantErr != "" {
				assert.EqualError(t, err, tt.wantErr)
				return
			}
			assert.NoError(t, err)
		})
	}
}
