package data

import (
	"log/slog"
	"net/http"
	"os"
	"strconv"
	"sync"
	"time"

	"github.com/robfig/cron/v3"
	"maki.cafe/src/util"
)

const (
	ipExpireDuration = time.Hour
)

var (
	counterMutex sync.RWMutex
	// save this?
	ipExpireMap = map[string]time.Time{}
)

func init() {
	os.Mkdir("data", 0755)

	// reap once an hour
	c := cron.New()
	c.AddFunc("0 * * * *", func() {
		slog.Debug("reaping expired ips")
		for ip, expire := range ipExpireMap {
			if time.Now().After(expire) {
				slog.Debug("expired", "ip", ip)
				delete(ipExpireMap, ip)
			}
		}
	})
	c.Start()
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

func AddOneToCounter(r *http.Request) {
	ip := util.HTTPGetIPAddress(r)

	expireTime, ok := ipExpireMap[ip]
	if ok && time.Now().Before(expireTime) {
		return
	}

	slog.Debug("counter +1 from " + ip)

	value := ReadCounter()
	value++

	ipExpireMap[ip] = time.Now().Add(ipExpireDuration)

	counterMutex.Lock()
	os.WriteFile(
		"data/counter.txt", []byte(strconv.FormatUint(value, 10)), 0644,
	)
	counterMutex.Unlock()
}
