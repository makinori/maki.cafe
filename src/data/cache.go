package data

import (
	"os"
	"path"
)

func writeCachePublic(filename string, data []byte) error {
	filePath := "cache/public/" + filename

	os.MkdirAll(path.Dir(filePath), 0755)
	err := os.WriteFile(filePath, data, 0644)
	if err != nil {
		return err
	}

	return nil
}
