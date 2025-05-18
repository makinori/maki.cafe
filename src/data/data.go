package data

import (
	"sync"
	"time"

	"github.com/charmbracelet/log"
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
				log.Fatal("failed to parse cron spec", "err", err.Error())
				// exit program entirely
			}

			expiresAt := schedule.Next(time.Now())

			// get data
			cachedData.Data, err = cachedData.retrieve()
			if err != nil {
				log.Error(
					"failed to get data",
					"key", cachedData.Key, "err", err.Error(),
				)
				return
			}

			err = setCache(cachedData.Key, cachedData.Data, expiresAt)
			if err != nil {
				log.Error(
					"failed to set cache",
					"key", cachedData.Key, "err", err.Error(),
				)
			}
		}

		// try from cache
		err := getCache(cachedData.Key, &cachedData.Data)
		if err != nil {
			log.Infof(`fetching fresh "%s", starting cron`, cachedData.Key)
			getFreshCachedData()
		} else {
			log.Infof(`already cached "%s", starting cron`, cachedData.Key)

		}

		// setup cron
		c.AddFunc(cachedData.CronSpec, func() {
			getFreshCachedData()
		})

		// log.Println("starting cron for " + cachedData.Key)
	}()
}

func InitData() {
	c := cron.New()

	var wg sync.WaitGroup

	initCachedData(c, &wg, &Anilist)

	wg.Wait()

	c.Start()
}
