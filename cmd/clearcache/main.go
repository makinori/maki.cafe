package main

import (
	"fmt"
	"os"
	"path/filepath"

	"go.etcd.io/bbolt"
	"maki.cafe/cmd"
)

func exists(path string) bool {
	file, err := os.Open(path)
	exists := !os.IsNotExist(err)
	file.Close()
	return exists
}

func main() {
	databasePath := filepath.Join(cmd.GetRootDir(), "data.db")

	fmt.Println("opening: " + databasePath)

	if !exists(databasePath) {
		fmt.Println("database doesnt exist")
		os.Exit(1)
	}

	db, err := bbolt.Open(databasePath, 0644, nil)
	if err != nil {
		fmt.Println("failed to open database: " + err.Error())
		os.Exit(1)
	}

	err = db.Update(func(tx *bbolt.Tx) error {
		tx.DeleteBucket([]byte("cache"))
		fmt.Println("deleted bucket: cache")

		tx.DeleteBucket([]byte("cache_files"))
		fmt.Println("deleted bucket: cache_files")

		return nil
	})

	if err != nil {
		fmt.Println("failed to clear cache: " + err.Error())
		os.Exit(1)
	}
}
