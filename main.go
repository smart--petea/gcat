package main

import (
    "os"
    "fmt"

    "github.com/smart--petea/gcat"
)

func main() {
    files, params := parseArgs(os.Args[1:])

    fmt.Println(files)
    fmt.Println(params)
}

