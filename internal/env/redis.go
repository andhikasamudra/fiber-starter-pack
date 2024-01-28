package env

import "os"

func RedisURL() string {
	return os.Getenv("REDIS_URL")
}
