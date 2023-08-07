package utils

import (
    "os"
)

func ChangeDirectory(path string) error {
    return os.Chdir(path)
}

func GetCurrentWorkingDirectory() string {
    dir, err := os.Getwd()
    if err != nil {
        return ""
    } 
    return dir
}
