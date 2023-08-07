package main

import (
	"fmt"
	"log"
    "github.com/aquav3/file-manager/file-man"
	"github.com/aquav3/file-manager/utils"
)

 

func main() {
    var input string

    for {
        input, err := utils.GetInput(" > ")
        if err != nil {
            log.Fatal(err)
        }
        command := fileman.ParseCommand(input)

    }

}
