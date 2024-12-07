// Разработать программу, которая будет последовательно отправлять значения в канал, 
//а с другой стороны канала — читать. 
// По истечению N секунд программа должна завершаться.

package main 

import (
	"time"
	"fmt"
	"sync"
)


func writer(wg* sync.WaitGroup, timeout time.Duration){
	go func() {
		for i := 1; ; i++ {
			time.Sleep(1 * time.Second)
			select {
			case <- time.After(timeout):
				wg.Done()
				return
			default:
				fmt.Println(i) 
			}
		}
	}()
}


func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	timeout := 10 * time.Second
	writer(&wg, timeout)
	wg.Wait()
}