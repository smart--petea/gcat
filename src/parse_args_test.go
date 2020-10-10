package gcat

import (
    "testing"
    "fmt"
)

func TestParseArgs(t *testing.T) {
    var inputArgs []string
    var outputFiles []string
    var outputParams *Params

    inputArgs = []string{"a.txt", "b.txt"}
    outputFiles, outputParams = ParseArgs(inputArgs)
    if len(outputFiles) != 2 {
        t.Errorf("inputArgs=%v outputFiles=%v", inputArgs, outputFiles)
    }

    fmt.Println(outputParams)
}
