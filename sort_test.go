package utils

import (
    "fmt"
    "testing"
)

func TestSortMapListById(t *testing.T) {
    mapList := make([]map[string]string, 5)
    mapList[0] = map[string]string {
        "id" : "c",
    }
    mapList[1] = map[string]string {
        "id" : "b",
    }
    mapList[2] = map[string]string {
        "id" : "a",
    }
    mapList[3] = map[string]string {
        "id" : "e",
    }
    mapList[4] = map[string]string {
        "id" : "f",
    }
    fmt.Println(mapList)
    SortMapListById("id", mapList)
    fmt.Println(mapList)
}
