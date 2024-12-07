package main

import "fmt"

type Human struct {
	name   string
	age    int
	sex    string
	height int
	gopher bool
}

type Action struct {
	Human
}

func NewHuman(name string, age int, sex string, height int, gopher bool) Human {
	human := Human{name: name, age: age, sex: sex, height: height, gopher: gopher}
	return human
}

func (h *Human) GetName() string {
	return h.name

}

func (h *Human) GetAge() int {
	return h.age
}

func (h *Human) GetSex() string {
	return h.sex
}

func (h *Human) GetHeight() int {
	return h.height
}

func (h *Human) IsGopher() bool {
	return h.gopher
}

func main() {
	human := NewHuman("Kris", 20, "m", 185, true)
	action := Action{human}
	fmt.Println(action.GetName())
	fmt.Println(action.GetSex())
	fmt.Println(action.GetAge())
	fmt.Println(action.GetHeight())
	fmt.Println(action.IsGopher())

}
