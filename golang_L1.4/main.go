// # Реализовать постоянную запись данных в канал (главный поток).
// # Реализовать набор из N воркеров, которые читают произвольные данные из канала и выводят в stdout.
// #  Необходима возможность выбора количества воркеров при старте.

// # Программа должна завершаться по нажатию Ctrl+C. Выбрать и обосновать способ завершения работы всех воркеров.

package main

import (
	"fmt"
	"time"
	"os"
	"strconv"
)

func writer(cancel chan struct{}) chan int {
	input := make(chan int, 10)
	go func() {
		
		for i := 0; ; i++ {
			time.Sleep(1 * time.Second)
			select {
			case input <- i:
			case <-cancel:
				close(input)
				return
			}
		}
	}()
	return input
}

func reader(cancel chan struct{}, in chan int) {
	go func() {
		for {
			time.Sleep(3 * time.Second)
			select {
			case message := <-in:
				fmt.Println(message)
			case <-cancel:
				return
			}
		}
	}()

}

func main() {
	arguments := os.Args[1]
	workers, err := strconv.Atoi(arguments)
	if err != nil {
		fmt.Println("Worker must be digit")
		return
	}
	fmt.Println(workers)
	cancel := make(chan struct{})
	defer close(cancel)
	w := writer(cancel)
	for i := 0; i < workers; i++ {
		reader(cancel, w)
	}
	<- cancel
}
