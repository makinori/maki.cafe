package data

import (
	"errors"
	"io"
	"io/fs"
	"time"

	"go.etcd.io/bbolt"
)

// bucket key
type BucketFS []byte

func (bucketName BucketFS) Open(name string) (fs.File, error) {
	file := BucketFile{name: name, ptr: 0}

	// would prefer not to read on open
	// but cant figure out how not to with bolt

	err := Database.View(func(tx *bbolt.Tx) error {
		bucket := tx.Bucket(bucketName)
		if bucket == nil {
			return errors.New("bucket doesnt exist: " + string(bucketName))
		}

		data := bucket.Get([]byte(name))
		if len(data) == 0 {
			return fs.ErrNotExist
		}

		file.data = make([]byte, len(data))
		copy(file.data, data)

		return nil
	})

	if err != nil {
		return nil, err
	}

	return &file, nil
}

type BucketFile struct {
	name string
	data []byte
	ptr  int
}

func (file *BucketFile) Stat() (fs.FileInfo, error) {
	return BucketFileInfo(BucketFileInfo{file}), nil
}

func (file *BucketFile) Read(out []byte) (int, error) {
	need := len(out)
	if need == 0 {
		// need none
		return 0, nil
	}

	if file.data == nil {
		return 0, fs.ErrClosed
	}

	size := len(file.data)

	if file.ptr >= size {
		// none left
		return 0, io.EOF
	}

	to := min(file.ptr+need, size)

	n := copy(out, file.data[file.ptr:to])
	file.ptr += n

	return n, nil
}

func (file *BucketFile) Close() error {
	file.data = nil
	file.ptr = 0
	return nil
}

type BucketFileInfo struct {
	*BucketFile
}

func (info BucketFileInfo) Name() string {
	return info.name
}

func (info BucketFileInfo) Size() int64 {
	// will be 0 when closed
	return int64(len(info.data))
}

func (info BucketFileInfo) Mode() fs.FileMode {
	return 0644
}

func (info BucketFileInfo) ModTime() time.Time {
	return time.Unix(0, 0)
}

func (info BucketFileInfo) IsDir() bool {
	return false
}

func (info BucketFileInfo) Sys() any {
	return nil
}
