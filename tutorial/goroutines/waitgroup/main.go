package main

import (
	"fmt"
	"sync"
)

func main() {

	wg := sync.WaitGroup{}
	wg.Add(1)

	fmt.Println("Hola")
	go func() {
		fmt.Println("Hola desde goroutine")
		wg.Done()
	}()
	wg.Add(1)
	go printNumbers(&wg)
	fmt.Println("Finalizando")

	wg.Wait()

}

func printNumbers(wg *sync.WaitGroup) {
	mu := sync.Mutex{}
	mu.Lock()
	defer mu.Unlock()
	for i := 0; i < 100; i++ {
		fmt.Println(i + 1)
	}

	wg.Done()
}
