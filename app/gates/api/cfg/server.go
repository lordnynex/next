package cfg

import (
	"os"
)

const (
	defaultAddr = "localhost:8080"

	envGateAddr = "API_GATE_ADDR"
)

func GetAddr() string {
	addr, ok := os.LookupEnv(envGateAddr)
	if !ok {
		return defaultAddr
	}
	return addr
}
