package gotofsm

import (
	"go/token"
	"go/parser"
	"go/ast"
)

// A State has a name and a list of state transitions. 
type State struct {
	Name string
	Next []*State
}

type analyzer struct {
	funcName string // Function name to analyze.

	stateMap  map[string]*State // Map from state name to state.
	stateList []*State          // A list of all states.
	state     *State            // The current state.

	inFunc        bool // True when we're in the correct function.
}

// Analyze the given file and function, returning a list of states. 
func Analyze(path, funcName string) ([]*State, error) {

	a := new(analyzer)
	a.stateMap = make(map[string]*State)
	a.funcName = funcName

	fset := token.NewFileSet()

	f, err := parser.ParseFile(fset, path, nil, parser.ParseComments)
	if err != nil {
		return nil, err
	}

	a.state = a.getOrAddState("start")

	ast.Inspect(f, a.inspect)

	return a.stateList, nil
}

func (a *analyzer) getOrAddState(name string) *State {
	state, ok := a.stateMap[name]
	if !ok {
		state = &State{Name: name}
		a.stateMap[name] = state
		a.stateList = append(a.stateList, state)
	}
	return state
}

func (a *analyzer) inspect(n ast.Node) bool {
	switch x := n.(type) {

	case *ast.FuncDecl:
		if x.Name.Name == a.funcName {
			a.inFunc = true
		} else {
			a.inFunc = false
		}

	case *ast.BranchStmt:
		if a.inFunc && x.Tok == token.GOTO {
			nextState := a.getOrAddState(x.Label.Name)
			a.state.Next = append(a.state.Next, nextState)
		}

	case *ast.LabeledStmt:
		if a.inFunc {
			a.state = a.getOrAddState(x.Label.Name)
		}
	}

	return true
}
