package main
import (
	"github.com/aquav3/file-manager/file-man"
	"github.com/aquav3/file-manager/utils"
	"log"
)
func main() {
	for {
		input, err := utils.GetInput(" > ")
		if err != nil {
			log.Fatal(err)
		}
		err = fileman.RunCommand(fileman.ParseCommand(input))
		if input == "exit" {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
	}
}
