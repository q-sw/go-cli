package utils

import (
    "fmt"
    "os"
    "strings"
)

func GetCurrentDirName() string {
    currentPath, err := os.Getwd()
    if err != nil {
        fmt.Println(err)
    }

    pathSplit := strings.Split(currentPath, "/")
    return pathSplit[len(pathSplit)-1]

}
