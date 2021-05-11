package parse

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"

	"github.com/iancoleman/strcase"
)

type FieldType struct {
	Name string
	Type string
}

type ModelDesc struct {
	UperName   string      // struct name uperCamel
	LowerName  string      // struct name lowerCamel
	FieldTypes []FieldType // struct field name/type slice
}

func Model(sourcePath string) ModelDesc {
	var result ModelDesc
	// Create the AST by parsing src.
	fset := token.NewFileSet()
	// Set src = nil , only generate from .go file
	f, err := parser.ParseFile(fset, sourcePath, nil, 0)
	if err != nil {
		panic(err)
	}

	for _, node := range f.Decls {
		switch node.(type) {

		case *ast.GenDecl:
			genDecl := node.(*ast.GenDecl)
			for _, spec := range genDecl.Specs {
				switch spec.(type) {
				case *ast.TypeSpec:
					typeSpec := spec.(*ast.TypeSpec)

					fmt.Printf("Struct: name=%s\n", typeSpec.Name.Name)
					result.UperName = strcase.ToCamel(typeSpec.Name.Name)
					result.LowerName = strcase.ToLowerCamel(typeSpec.Name.Name)
					switch typeSpec.Type.(type) {
					case *ast.StructType:
						structType := typeSpec.Type.(*ast.StructType)
						for _, field := range structType.Fields.List {
							i := field.Type.(*ast.Ident)
							fieldType := i.Name

							for _, name := range field.Names {
								fmt.Printf("\tField: name=%s type=%s\n", name.Name, fieldType)
								result.FieldTypes = append(result.FieldTypes, FieldType{Name: name.Name, Type: fieldType})
							}

						}

					}
				}
			}
		}
	}

	return result
}
