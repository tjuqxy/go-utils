package utils

import (
	"fmt"
	"strconv"
)

func ShowRedisInfo(addr, info string) ([]string, error) {
	req := fmt.Sprintf("*1\r\n$%d\r\n%s\r\n", len(info), info)
	return callCmd(addr, []byte(req))
}

func CallRedis(addr, cmd string, param ...interface{}) ([]string, error) {
	req := ""
	req += "*" + strconv.Itoa(len(param) + 1) + "\r\n"
	req += "$" + strconv.Itoa(len(cmd)) + "\r\n"
	req += cmd + "\r\n"
	for _, p := range param {
		req += "$" + strconv.Itoa(len(p.(string))) + "\r\n"
		req += p.(string) + "\r\n"
	}
	return callCmd(addr, []byte(req))
}
