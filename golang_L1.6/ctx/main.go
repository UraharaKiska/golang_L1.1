// Разработать программу, которая будет последовательно отправлять значения в канал,
//а с другой стороны канала — читать.
// По истечению N секунд программа должна завершаться.

package main

import (
	"context"
	"fmt"
	"time"
)


func writer(ctx context.Context) chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for i := 1; ; i++ {
			time.Sleep(1 * time.Second)
			select {
			case out <- i:
			case <- ctx.Done():
				return
			}
		}
	}()
	return out
}

func reader(ctx context.Context, in chan int) {
	for {
		select {
		case m := <- in:
			fmt.Println(m)
		case <- ctx.Done():
			return
		}
	}
}

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	w := writer(ctx)
	timeout := 10 * time.Second
	go func() {
		select {
		case <- time.After(timeout):
			cancel()	
		}
	}()
	reader(ctx, w)
}