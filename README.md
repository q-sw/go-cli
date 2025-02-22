# go-cli
This is my own CLI write in GO with Cobra

## CLI

```plaintext
CLI is a collection of tools designed to improve the daily workflow of a DevOps engineer.
The CLI is structured using verbs and actions.

Usage:
  cli [command]

Available Commands:
    completion -> Generate the autocompletion script for the specified shell
    create -> Create project, files...
        go-project -> Create skaffold for Go project
            cli create go-project [flags]
                -d, --description string   [Global] Go project description
                -g, --enable-git           [Global] Activate Git
                -h, --help                 help for go-project
                -n, --name string          [Global] Go project name without github profile url
                -t, --type string          [Global] Go project type standard|api|cli (default "package")
        terraform -> Create skaffold for terraform
            cli create terrafrom [project|module|files] [flags]
                -g, --enable-git    [Project] Activate Git (git init) on your folder/project
                -h, --help          help for terraform
                -n, --name string   [Global] Name for your terraform module or project
                -p, --path string   [Global] Path where create your instance (default "./")
                -t, --type string   [Project] Project type, accept standard;env (default "standard")
    get -> Get information from your system or projet
        dev-status -> show the status of your dev environments
            -h, --help                help for dev-status
                --show-all-branches   [Global] Show local and remote branches
            -b, --show-branch         [Global] Show actual branch
            -c, --show-change         [Global] Show files changed
            -v, --verbose             [Global] Show details about repository status
    help -> Help about any command
```

## Init the project

### Install Cobra-cli
> Before installed tools with Go check your `$GOPATH` and check if it is in your `$PATH`  
> To get your environment information run `go env`

```shell
go install github.com/spf13/cobra-cli@latest
```

### Init the GO project
```shell
go mod init github.com/q-sw/go-qsw-cli
```

### Init the Cobra project
```shell
cobra-cli init qsw-cli --author qsw
```
At this point, cobra-cli creates a new folder named `qsw-cli`.  
To have a simpless project structure, I moved all folder content to the Root folder

```shell
mv qsw-cli/* .
```

## Install command line

```shell
make install
```

## Config File template
```yaml
---
mainPath: "/home/xxxx"
toCheck:
  - path: "dotfiles"
    is_repo: true
  - path: "dev/tmp"
    is_repo: false
  - path: "dev/github/public"
    is_repo: false
git_username: q-sw
git_email: xxxxx.xxxxxh@yyyy.yy
github_profile: github.com/q-sw
```
