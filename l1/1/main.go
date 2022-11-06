package main

import "fmt"

type Human struct {
	Name string
	Age  int
}

func (h *Human) Jump() {
	fmt.Printf("%s jumped\n", h.Name)
}

func (h *Human) Run() {
	fmt.Printf("%s is jumping\n", h.Name)
}

type Action struct {
	Human // Встраивание структуры Human в Action
}

func main() {
	a := Action{Human{"Jake", 15}}
	a.Jump()
	a.Run()
}
