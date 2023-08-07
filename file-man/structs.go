package fileman

type Command string

const (
    Create Command = "create"
    Read Command = "read"
    Delete Command = "delete"
    Ls Command = "ls"
    Cd Command = "cd"
    Pwd Command = "pwd"
    Help Command = "help"
    Exit Command = "exit"
)

type FullCommand struct {
    Cmd Command
    Name string
}
