package utils

import (
    "fmt"
)

func Pause() bool {
    //return true

    fmt.Println(GREEN("\nContinue?(y/n)"))
    s := ""
    fmt.Scan(&s)
    fmt.Println()
    if s == "n" {
        return false
    }
    return true
}
