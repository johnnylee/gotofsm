package main

import (
	"os"
	"fmt"
	"strings"
	"go/token"
	"go/parser"
	"go/ast"
)

func printUsage() {
	fmt.Printf("Usage: %v [filename] [function]\n", os.Args[0])
	os.Exit(1)
}

func main() {
	if len(os.Args) != 3 {
		printUsage()
		return
	}

	filename := os.Args[1]
	funcName := os.Args[2]
	
	fset := token.NewFileSet()

	f, err := parser.ParseFile(fset, filename, nil, 0)
	if err != nil {
		panic(err)
	}

	inFunc := false

	curState := ""
	nextState := ""
	
	lines := make([]string, 0, 1)
	lines = append(lines, "digraph {")
	
	ast.Inspect(f, func(n ast.Node) bool {
		switch x := n.(type) {

		case *ast.FuncDecl:
			if x.Name.Name == funcName {
				inFunc = true
			} 

		case *ast.BranchStmt:
			if inFunc && x.Tok == token.GOTO {
				nextState = x.Label.Name
				lines = append(lines, 
					fmt.Sprintf("%v->%v;", curState, nextState))
			}

		case *ast.LabeledStmt:
			if inFunc {
				curState = x.Label.Name
			}
		}
		return true
	})
	
	lines = append(lines, "}")
	fmt.Println(strings.Join(lines, "\n"))
}