package handler

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMakeErrorJson(t *testing.T) {
	msg := "Error JSON"
	result := jsonError(msg)
	require.Equal(t, string([]byte(`{"message":"Error JSON"}`)), string(result))
}
