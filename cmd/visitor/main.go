package main

import (
	"flag"
	"fmt"
	"github.com/ffurrer2/mini-go-static-analysis/pkg/visitor"
	"go/ast"
	"go/parser"
	"go/printer"
	"go/token"
	"os"
	"path/filepath"
)

func main() {

	filename, err := filepath.Abs("pkg/examples/minigo/main.go")

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
