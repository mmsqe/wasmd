package wasmtesting

import (
	storetypes "cosmossdk.io/store/types"
)

// MockCommitMultiStore mock with a CacheMultiStore to capture commits
type MockCommitMultiStore struct {
	storetypes.CommitMultiStore
	Committed []bool
}

func (m *MockCommitMultiStore) CacheMultiStore() storetypes.CacheMultiStore {
	m.Committed = append(m.Committed, false)
	return &mockCMS{m, &m.Committed[len(m.Committed)-1]}
}

type mockCMS struct {
	storetypes.CommitMultiStore
	committed *bool
}

// RunAtomic implements the CacheMultiStore interface.
func (m *mockCMS) RunAtomic(fn func(storetypes.CacheMultiStore) error) error {
	return fn(m)
}

func (m *mockCMS) Write() {
	*m.committed = true
}

func (m *mockCMS) Copy() storetypes.CacheMultiStore {
	return m
}
