// Имеется последовательность строк - (cat, cat, dog, cat, tree) создать для нее собственное множество.

package main

import (
	"fmt"
)

type Set struct {
	elements map[string]struct{}
}

func NewSet() *Set {
	var set Set
	set.elements = make(map[string]struct{})
	return &set
}

func (s* Set) Add(value string) {
	s.elements[value] = struct{}{}
}

func (s* Set) Print() {
	out := "{"
	for key, _ := range s.elements {
		out += fmt.Sprintf("%s, ", key)
	}
	if len(out) > 1 {

	}
	if len(out) > 1 {
		out = out[:len(out)-2]
	}
	out += "}"
	fmt.Println(out)
}

func main() {
	strs := []string{"cat", "cat", "dog", "cat", "tree"}
	set := NewSet()
	for _, s := range strs {
		set.Add(s)
	}
	subset := NewSet()
	subset.Add("cat")
	subset.Add("dog")
	set.Print()
	fmt.Print("Subset: ")
	subset.Print()

}