package main

import (
    "os"

    gcat "gcat/src"
)

func main() {
    gcat.NewApp(os.Args[1:]).Run()
}
