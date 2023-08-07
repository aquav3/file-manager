package utils

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func GetInput(prompt string) (string, error) {
    reader := bufio.NewReader(os.Stdin)    
    fmt.Print(prompt)
    result, err := reader.ReadString('\n')
    if err != nil {
        return "", err
    }
    return strings.Trim(result, "\n"), nil
}
