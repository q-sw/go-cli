package utils

import (
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/object"
	"log"
	"time"
)

func InitRepo(path string, files []string) {
	git.PlainInit(path, false)

	r, err := git.PlainOpen(path)
	if err != nil {
		log.Fatal("Error when open git repo")
	}
	w, err := r.Worktree()
	if err != nil {
		log.Fatal("Error when set worktree")
	}

	w.AddGlob("*")
	w.Commit("Init repository", &git.CommitOptions{
		Author: &object.Signature{
			Name:  "q-sw",
			Email: "quentin.swiech@gmail.com",
			When:  time.Now(),
		},
	})

}
