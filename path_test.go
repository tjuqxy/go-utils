package utils

import (
    "fmt"
    "testing"
)

func TestGetPath(t *testing.T) {
    fmt.Println(GetPath("."))
    fmt.Println(GetPath(".."))
    fmt.Println(GetPath("..."))

    fmt.Println(GetPath("./"))
    fmt.Println(GetPath("../"))
    fmt.Println(GetPath(".../"))

    fmt.Println(GetPath("./a.b.c."))
    fmt.Println(GetPath("./../a.b.c."))
    fmt.Println(GetPath("./a.b/../c"))

    fmt.Println(GetPath("./a/b/c/d/e/f"))
}
