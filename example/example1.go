package example

import "fmt"
import _ "runtime"
import _ "sync"

func snd1(ch chan int) {

	var x int = 42
	fmt.Printf("send %d \n", x)
	ch <- x
}

func rcv1(ch chan int) {

	y := <-ch
	fmt.Printf("received %d \n", y)
}

func ExampleFunc1() {

	var ch chan int = make(chan int)
	go snd1(ch)
	rcv1(ch)
}
