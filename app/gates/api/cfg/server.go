package cfg

import (
	"os"
)

const (
	envGateAddr = "API_GATE_ADDR"
)

func GetAddr() string {
	addr, ok := os.LookupEnv(envGateAddr)
	if !ok {
		return "localhost:8080" // Default address.
	}
	return addr
}
