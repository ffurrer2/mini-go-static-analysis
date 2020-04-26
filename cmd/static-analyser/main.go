// SPDX-License-Identifier: MIT
package main

import (
	"flag"
	"fmt"
	"github.com/ffurrer2/mini-go-static-analysis/pkg/algorithm"
	"github.com/ffurrer2/mini-go-static-analysis/pkg/regex"
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
