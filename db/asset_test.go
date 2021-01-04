package db

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestAssets(t *testing.T) {
	for _, asset := range Assets {
		require.True(t, IsValidAsset(asset.String()))
		require.Equal(t, ToAsset(asset.String()).String(), asset.String())
	}
	require.False(t, IsValidAsset("notarealasset"))
}
