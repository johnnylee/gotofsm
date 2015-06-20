gotofsm
=======

Tools for working with goto-based state machines in go (golang).

usage
-----

The program in bin/gotofsm-dot.go will print a graphviz file for a
state machine implemented using goto statements. Consider the
following file:

```go
package main

func Run() {

start: 
	if true {
		goto state1
	}
	goto state2

state1: 
	goto state3

state2: 
	if true {
		goto state4
	}
	goto state1

state3: 
	if true {
		goto start
	}
	goto end

state4: 
	if true {
		goto state1
	}
	goto start

end:
}

func main() {}

```

Assuming that this file is named example.go, we run:

```
gotofsm-dot example.go Run
```

Which prints

```
digraph {
start->state1;
start->state2;
state1->state3;
state2->state4;
state2->state1;
state3->start;
state3->end;
state4->state1;
state4->start;
}
```

The bash script bin/gotofsm-dot-viewer takes the same arguments as
gotofsm-dot, but pipes the output to a file and launches the xdot
viewer.

Here's what the output looks like as a png: 

![dot output](https://raw.githubusercontent.com/johnnylee/gotofsm/master/example/output.png)
