package util

import "os"

var (
	ENV_PORT               = os.Getenv("PORT")
	_, ENV_IS_DEV          = os.LookupEnv("DEV")
	_, ENV_PLAUSIBLE_DEBUG = os.LookupEnv("PLAUSIBLE_DEBUG")
)
