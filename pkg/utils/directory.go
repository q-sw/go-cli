package utils

import (
	"os"
	"path/filepath"
)

type Directory struct {
	DirName string
	DirPath string
}

func CreateDir(dir Directory) {
	os.MkdirAll(filepath.Join(dir.DirPath, dir.DirName), 0755)

}

func CreateDirs(dirs []Directory) {
	for _, d := range dirs {
		CreateDir(d)
	}

}
