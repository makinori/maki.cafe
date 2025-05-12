package data

import (
	"log"
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

func initCachedData[T any](c *cron.Cron, wg *sync.WaitGroup, cachedData *cachedData[T]) {
	wg.Add(1)
	go func() {
		defer wg.Done()

		getFreshCachedData := func() {
			// parse cron spec so we can get an expire time
			schedule, err := cron.ParseStandard(cachedData.CronSpec)
			if err != nil {
				log.Fatalln("failed to parse cron spec: " + err.Error())
				// exit program entirely
			}

			expiresAt := schedule.Next(time.Now())

			// get data
			cachedData.Data, err = cachedData.retrieve()
			if err != nil {
				log.Printf(
					"failed to get %s data: %s\n", cachedData.Key, err.Error(),
				)
				return
			}

			err = setCache(cachedData.Key, cachedData.Data, expiresAt)
			if err != nil {
				log.Printf(
					"failed to set %s cache: %s\n", cachedData.Key, err.Error(),
				)
			}
		}

		// try from cache
		err := getCache(cachedData.Key, &cachedData.Data)
		if err != nil {

			log.Println(`fetching fresh "` + cachedData.Key + `", starting cron`)
			getFreshCachedData()
		} else {
			log.Println(`already cached "` + cachedData.Key + `", starting cron`)

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
