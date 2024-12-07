package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	numbers := []int{2, 4, 6, 8, 10, 12, 14}
	inputChan := make(chan int)
	sqrtChan := make(chan int)
	done := make(chan int)
	go func() {
		defer close(inputChan)
		for _, n := range numbers {
			inputChan <- n
		}
	}()
	workers := 3
	for i := 0; i < workers; i++ {
		wg.Add(1)
		go CalcSqrt(&wg, inputChan, sqrtChan, done)
	}
	go func() {
		for i := 0; i < workers; i++ {
			<-done
		}
		close(sqrtChan)
	}()
	wg.Add(1)
	go Summer(&wg, sqrtChan)
	wg.Wait()

}

func CalcSqrt(wg *sync.WaitGroup, in chan int, out chan int, done chan int) {
	defer wg.Done()
	for n := range in {
		out <- n * n
	}
	done <- 0
}

func Summer(wg *sync.WaitGroup, in chan int) {
	defer wg.Done()
	sum := 0
	for sqrt := range in {
		sum += sqrt
	}
	fmt.Printf("sum = %d", sum)
}
