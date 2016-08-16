package example

import (
	"fmt"
	"reflect"
)

func snd6A(ch chan chan int) {
	y := make(chan int)
	go snd6B(y)
	ch <- y

}

func snd6B(ch chan int) {
	ch <- 42
}

func ExampleFunc6() {

	x := make(chan chan int)
	go snd6A(x)

	a := <-<-x

	fmt.Println("----Type: " + reflect.TypeOf(a).String())
}
