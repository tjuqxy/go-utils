package utils

import (
    "github.com/fatih/color"
)

var (
    GREEN  = color.New(color.FgGreen).SprintFunc()
    YELLOW = color.New(color.FgYellow).SprintFunc()
    CYAN   = color.New(color.FgCyan).SprintFunc()
)
