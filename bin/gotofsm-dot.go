package main

import (
	"os"
	"fmt"
	"strings"
	"github.com/johnnylee/gotofsm"
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

	path := os.Args[1]
	funcName := os.Args[2]
	
	states, err := gotofsm.Analyze(path, funcName)
	
	if err != nil {
		fmt.Printf("Error: %v", err)
		return
	}

	lines := make([]string, 0, 1)
	lines = append(lines, "digraph {")
	
	for _, state := range states {
		for _, nextState := range state.Next {
			lines = append(
				lines, 
				fmt.Sprintf("%v->%v;", state.Name, nextState.Name))
		}
	}
	
	lines = append(lines, "}")
	fmt.Println(strings.Join(lines, "\n"))
}