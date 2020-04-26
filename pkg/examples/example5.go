package examples

import "fmt"

func snd5(ch chan int) {

	var x int = 42
	ch <- x
}

func rcv5(ch chan int) {

	var x int
	x = <-ch
	fmt.Printf("received %d \n", x)
}

func ExampleFunc5() {

	var ch chan int = make(chan int)
	var ch2 chan int = make(chan int)
	go snd5(ch)
	go rcv5(ch2)
	ch2 <- <-ch
}
