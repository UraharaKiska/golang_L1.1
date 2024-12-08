// Разработать программу, которая в рантайме способна определить тип 
// переменной: int, string, bool, channel из переменной типа interface{}.

package main

import (
	"fmt"
)

func main() {
	var variable any = make(chan struct{})
	switch t := variable.(type){
	case int:
		fmt.Println("is int:", t)
	case bool:
		fmt.Println("is bool:", t)
	case string:
		fmt.Println("is string:", t)
	case chan int, chan bool, chan string, chan struct{}:
		fmt.Println("is channel:", t)
	default:
		fmt.Println("type unknown")
	}
}