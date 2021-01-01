package discord

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestConfig(t *testing.T) {
	t.Cleanup(func() {
		os.Remove("test-config.yml")
	})
	err := NewConfig("test-config.yml")
	require.NoError(t, err)
	cfg, err := LoadConfig("test-config.yml")
	require.NoError(t, err)
	require.Len(t, cfg.Watchers, 1)
	require.Equal(t, cfg.MainDiscordToken, "CHANGEME-MAIN")
	require.Equal(t, cfg.Watchers[0].DiscordToken, "CHANGEME-TOKEN")
	require.Equal(t, cfg.Watchers[0].Currency, "CHANGEME-CURRENCY")
}
