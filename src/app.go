package gcat

import (
    "os"
    "fmt"
    "bufio"
    "strings"
)

type App struct {
    Files []string
    Params *Params
    CurrentLineNumber uint32
    PrecedentLine string
    NotFirstLine bool
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
        app.PrintLine(scanner.Text())
    }

    if err := scanner.Err(); err != nil {
        return err
    }

    return nil
}

func (app *App) PrintLine(line string) {
    if app.Params.SqueezeBlank && app.NotFirstLine && app.PrecedentLine == line {
        return
    }

    app.NotFirstLine = true
    app.PrecedentLine = line

    prefix := "\t"
    if app.Params.NumberNonBlank {
        if len(strings.TrimSpace(line)) > 0 {
            app.CurrentLineNumber = app.CurrentLineNumber + 1
            prefix = fmt.Sprintf("%s%d ", prefix, app.CurrentLineNumber)
        }
    } else if app.Params.NumberAllLines {
            app.CurrentLineNumber = app.CurrentLineNumber + 1
            prefix = fmt.Sprintf("%s%d ", prefix, app.CurrentLineNumber)
    }

    var sufix string
    if app.Params.ShowEnd {
        sufix = "$"
    }

    fmt.Printf("%s%s%s\n", prefix, line, sufix)
}
