package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
)

func GetRootDir() string {
	_, goFilePath, _, ok := runtime.Caller(0)
	if !ok {
		fmt.Println("failed to get caller file")
		os.Exit(1)
	}
	return filepath.Join(filepath.Dir(goFilePath), "..")
}
