package data

import (
	"errors"
	"log/slog"
	"os"

	"go.etcd.io/bbolt"
)

func migrateCounterTxtToDatabase() error {
	if Database == nil {
		return errors.New("database is nil")
	}

	oldCounterBytes, err := os.ReadFile("data/counter.txt")
	if err != nil {
		// file doesnt exist, already migrated
		return nil
	}

	err = Database.Update(func(tx *bbolt.Tx) error {
		statsBucket := tx.Bucket(STATS_BUCKET)
		if statsBucket == nil {
			return errors.New("failed to find stats bucket")
		}
		return statsBucket.Put([]byte("counter"), oldCounterBytes)
	})

	if err != nil {
		return err
	}

	slog.Info("migrated data/counter.txt to database. file deleted!")

	return os.Remove("data/counter.txt")
}
