package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"os"
)

func main() {
	// TODO cmd line parsing
	dirPath := os.Args[1]
	fset := token.NewFileSet()

	pkgs, err := parser.ParseDir(fset, dirPath, nil, 0)
	if err != nil {
		fmt.Println(err)
		return
	}

	for pkgName, pkg := range pkgs {
		fmt.Println("package", pkgName)
		for fileName, file := range pkg.Files {
			fmt.Println("  fileName", fileName)
			for _, decl := range file.Decls {
				switch declVal := decl.(type) {
				case *ast.FuncDecl:
					fmt.Printf("    FuncDecl: %+v\n", declVal)
				case *ast.GenDecl:
					fmt.Printf("    GenDecl: %+v\n", declVal)
					for _, spec := range declVal.Specs {
						switch specVal := spec.(type) {
						case *ast.ImportSpec:
							fmt.Printf("      ImportSpec: %+v\n", specVal)
						case *ast.ValueSpec:
							fmt.Printf("      ValueSpec: %+v\n", specVal)
						case *ast.TypeSpec:
							fmt.Printf("      TypeSpec: %+v\n", specVal)
							switch typeVal := specVal.Type.(type) {
							case *ast.StructType:
								fmt.Printf("        Struct: %+v\n", typeVal)
								for _, field := range typeVal.Fields.List {
									fmt.Printf("          Field: %+v\n", field)
								}
							default:
								fmt.Printf("        default: %+v\n", typeVal)
							}
						default:
							fmt.Printf("      default: %+v\n", specVal)
						}
					}
				default:
					fmt.Printf("   default: %+v\n", declVal)
				}
			}
		}
	}
}
