package env

import "os"

func JWTSecret() string {
	return os.Getenv("SECRET_KEY")
}
