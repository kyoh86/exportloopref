//go:build print_ast

package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		panic("use it like: print_ast <filename>")
	}
	filename := os.Args[1]
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, filename, nil, parser.Mode(0))

	if err != nil {
		panic(err)
	}

	for _, d := range f.Decls {
		ast.Print(fset, d)
		fmt.Println() // \n したい...
	}
}
