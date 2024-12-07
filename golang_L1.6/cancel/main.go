// Разработать программу, которая будет последовательно отправлять значения в канал, 
//а с другой стороны канала — читать. 
// По истечению N секунд программа должна завершаться.

package main 

import (
	"time"
	"fmt"
)


func writer(cancel chan struct{}) chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for i := 1; ; i++ {
			time.Sleep(1 * time.Second)
			select {
			case out <- i:
			case <- cancel:
				return
			}
		}
	}()
	return out
}

func reader(cancel chan struct{}, in chan int) {
	for {
		select {
		case m := <- in:
			fmt.Println(m)
		case <- cancel:
			return
		}
	}
}

func main() {
	cancel := make(chan struct{})
	defer close(cancel)
	w := writer(cancel)
	reader(cancel, w)
}