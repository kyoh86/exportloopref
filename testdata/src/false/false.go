package main

type List []*int

func (l *List) AppendP(p *int) {
	*l = append(*l, p)
}

func main() {
	var slice []*int
	list := List{}

	println("loop expect exporting 10, 11, 12, 13")
	for _, v := range []int{10, 11, 12, 13} {
		list.AppendP(&v) // wanted "exporting a pointer for the loop variable v", but cannot be found

		p := &v                  // p is the local variable
		slice = append(slice, p) // wanted "exporting a pointer for the loop variable v", but cannot be found
	}

	println(`slice expecting "10, 11, 12, 13" but "13, 13, 13, 13"`)
	for _, p := range slice {
		printp(p)
	}
	println(`array expecting "10, 11, 12, 13" but "13, 13, 13, 13"`)
	for _, p := range ([]*int)(list) {
		printp(p)
	}
}

func printp(p *int) {
	println(*p)
}
