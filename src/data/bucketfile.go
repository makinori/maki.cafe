package data

import (
	"errors"
	"io/fs"

	"go.etcd.io/bbolt"
)

// as far as i know these arent standard, but we can replicate os package

func (bucketName BucketFS) WriteFile(name string, data []byte) (fs.File, error) {
	err := Database.Update(func(tx *bbolt.Tx) error {
		bucket := tx.Bucket(bucketName)
		if bucket == nil {
			return errors.New("bucket doesnt exist: " + string(bucketName))
		}

		return bucket.Put([]byte(name), data)
	})

	if err != nil {
		return nil, err
	}

	// return bucketName.Open(name)

	return &BucketFile{
		name: name,
		data: data,
		ptr:  0,
	}, nil
}
