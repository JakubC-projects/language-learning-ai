package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConnection(t *testing.T) {
	_, _, err := Client.Collection("users").Add(context.Background(), map[string]interface{}{
		"first": "Ada",
		"last":  "Lovelace",
		"born":  1815,
	})
	assert.NoError(t, err)
}
