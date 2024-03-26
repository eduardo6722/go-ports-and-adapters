package handlers

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestJsonError(t *testing.T) {
	msg := "custom error"
	result := jsonError(msg)
	require.Equal(t, `{"message":"custom error"}`, string(result))
}
