// SPDX-License-Identifier: MIT
package example

import (
	"fmt"
	"reflect"
)

// If := (Init|Eps)⋅Cond⋅(Body|Else|Eps)
// Else := Body|If

// Init := (SndSmt⋅(RcvSmt*))
// Cond := RcvSmt*
// Body := (SndSmt|RcvSmt)*

func ExampleFuncIf() {

	// IfStmt: Init == nil, Cond, Body (BlockStmt), Else == nil
	if true {
		fmt.Println("Expected case: Body")
	}

	// IfStmt: Init, Cond, Body (BlockStmt), Else == nil
	if a := true; a {
		fmt.Println("Expected case: Body")
	}

	// IfStmt: Init == nil, Cond, Body (BlockStmt), Else (BlockStmt)
	if true {
		fmt.Println("Expected: Body")
	} else {
		fmt.Println("Unexpected: Else")
	}

	// IfStmt: Init, Cond, Body (BlockStmt), Else (BlockStmt)
	if a := true; a {
		fmt.Println("Expected: Body")
	} else {
		fmt.Println("Unexpected: Else")
	}

	// ch! .(Body|ch2!).(Body2|E)

	// IfStmt: Init == nil, Cond, Body (BlockStmt), Else (IfStmt)
	if true {
		fmt.Println("Expected: Body")
	} else if false {
		fmt.Println("Unexpected: else if Body")
	}

	// IfStmt: Init, Cond, Body (BlockStmt), Else (IfStmt)
	if a := true; a {
		fmt.Println("Expected: Body")
	} else if !a {
		fmt.Println("Unexpected: else if Body")
	}

	// #################################

	a := make(chan chan chan chan int)
	b := make(chan chan chan int)

	go func(ch chan chan chan int) {
		ch <- make(chan chan int)
	}(b)

	go func(ch chan chan chan chan int, ch2 chan chan chan int) {
		ch <- ch2
	}(a, b)

	c := <-<-a

	fmt.Println("----Type a: " + reflect.TypeOf(a).String())
	fmt.Println("----Type b: " + reflect.TypeOf(b).String())
	fmt.Println("----Type c: " + reflect.TypeOf(c).String())
}

func ExampleFuncIf_Snd_Rcv() {

	var ch chan bool = make(chan bool)

	// IfStmt: Init == nil, Cond, Body (BlockStmt), Else == nil; RcvSmt in Cond
	go snd(ch)
	if <-ch {
		fmt.Println("Expected case: Body")
	}

	// IfStmt: Init, Cond, Body (BlockStmt), Else == nil; SndSmt in Init
	go rcv(ch)
	if ch <- true; true {
		fmt.Println("Expected case: Body")
	}

	// IfStmt: Init, Cond, Body (BlockStmt), Else == nil; RcvSmt in Init
	go snd(ch)
	if a := <-ch; a {
		fmt.Println("Expected case: Body")
	}

	// IfStmt: Init == nil, Cond, Body (BlockStmt), Else (BlockStmt); SndSmt in Body
	go rcv(ch)
	if true {
		ch <- true
		fmt.Println("Expected: Body")
	} else {
		fmt.Println("Unexpected: Else")
	}

	// IfStmt: Init == nil, Cond, Body (BlockStmt), Else (BlockStmt); RcvSmt in Body
	go snd(ch)
	if true {
		<-ch
		fmt.Println("Expected: Body")
	} else {
		fmt.Println("Unexpected: Else")
	}

	// IfStmt: Init == nil, Cond, Body (BlockStmt), Else (BlockStmt); SndSmt in Else
	go rcv(ch)
	if false {
		fmt.Println("Unexpected: Body")
	} else {
		ch <- true
		fmt.Println("Expected: Else")
	}

	// IfStmt: Init == nil, Cond, Body (BlockStmt), Else (BlockStmt); RcvSmt in Else
	go snd(ch)
	if false {
		<-ch
		fmt.Println("Unexpected: Body")
	} else {
		<-ch
		fmt.Println("Expected: Else")
	}
}

func snd(ch chan bool) {
	ch <- true
}

func rcv(ch chan bool) {
	<-ch
}
