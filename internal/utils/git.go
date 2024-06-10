package utils

import (
	"log"
	"time"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/object"
	"github.com/spf13/viper"
)

func GitInit(path string) {

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
			Name:  viper.GetString("git_username"),
			Email: viper.GetString("git_email"),
			When:  time.Now(),
		},
	})

}
