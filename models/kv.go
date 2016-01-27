package models

import (
	"fmt"
	"os"

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
		os.Exit(1)
	}

	kv = _kv
}

// GetKv Get Store
func GetKv() store.Store {
	return kv
}
