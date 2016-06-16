package utils

import (
    "fmt"
    "testing"
)

func TestRedisConfigGet(t *testing.T) {
    fmt.Println(GetRedisConfig(":6379", "maxmemory-policy"))
}
