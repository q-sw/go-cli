package utils

import (
	"log"
	"os"
	"path/filepath"
)

type File struct {
	Filename string
	Path     string
}

func CreateFile(file File) os.File {
	if file.Path == "" {
		file.Path = "./"
	}
	f, err := os.Create(filepath.Join(file.Path, file.Filename))
	if err != nil {
		log.Fatal(err)
	}

	return *f
}

func CreateFiles(files []File) {
	for _, f := range files {
		CreateFile(f)
	}

}
