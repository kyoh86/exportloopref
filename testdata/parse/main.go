package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
)

func main() {
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "./testdata/src/a/a.go", nil, parser.Mode(0))

	if err != nil {
		panic(err)
	}

	for _, d := range f.Decls {
		ast.Print(fset, d)
		fmt.Println()
	}
}
