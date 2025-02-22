package devenvstatus

import (
    "fmt"
    "os"
    "strings"

    "github.com/fatih/color"
    "github.com/go-git/go-git/v5"
    "github.com/go-git/go-git/v5/plumbing"
    "github.com/spf13/viper"
)

func GetDevStatus(showChange, showBranch, showAllBranches bool) {
    m := viper.GetString("mainPath")
    viperToCheck := viper.Get("ToCheck")
    t, _ := viperToCheck.([]interface{})

    for _, i := range t {
        toCheck := i.(map[string]interface{})

        if toCheck["is_repo"] == true {
            path := m + "/" + toCheck["path"].(string)
            getRepoStatus(path, showChange)
            continue
        }

        path := m + "/" + toCheck["path"].(string)
        dir, err := os.ReadDir(path)
        if err != nil {
            fmt.Printf("error during read Directory at path %v with error: %v", path, err)
        }

        for _, d := range dir {
            dirName := d.Name()
            fullPath := path + "/" + dirName
            getRepoStatus(fullPath, showChange)
            listLocalBranch(fullPath, showBranch, showAllBranches)
        }
    }
}

func getRepoStatus(repoPath string, verbose bool) {
    repo, err := git.PlainOpen(repoPath)

    if err != nil {
        //fmt.Printf("%v is not a git repository \n\n", repoPath)
        return
    }else{
        fmt.Printf("%v \n", repoPath)
    } 

    w, err := repo.Worktree()
    if err != nil {
        fmt.Printf("error during open git Worktree %v with error %v \n\n", repoPath, err)
    }

    status, err := w.Status()
    if err != nil {
        fmt.Printf("error during  git status Worktree %v with error %v \n\n", repoPath, err)
    }

    if status.IsClean() {
        green := color.New(color.FgGreen)
        boldGreen := green.Add(color.Bold)
        boldGreen.Println("The repo is clean")
        fmt.Println()
    } else {
        color.HiRed("The repo is not clean")
        fmt.Println()
    }
    if verbose == true {
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
        fmt.Println()
    }
}

func listLocalBranch(repoPath string, showBranch, showAllBranches bool) {

    repo, err := git.PlainOpen(repoPath)

    if err != nil {
        fmt.Printf("%v is not a git repository \n\n", repoPath)
        return
    }
    head, err := repo.Head()
    if err != nil {
        fmt.Println("error to get HEAD")
    }
    if showBranch{
        color.Green(fmt.Sprintf("Attach on branch: %v\n", head.Name().String()))
    }

    refs, err := repo.References()
    if err != nil {
        fmt.Println("error to get ref")
    }

    var remote []string
    var local []string

    refs.ForEach(func(r *plumbing.Reference) error {
        if r.Name().IsRemote() {
            b := strings.Split(string(r.Name().Short()), "/")
            if b[1] != "HEAD" {
                remote = append(remote, b[1])
            }
        } else if r.Name().IsBranch() {
            local = append(local, string(r.Name().Short()))
        }
        return nil
    })

    if showAllBranches {
        fmt.Printf("Remote: %v\n", remote)
        fmt.Printf("Local: %v\n", local)
    }
    fmt.Println()
}
