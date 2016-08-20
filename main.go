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
	"github.com/ffurrer2/mini-go-static-analysis/regex"
	"flag"
	"github.com/ffurrer2/mini-go-static-analysis/algorithm"
)

func main() {

	filename, err := filepath.Abs("src/github.com/ffurrer2/mini-go-static-analysis/example/minigo/main.go")

	path := flag.String("p", filename, "path to code file")
	flag.Parse()

	fset := token.NewFileSet()
	file, err := parser.ParseFile(fset, *path, nil, parser.AllErrors) //|parser.Trace
	if err != nil {
		fmt.Println(err)
		return
	}

	// Init CleanUpVisitor
	v := new(visitor.CleanUpVisitor)
	// Run CleanUpVisitor
	ast.Walk(v, file)

	fmt.Println("######################")
	fmt.Println("File: " + *path)
	fmt.Println("######################")

	// Print simplified source code
	printer.Fprint(os.Stdout, fset, file)

	fmt.Println("######################")
	fmt.Println()

	// Init RegexBuilder
	v2 := new(algorithm.RegexBuilder)
	// Run RegexBuilder
	algorithm.Walk(v2, file)

	fmt.Println("########### Functions ###########")

	var v3 *algorithm.RegexPrinter
	for _, fun := range v2.Root.FuncList {
		v3 = new(algorithm.RegexPrinter)
		v3.Fset = fset
		regex.WalkRegex(v3, fun.Regex)
		fmt.Println(v3.Buffer.String())
	}
	fmt.Println()
	fmt.Println("############# Forks #############")
	for _, fork := range v2.Root.ForkList {
		if fork.Func != nil {
			v3 = new(algorithm.RegexPrinter)
			v3.Fset = fset
			regex.WalkRegex(v3, fork.Func.Regex)
			fmt.Println(v3.Buffer.String())
		}
	}
}
