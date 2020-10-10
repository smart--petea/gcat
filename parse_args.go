package gcat

func ParseArgs(args []string) (files []string, params *Params) {
    files = make([]string, 0)
    params = &Params{}

    for _, arg := range os.Args[1:] {
        switch arg {
        case "-A":
            params.ShowNonprinting = true
            params.ShowEnd = true
            params.ShowTabs = true
        case "--show-all":
            params.ShowNonprinting = true
            params.ShowEnd = true
            params.ShowTabs = true
        case "-b":
            params.NumberNonBlank = true
        case "--number-nonblank":
            params.NumberNonBlank = true
        case "-e":
            params.ShowNonprinting = true
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
            params.SqeezeBlank = true
        case "--squeze-blank":
            params.SqeezeBlank = true
        case "-t":
            params.ShowNonprinting = true
            params.ShowTabs = true
        case "-T":
            params.ShowTabs = true
        case "--show-tabs":
            params.ShowTabs = true
        case "-v":
            params.ShowNonprinting = true
        case "--show-nonprinting":
            params.ShowNonprinting = true
        case "--help":
            params.Help = true
        case "--version":
            params.Version = true
        }

        files = append(files, arg)
    }

    return
}
