package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	numbers := []int{2, 4, 6, 8, 10, 12, 14}
	input := make(chan int)
	go func() {
		for _, n := range numbers {
			input <- n
		}
		close(input)
	}()
	wg.Add(1)
	go CalcSqrt(&wg, input)
	wg.Add(1)
	go CalcSqrt(&wg, input)
	wg.Wait()

}

func CalcSqrt(wg *sync.WaitGroup, out chan int) {
	defer wg.Done()
	for n := range out {
		fmt.Printf("sqrt(%d) = %d\n", n, n * n)
	}
	
}
