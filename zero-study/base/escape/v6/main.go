package main

import (
	"fmt"
)

type MyInterface interface {
	DoSomething()
}

type MyStruct struct {
	Name string
}

func (m *MyStruct) DoSomething() {
	fmt.Println("Doing something:", m.Name)
}

func ProcessInterface(i MyInterface) {
	i.DoSomething()
}

func main() {
	m := &MyStruct{Name: "Test"}
	ProcessInterface(m)
}
