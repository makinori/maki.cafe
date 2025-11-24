package data

import (
	"errors"
	"log/slog"
	"net/http"
	"strconv"
	"time"

	"github.com/makinori/foxlib/foxhttp"
	"github.com/robfig/cron/v3"
	"go.etcd.io/bbolt"
)

const (
	ipExpireDuration = time.Hour
)

var (
	// TODO: save this?
	ipExpireMap = map[string]time.Time{}
)

func init() {
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
	if Database == nil {
		slog.Error(
			"database not initialized for counter",
		)
		return 0
	}

	var value uint64

	err := Database.View(func(tx *bbolt.Tx) error {
		bucket := tx.Bucket(STATS_BUCKET)
		if bucket == nil {
			return errors.New("stats bucket not found")
		}

		counterBytes := bucket.Get([]byte("counter"))
		if len(counterBytes) == 0 {
			return nil
		}

		var err error
		value, err = strconv.ParseUint(string(counterBytes), 10, 64)
		if err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		slog.Error("failed to read counter", "err", err)
		return 0
	}

	return value
}

func AddOneToCounter(r *http.Request) {
	ip := foxhttp.GetIPAddress(r)

	expireTime, ok := ipExpireMap[ip]
	if ok && time.Now().Before(expireTime) {
		return
	}

	slog.Debug("counter +1 from " + ip)

	value := ReadCounter()
	value++

	data := []byte(strconv.FormatUint(value, 10))

	ipExpireMap[ip] = time.Now().Add(ipExpireDuration)

	err := Database.Update(func(tx *bbolt.Tx) error {
		bucket := tx.Bucket(STATS_BUCKET)
		if bucket == nil {
			return errors.New("stats bucket not found")
		}

		return bucket.Put([]byte("counter"), data)
	})

	if err != nil {
		slog.Error("failed to add one to counter", "err", err)
		return
	}
}
