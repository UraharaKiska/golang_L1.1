// Реализовать пересечение двух неупорядоченных множеств.

package main

import (
	"fmt"
)

type Set struct {
	elements map[int]struct{}
}

func NewSet() *Set {
	var set Set
	set.elements = make(map[int]struct{})
	return &set
}

func (s* Set) Add(value int) {
	s.elements[value] = struct{}{}
}

func (s* Set) Print() {
	out := "{"
	for key, _ := range s.elements {
		out += fmt.Sprintf("%d, ", key)
	}
	if len(out) > 1 {

	}
	if len(out) > 1 {
		out = out[:len(out)-2]
	}
	out += "}"
	fmt.Println(out)
}

func Intersection(s1, s2 Set) *Set {
	intersection := NewSet()
	for key, _ := range s1.elements {
		if _, exist := s2.elements[key]; exist {
			intersection.Add(key)
		}
	}
	return intersection
}

func main () {
	numbers1 := []int{0, 1, 232, 3, 4, 5, 6, -10}
	numbers2 := []int{7, 5, 4, 1, -10, 2012, 0}
	set1 := NewSet()
	set2 := NewSet()
	for _, n :=  range numbers1 {
		set1.Add(n)
	}
	for _, n :=  range numbers2 {
		set2.Add(n)
	}
	set1.Print()
	set2.Print()
	Intersection(*set1, *set2).Print()
}