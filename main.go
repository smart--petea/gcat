package main

import (
    "os"
    "fmt"

    gcat "gcat/src"
)

func main() {
    files, params := gcat.ParseArgs(os.Args[1:])

    fmt.Println(files)
    fmt.Println(params)
}

