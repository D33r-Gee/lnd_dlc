package dlc

import (
	"go.etcd.io/bbolt"
)

type DlcManager struct {
	DLCDB *bbolt.DB
}

// NewManager generates a new manager to add to the LND Node
func NewManager(dbPath string) (*DlcManager, error) {

	var mgr DlcManager
	err := mgr.InitDB(dbPath)
	if err != nil {
		return nil, err
	}

	return &mgr, nil
}
