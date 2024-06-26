package node

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/celestiaorg/celestia-node/core"
)

// MockStore provides mock in memory Store for testing purposes.
func MockStore(t *testing.T, cfg *Config) Store {
	t.Helper()
	store := NewMemStore()
	err := store.PutConfig(cfg)
	require.NoError(t, err)
	cstore, err := store.Core()
	require.NoError(t, err)
	err = cstore.PutConfig(core.MockConfig(t))
	require.NoError(t, err)
	return store
}
