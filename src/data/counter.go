package data

import (
	"log/slog"
	"os"
	"strconv"
	"sync"
)

var counterMutex sync.RWMutex

func init() {
	os.Mkdir("data", 0755)
}

func ReadCounter() uint64 {
	counterMutex.RLock()
	data, err := os.ReadFile("data/counter.txt")
	counterMutex.RUnlock()

	if err != nil {
		return 0
	}

	value, err := strconv.ParseUint(string(data), 10, 64)
	if err != nil {
		slog.Error(
			"failed to parse counter.txt",
			"err", err, "data", string(data),
		)
		return 0
	}

	return value
}

func AddOneToCounter() {
	value := ReadCounter()
	value++

	counterMutex.Lock()
	os.WriteFile(
		"data/counter.txt", []byte(strconv.FormatUint(value, 10)), 0644,
	)
	counterMutex.Unlock()
}
