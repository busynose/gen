package main

import (
	"gen"
	"gen/parse"
	"gen/template"
	"os"
)

const (
	defaultFileMode = 0o600
)

const (
	sourcePath  = "model/demo.go"
	packageName = "nextMaker"
)

func main() {
	f, err := os.OpenFile("gen.go", os.O_CREATE|os.O_WRONLY, defaultFileMode)
	if err != nil {
		return
	}
	modelDesc := parse.Model(sourcePath)
	gen.Repo(template.RepoTemplate, modelDesc, packageName, f)
}
