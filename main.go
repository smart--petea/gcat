package main

import (
    "os"
    "fmt"
    "bufio"
    "log"

    gcat "gcat/src"
)

func main() {
    fileNames, _ := gcat.ParseArgs(os.Args[1:])

    for _, fileName := range fileNames {
        file, err := os.Open(fileName)
        if err != nil {
            log.Fatal(err)
        }


        scanner := bufio.NewScanner(file)
        for scanner.Scan() {
            fmt.Println(scanner.Text())
        }

        if err := scanner.Err(); err != nil {
            log.Fatal(err)
        }

        defer file.Close()
    }
}

