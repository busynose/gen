package gen

import (
	"gen/parse"
	"html/template"
	"io"
)

type Data struct {
	parse.ModelDesc
	PackageName string
}

func buildData(modelDesc parse.ModelDesc, packageName string) Data {
	var result Data
	result.UperName = modelDesc.UperName
	result.LowerName = modelDesc.LowerName
	result.FieldTypes = modelDesc.FieldTypes
	result.PackageName = packageName
	return result
}

func Repo(templateRepo string, modelDesc parse.ModelDesc, packageName string, wr io.Writer) {
	t, _ := template.New("repo").Parse(templateRepo)
	_ = t.Execute(wr, buildData(modelDesc, packageName))
}
