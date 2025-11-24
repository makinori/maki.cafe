package data

import (
	"errors"

	"github.com/makinori/foxlib/foxcache"
	"go.etcd.io/bbolt"
)

var (
	STATS_BUCKET       = []byte("stats")
	CACHE_BUCKET       = []byte("cache")
	CACHE_FILES_BUCKET = []byte("cache_files")

	ALL_BUCKETS = [][]byte{
		STATS_BUCKET, CACHE_BUCKET, CACHE_FILES_BUCKET,
	}

	Database *bbolt.DB

	CacheFilesFS = BucketFS(CACHE_FILES_BUCKET)
)

func Init() error {
	var err error
	Database, err = bbolt.Open("data.db", 0644, nil)
	if err != nil {
		return errors.New("failed to open database: " + err.Error())
	}

	// ensure buckets
	err = Database.Update(func(tx *bbolt.Tx) error {
		for _, name := range ALL_BUCKETS {
			_, err := tx.CreateBucketIfNotExists(name)
			if err != nil {
				return err
			}
		}
		return nil
	})
	if err != nil {
		return errors.New("failed to create bucket: " + err.Error())
	}

	// migrations
	err = migrateCounterTxtToDatabase()
	if err != nil {
		return errors.New("failed to migrate: " + err.Error())
	}

	// init cache
	err = foxcache.Init(Database, CACHE_BUCKET, []foxcache.DataInterface{
		&Anilist, &Squirrels,
	})
	if err != nil {
		return errors.New("failed to init cache: " + err.Error())
	}

	return nil
}
