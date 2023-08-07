package main

import (
	"fmt"
	"log"

	"github.com/aquav3/file-manager/utils"
)

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

func main() {

}
