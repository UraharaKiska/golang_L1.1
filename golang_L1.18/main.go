// Реализовать структуру-счетчик, которая будет инкрементироваться в конкурентной среде.
// По завершению программа должна выводить итоговое значение счетчика.

package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

func Increment(wg* sync.WaitGroup, count *atomic.Int32) {
	time.Sleep(1 * time.Second)
	count.Add(1)
	wg.Done()
}

func main() {
	var count atomic.Int32
	var wg sync.WaitGroup
	for i := 0; i < 25; i++ {
		wg.Add(1)
		go Increment(&wg, &count)
	}
	wg.Wait()
	fmt.Println(count.Load())
}
