package data

import (
	"log/slog"
	"os"
	"sync"
	"time"

	"github.com/robfig/cron/v3"
)

type cachedData[T any] struct {
	Key      string
	CronSpec string
	retrieve func() (T, error)
	Data     T
}

func initCachedData[T any](
	c *cron.Cron, wg *sync.WaitGroup, cachedData *cachedData[T],
) {
	wg.Add(1)
	go func() {
		defer wg.Done()

		getFreshCachedData := func() {
			// parse cron spec so we can get an expire time
			schedule, err := cron.ParseStandard(cachedData.CronSpec)
			if err != nil {
				slog.Error("failed to parse cron spec", "err", err.Error())
				os.Exit(1)
				// exit program entirely
			}

			expiresAt := schedule.Next(time.Now())

			// get data
			cachedData.Data, err = cachedData.retrieve()
			if err != nil {
				slog.Error(
					"failed to get data",
					"key", cachedData.Key, "err", err.Error(),
				)
				return
			}

			err = setCache(cachedData.Key, cachedData.Data, expiresAt)
			if err != nil {
				slog.Error(
					"failed to set cache",
					"key", cachedData.Key, "err", err.Error(),
				)
			}
		}

		// try from cache
		err := getCache(cachedData.Key, &cachedData.Data)
		if err != nil {
			slog.Info("fetching fresh", "key", cachedData.Key)
			getFreshCachedData()
		} else {
			slog.Info("already cached", "key", cachedData.Key)

		}

		// setup cron
		c.AddFunc(cachedData.CronSpec, func() {
			getFreshCachedData()
		})

		// slog.Println("starting cron for " + cachedData.Key)
	}()
}

func InitData() {
	c := cron.New()

	var wg sync.WaitGroup

	initCachedData(c, &wg, &Anilist)

	wg.Wait()

	c.Start()
}
