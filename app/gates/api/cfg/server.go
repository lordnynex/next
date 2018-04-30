package cfg

import "os"

const (
	envAddr = "API_GATE_ADDR"
)

func GetAddr() string {
	addr, ok := os.LookupEnv(envAddr)
	if !ok {
		panic(envAddr + " is not defined")
	}
	return addr
}
