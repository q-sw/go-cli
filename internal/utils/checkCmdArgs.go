package utils

import (
    "fmt"
)

func CheckArgs(validArgs []string, cmdArgs []string) error {
    for _, a := range cmdArgs {
        for _, v := range validArgs {
            if v == a {
                return nil
            }
        }
        return fmt.Errorf("Invalid args: select one of %v", validArgs)
    }
    return nil
}
