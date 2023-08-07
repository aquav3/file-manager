package fileman

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"sync"
	"github.com/aquav3/file-manager/utils"
)



func createFile(filename string) error {
    _, err := os.Create(filename)
    if err != nil {
        return err
    }
    return nil
}

func readFile(filename string) error {
    content, err := os.ReadFile(filename)
    if err != nil {
        return err
    }
    fmt.Println(string(content))
    return nil
}

func deleteFile(filename string) error {
    err := os.Remove(filename)
    if err != nil {
        return err
    }
    return nil
}

func ls() error {
    files, err := ioutil.ReadDir(".")
    if err != nil {
        return err
    }
    for _, file := range files {
        fmt.Println(file.Name())
    }
    return nil
}

func cd(dir string) error {
    err := utils.ChangeDirectory(dir)
    return err
}

func pwd() {
    fmt.Println(utils.GetCurrentWorkingDirectory())
}

func help() {
    fmt.Println("create <filename>")
    fmt.Println("read <filename>")
    fmt.Println("delete <filename>")
    fmt.Println("ls")
    fmt.Println("cd <directory>")
    fmt.Println("pwd")
    fmt.Println("help")
}

func ParseCommand(cmd string) FullCommand {
    result := FullCommand {
        Cmd: Create,
        Name: "",
    }
    tokens := strings.Split(cmd, " ")
    switch tokens[0] {
        case "create":
            result.Cmd = Create
        case "read":
            result.Cmd = Read
        case "delete":
            result.Cmd = Delete
        case "ls":
            result.Cmd = Ls
        case "cd":
            result.Cmd = Cd
        case "pwd":
            result.Cmd = Pwd
        case "help":
            result.Cmd = Help
    }
    result.Name = "Empty"
    if len(tokens) == 2 {
        result.Name = tokens[1]
    }
    return result
}

func RunCommand(cmd FullCommand) error {
    var wg sync.WaitGroup
    var err error
    wg.Add(1)
    switch cmd.Cmd {
        case Create:
            go func() {
                defer wg.Done()
                err = createFile(cmd.Name)
            }()
        case Read:
            go func() {
                defer wg.Done()
                err = readFile(cmd.Name)
            }()
        case Delete:
            go func() {
                defer wg.Done()
                err = deleteFile(cmd.Name)
            }()
        case Ls:
            go func() {
                defer wg.Done()
                err = ls()
            }()
        case Pwd:
            go func() {
                defer wg.Done()
                pwd()
            }()
        case Help:
            go func() {
                defer wg.Done()
                help()
            }()
    }   
    return err 
}
