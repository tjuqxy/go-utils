package utils

import (
	"fmt"
)

func SetCache(addr, key, flag, exptime, value string) ([]string, error) {
	req := fmt.Sprintf("set %s %s %s %d\r\n%s\r\n", key, flag, exptime, len(value), value)
	return callCmd(addr, []byte(req))
}

func AddCache(addr, key, flag, exptime, value string) ([]string, error) {
	req := fmt.Sprintf("add %s %s %s %d\r\n%s\r\n", key, flag, exptime, len(value), value)
	return callCmd(addr, []byte(req))
}

func GetCache(addr, key string) ([]string, error) {
	req := fmt.Sprintf("get %s\r\n", key)
	return callCmd(addr, []byte(req))
}

func DelCache(addr, key string) ([]string, error) {
	req := fmt.Sprintf("delete %s\r\n", key)
	return callCmd(addr, []byte(req))
}
