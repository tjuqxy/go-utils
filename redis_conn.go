package utils

import (
    "time"
    "strings"

    "github.com/garyburd/redigo/redis"
)

const (
    RedisConnTimeout  = time.Duration(time.Second)
    RedisReadTimeout  = time.Duration(5 * time.Second)
    RedisWriteTimeout = time.Duration(5 * time.Second)
)

func NewRedisConn(addr string) (redis.Conn, error) {
    return redis.DialTimeout("tcp", addr,
        RedisConnTimeout, RedisReadTimeout, RedisWriteTimeout)
}

func ReqRedis(addr, cmd string, param ...interface{}) (interface{}, error) {
    conn, err := NewRedisConn(addr)
    if err != nil {
        return nil, err
    }

    return conn.Do(cmd, param...)
}

func GetRedisConfig(addr, name string) ([]string, error) {
    return redis.Strings(ReqRedis(addr, "config", "get", name))
}

func GetClusterNodes(addr string, cmd string, param ...interface{}) ([]string, error) {
    info, err := ReqRedis(addr, cmd, param...)
    if err != nil {
        return nil, err
    }

    infoStr, err := String(info)
    if err != nil {
        return nil, err
    }

    strList := strings.Split(infoStr, "\n")
    for ind, str := range strList {
        strList[ind] = strings.Split(str, "\r")[0]
    }
    return strList, nil
}
