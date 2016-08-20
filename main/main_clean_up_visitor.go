package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"path/filepath"
	"github.com/ffurrer2/mini-go-static-analysis/visitor"
	"os"
	"go/printer"
	"flag"
)

func main() {

	filename, err := filepath.Abs("src/github.com/ffurrer2/mini-go-static-analysis/example/minigo/main.go")

	path := flag.String("p", filename, "path to code file")
	flag.Parse()

	fset := token.NewFileSet()
	file, err := parser.ParseFile(fset, *path, nil, parser.AllErrors)
	if err != nil {
		fmt.Println(err)
		return
	}

	v := new(visitor.CleanUpVisitor)
	ast.Walk(v, file)

	fmt.Println("######################")
	printer.Fprint(os.Stdout, fset, file)
	fmt.Println("######################")
}
