package utils

import (
    "fmt"
)

func ValueOf(info map[string]string, itemList []string) []string {
    retList := make([]string, len(itemList))
    for ind, item := range itemList {
        infoMsg, ok := info[item]
        if !ok {
            fmt.Printf("Item(%s) is not exist\n", item)
        }
        retList[ind] = infoMsg
    }
    return retList
}
