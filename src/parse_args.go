package gcat

import (
)

func ParseArgs(args []string) (files []string, params *Params) {
    files = make([]string, 0)
    params = &Params{}

    for _, arg := range args {
        switch arg {
        case "-A":
            params.ShowEnd = true
        case "--show-all":
            params.ShowEnd = true
        case "-b":
            params.NumberNonBlank = true
        case "--number-nonblank":
            params.NumberNonBlank = true
        case "-e":
            params.ShowEnd = true
        case "-E":
            params.ShowEnd = true
        case "--show-ends":
            params.ShowEnd = true
        case "-n":
            params.NumberAllLines = true
        case "--number":
            params.NumberAllLines = true
        case "-s":
            params.SqueezeBlank = true
        case "--squeeze-blank":
            params.SqueezeBlank = true
        }

        files = append(files, arg)
    }

    return
}
