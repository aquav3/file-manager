package utils

import (
    "os"
)

func ChangeDirectory(path string) error {
    return os.Chdir(path)
}

func GetCurrentWorkingDirectory() (string, error) {
    dir, err := os.Getwd()
    if err != nil {
        return "", err
    } 
    return dir, nil
}
