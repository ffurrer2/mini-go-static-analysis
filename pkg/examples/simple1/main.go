// SPDX-License-Identifier: MIT
package main

import (
	"fmt"
)

func main() {

	var ch chan int = make(chan int)

	for i := <-ch; i < <-ch; i++ {
		fmt.Printf("Hello World!\n")
		<-ch
	}

	chanArray := []chan int{make(chan int), make(chan int), make(chan int)}

	<-chanArray[<-ch]

	go func() {
		ch <- 5
	}()

	if 1 < <-ch {
		ch <- 5
	}

	go snd11(ch)
	rcv11(ch)

	<-getChan()
	getChan() <- 5

	close(ch)
}

func snd11(ch chan int) {
	var x int = 42
	ch <- x
}

func rcv11(ch chan int) {
	y := <-ch
	ch <- y
}

func getChan() chan int {
	return make(chan int)
}
