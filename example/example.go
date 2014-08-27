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
