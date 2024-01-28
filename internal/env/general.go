package env

import "os"

func BaseURL() string {
	return os.Getenv("BASE_URL")
}

func GRPCServerPort() string {
	return os.Getenv("GRPC_ACCOUNT_SERVER_PORT")
}
