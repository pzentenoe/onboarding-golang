package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func main() {

	sigs := make(chan os.Signal, 1)
	done := make(chan bool, 1)

	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		sig := <-sigs
		fmt.Println()
		fmt.Println(sig)
		done <- true
	}()

	fmt.Println("awaiting signal")
	<-done
	fmt.Println("exiting")

/*	channelNumbers := make(chan int, 10)
	go sendNumber(channelNumbers)

	for number := range channelNumbers {
		fmt.Println(number)
	}*/

}

func sendNumber(channel chan int) {
	for i := 0; i < 10; i++ {
		channel <- i + 1
	}
	close(channel)
}
