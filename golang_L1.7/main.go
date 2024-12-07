package main

import (

	"fmt"
	"sync"
)

func main() {
	var dict sync.Map

	var wg sync.WaitGroup

	wg.Add(2)
	go generate(&wg, 100, &dict)
	go generate(&wg, 100, &dict)

	wg.Wait()
	fmt.Println(dict)
	

}

func generate(wg* sync.WaitGroup, mapLen int, dict* sync.Map) {
	for i := 0; i < mapLen; i++ {
		dict.Store(i, i)
	}
	wg.Done()
}


// func main() {
// 	dict := map[int]int{}
// 	var wg sync.WaitGroup
// 	var lock sync.Mutex
// 	wg.Add(2)

// 	go generate(&wg, &lock, 100, dict)
// 	go generate(&wg, &lock, 100, dict)

// 	wg.Wait()
// 	fmt.Print(dict)
// }



// func generate(wg* sync.WaitGroup, lock* sync.Mutex, mapLen int, dict map[int]int){
// 	for i := 0; i < mapLen; i++ {
// 		lock.Lock()
// 		dict[i]++
// 		lock.Unlock()
// 	}
// 	wg.Done()
// }

