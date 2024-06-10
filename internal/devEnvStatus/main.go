package devenvstatus

import (
	"fmt"
	"os"

	"github.com/fatih/color"
	"github.com/go-git/go-git/v5"
	"github.com/spf13/viper"
)

func GetDevStatus() {

	m := viper.GetString("mainPath")
	toCheck := viper.GetStringSlice("toCheck")
	for _, c := range toCheck {

		path := m + "/" + c
		dir, err := os.ReadDir(path)
		if err != nil {
			fmt.Println(err)
		}

		for _, d := range dir {
			dirName := d.Name()
			fullPath := path + "/" + dirName

			fmt.Println(fullPath)
			repo, err := git.PlainOpen(fullPath)
			if err != nil {
				fmt.Println(err)
				continue
			}
			w, _ := repo.Worktree()
			status, _ := w.Status()
			if status.IsClean() {
				green := color.New(color.FgGreen)
				boldGreen := green.Add(color.Bold)
				boldGreen.Println("The repo is clean")
			} else {
				color.HiRed("The repo is not clean")
			}
			for k, v := range status {
				switch {
				case v.Worktree == git.StatusCode('?'):
					r := string(v.Worktree) + " " + k
					color.Yellow(r)
				case v.Worktree == git.StatusCode('M'):
					r := string(v.Worktree) + " " + k
					color.Cyan(r)
				case v.Worktree == git.StatusCode('A'):
					r := string(v.Worktree) + " " + k
					color.Green(r)
				case v.Worktree == git.StatusCode('D'):
					r := string(v.Worktree) + " " + k
					color.Red(r)
				case v.Worktree == git.StatusCode(' '):
					r := string(v.Worktree) + " " + k
					color.Blue(r)
				}
			}
		}
	}
}
