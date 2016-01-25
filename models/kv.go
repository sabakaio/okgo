package models

import (
	"fmt"

	"github.com/docker/libkv"
	"github.com/docker/libkv/store"
	"github.com/docker/libkv/store/boltdb"
)

var kv store.Store

func init() {
	// Register store to libkv
	boltdb.Register()

	_kv, err := libkv.NewStore(
		store.BOLTDB,
		[]string{"/tmp/boltdbtest"},
		&store.Config{Bucket: "boltDBTest"},
	)

	if err != nil {
		fmt.Println(err)
		return
	}

	kv = _kv
}

// GetKv Get Store
func GetKv() store.Store {
	return kv
}
