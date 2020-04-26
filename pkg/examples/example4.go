// SPDX-License-Identifier: MIT
package examples

import "fmt"
import "time"

func snd4(ch chan int) {

	var x int = 0
	for {
		x++
		ch <- x
		time.Sleep(1 * 1e9)
	}
}

func rcv4(ch chan int) {

	var i int
	for i = 0; i < 10; i = <-ch {
		fmt.Printf("received %d \n", i)
	}
}

func ExampleFunc4() {

	var ch chan int = make(chan int)
	go snd4(ch)
	rcv4(ch)
}
