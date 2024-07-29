# go-qsw-cli
This is my own CLI write in GO with Cobra

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
