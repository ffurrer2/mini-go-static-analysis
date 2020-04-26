// SPDX-License-Identifier: MIT
package example

import (
	"fmt"
)

// Switch := (Init|Eps)⋅(Tag|Eps)⋅Body

// Body := CaseClause*
// Init := (SndSmt⋅(RcvSmt*))
// Cond := RcvSmt*

// CaseClause := (Expr*)⋅Body

// Body := (SndSmt|RcvSmt)*

func ExampleFuncSwitch() {

	// SwitchStmt: Init == nil, Tag == nil, Body (BlockStmt); no CaseClauses
	switch {
	}

	// SwitchStmt: Init == nil, Tag, Body (BlockStmt); no CaseClauses
	switch true {
	}

	// SwitchStmt: Init, Tag, Body (BlockStmt); no CaseClauses
	switch a := true; a {
	}

	// SwitchStmt: Init == nil, Tag == nil, Body (BlockStmt); only default CaseClause
	switch {
	default:
		fmt.Println("Expected case: default")
	}

	// SwitchStmt: Init, Tag == nil, Body (BlockStmt); only default CaseClause
	switch a := true; {

	default:
		fmt.Printf("Expected case: default, a := %v\n", a)
	}

	// SwitchStmt: Init, Tag == nil, Body (BlockStmt); only default CaseClause
	switch true {

	default:
		fmt.Printf("Expected case: default")
	}

	// SwitchStmt: Init == nil, Tag == nil, Body (BlockStmt)
	switch {

	case true:
		fmt.Println("Expected case: true")
		fmt.Println("Expected case: true")
		fallthrough

	case false:
		fmt.Println("Unexpected case")

	default:
		fmt.Println("Unexpected case")
	}

	// SwitchStmt: Init, Tag == nil, Body (BlockStmt)
	switch a := true; {

	case a:
		fmt.Println("Expected case: true")

	case !a:
		fmt.Println("Unexpected case")

	default:
		fmt.Println("Unexpected case")
	}

	// SwitchStmt: Init == nil, Tag, Body (BlockStmt)
	switch true {

	case true:
		fmt.Println("Expected case: true")

	case false:
		fmt.Println("Unexpected case")

	default:
		fmt.Println("Unexpected case")
	}

	// SwitchStmt: Init, Tag Body (BlockStmt)
	switch a := true; a {

	case a:
		fmt.Println("Expected case: true")

	case !a:
		fmt.Println("Unexpected case")

	default:
		fmt.Println("Unexpected case")
	}

	// SwitchStmt: Init == nil, Tag == nil, Body (BlockStmt); list of expressions
	switch {

	case true, true == true, 1 == 1:
		fmt.Println("Expected case: true")

	case false:
		fmt.Println("Unexpected case")

	default:
		fmt.Println("Unexpected case")
	}
}

func ExampleFuncSwitch_Snd_Rcv() {

	var ch chan bool = make(chan bool)

	// SwitchStmt: Init == nil, Tag == nil, Body (BlockStmt); no CaseClauses; no SndSmt/RcvSmt
	switch {
	}

	// SwitchStmt: Init, Tag == nil, Body (BlockStmt); no CaseClauses; SndSmt in Init
	go rcv(ch)
	switch ch <- true; {
	}

	// SwitchStmt: Init, Tag == nil, Body (BlockStmt); no CaseClauses; RcvSmt in Init
	go snd(ch)
	switch <-ch; {
	}

	// SwitchStmt: Init == nil, Tag, Body (BlockStmt); no CaseClauses; RcvSmt in Tag
	go snd(ch)
	switch <-ch {
	}

	// SwitchStmt: Init == nil, Tag == nil, Body (BlockStmt); only default CaseClause; SndSmt in Body
	go rcv(ch)
	switch {
	default:
		ch <- true
		fmt.Println("Expected case: default")
	}

	// SwitchStmt: Init == nil, Tag == nil, Body (BlockStmt); only default CaseClause; RcvSmt in Body
	go snd(ch)
	switch {
	default:
		<-ch
		fmt.Println("Expected case: default")
	}

	// SwitchStmt: Init == nil, Tag == nil, Body (BlockStmt); RcvSmt in Body (expression)
	go snd(ch)
	switch {

	case <-ch:
		fmt.Println("Expected case: true")

	case false:
		fmt.Println("Unexpected case")

	default:
		fmt.Println("Unexpected case")
	}

	// SwitchStmt: Init == nil, Tag == nil, Body (BlockStmt); multiple RcvSmt in Body (expression list)
	go snd(ch)
	go snd(ch)
	switch {

	case <-ch, <-ch:
		fmt.Println("Expected case: true")

	case false:
		fmt.Println("Unexpected case")

	default:
		fmt.Println("Unexpected case")
	}
	return

	// SwitchStmt: Init, Tag, Body
	go snd(ch)
	go snd(ch)
	switch x := <-ch; x && <-x {

	case <-ch, <-ch:
		fmt.Println("Expected case: true")

	case false:
		fmt.Println("Unexpected case")

	default:
		fmt.Println("Unexpected case")
	}
	return
}

func snd(ch chan bool) {
	ch <- true
}

func rcv(ch chan bool) {
	<-ch
}
