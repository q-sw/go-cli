package utils

import (
	"fmt"
	"os"
	"path/filepath"
)

func CreateFile(path, name string) error {

	_, err := os.Create(filepath.Join(path, name))
	if err != nil {
		return fmt.Errorf("error during file create with error: %v", err)
	}
	return nil
}
