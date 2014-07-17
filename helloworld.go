// garbage program for garbage

package main

import "fmt"

type HelloWorld struct {
    Msg string
}

func MakeHelloWorld(msg string) HelloWorld {
    var h HelloWorld
    h.Msg = msg
    return h
}

func PrintHelloWorld(h HelloWorld) {
    fmt.Println(h.Msg)
}

func main() {
    h := MakeHelloWorld("Buttes and Donges")
    PrintHelloWorld(h)
}