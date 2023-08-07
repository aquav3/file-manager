package fileman

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

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
    fmt.Println(content)

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

    if len(tokens) == 2 {
        result.Name = tokens[1]
    }
    result.Name = "Empty"

    return result
}

func RunCommand(cmd FullCommand) {
    switch cmd.Cmd {
        case Create:
            createFile(cmd.Name)
        case Read:
            readFile(cmd.Name)
        case Delete:
            deleteFile(cmd.Name)
        case Ls:
            ls()
        case Pwd:
            pwd()
        case Help:
            help()
    }     
}
