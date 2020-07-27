package main

import "fmt"

type List []*int

func (l *List) AppendP(p *int) {
	*l = append(*l, p)
}

func main() {
	list := List{}

	println("loop expect exporting 10, 11, 12, 13")
	for _, p := range []int{10, 11, 12, 13} {
		fmt.Println(&p)  // not a diagnositc
		list.AppendP(&p) // want "exporting a pointer for the loop variable p"
	}

	println(`array expecting "10, 11, 12, 13" but "13, 13, 13, 13"`)
	for _, p := range ([]*int)(list) {
		printp(p)
	}
}

func printp(p *int) {
	println(*p)
}
