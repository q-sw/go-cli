package templates

import (
	"embed"
	_ "embed"
	"fmt"
	"os"
	"text/template"
)

type Readme struct {
	Title       string
	Description string
}

//go:embed files
var TemplateFile embed.FS

func RenderReadme(title, description string) {

	tmpl, err := template.ParseFS(TemplateFile, "files/readme.tmpl")
	if err != nil {
		fmt.Println(err)
	}

	var content Readme
	content.Title = title
	content.Description = description

	var tmplFile *os.File
	tmplFile, err = os.Create("README.md")
	err = tmpl.Execute(tmplFile, content)

}

func RenderMakefile(name string) {

	tmpl, err := template.ParseFS(TemplateFile, "files/makefile.tmpl")
	if err != nil {
		fmt.Println(err)
	}

	content := map[string]string{"Name": name}
	var tmplFile *os.File
	tmplFile, err = os.Create("makefile")
	err = tmpl.Execute(tmplFile, content)

}

func RenderMain(name string) {

	tmpl, err := template.ParseFS(TemplateFile, "files/main.tmpl")
	if err != nil {
		fmt.Println(err)
	}

	var tmplFile *os.File
	tmplFile, err = os.Create("main.go")
	err = tmpl.Execute(tmplFile, name)

}
