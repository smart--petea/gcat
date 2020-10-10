package gcat

import (
    "os"
    "fmt"
    "bufio"
)

type App struct {
    Files []string
    Params *Params
}

func NewApp(args []string) *App {
    files,  params := ParseArgs(args)

    return &App{
        Files: files,
        Params: params,
    }
}

func (app *App) GetFile(fileName string) (*os.File, error) {
    if fileName != "-" {
        return os.Open(fileName)
    }

    stat, err := os.Stdin.Stat()
    if err != nil {
        return nil, err
    }

    if (stat.Mode() & os.ModeCharDevice) == 0 {
        return os.Stdin, nil
    }

    //stdin is terminal. no data
    return nil, nil
}


func (app *App) Run() {
    for _, file := range app.Files {
        app.PrintFile(file)
    }
}

func (app *App) PrintFile(fileName string) error {
    file, err := app.GetFile(fileName)
    if err != nil || file == nil {
        return err
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        fmt.Println(scanner.Text())
    }

    if err := scanner.Err(); err != nil {
        return err
    }

    return nil
}
