package utils

import (
    "net"
    "strings"
    "strconv"
)

func GetIpAndPortByAddr(addr string) (string, int) {
    spList := strings.Split(addr, ":")
    if len(spList) == 1 || spList[1] == "" {
        return "", 0
    }
    if spList[0] == "" {
        spList[0] = "127.0.0.1"
    }
    port, err := strconv.Atoi(spList[1])
    if err != nil {
        return "", 0
    }
    return spList[0], port
}

func GetHostByIp(ip string) string {
    hosts, err := net.LookupAddr(ip)
    if err != nil {
        return ""
    }
    return hosts[0]
}

func GetIpByHost(host string) string {
    return host
}
