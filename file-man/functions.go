package fileman

import "strings"

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
        case "exit":
            result.Cmd = Exit
    }

    if len(tokens) == 2 {
        result.Name = tokens[1]
    }
    result.Name = "Empty"

    return result

}
