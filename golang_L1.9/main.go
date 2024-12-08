// Разработать конвейер чисел. Даны два канала: в первый пишутся числа (x) из массива,
// во второй — результат операции x*2,
// после чего данные из второго канала должны выводиться в stdout.

package main

import (
	"context"
	"fmt"
	"time"
)

func sender(ctx context.Context, data []int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for _, d := range data {
			time.Sleep(time.Second)
			select {
			case out <- d:
			case <-ctx.Done():
				return
			}
		}
	}()
	return out
}

func sqrter(ctx context.Context, in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for num := range in {
			select {
			case out <- num * num:
			case <-ctx.Done():
				return
			}
		}
	}()
	return out
}

func reader(in <-chan int) {
	for num := range in {
		fmt.Println(num)
	}
}

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	data := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	w := sender(ctx, data)
	s := sqrter(ctx, w)
	reader(s)

}
