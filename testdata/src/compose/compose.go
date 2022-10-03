package main

func main() {
	type singleLayer struct {
		num *int
	}
	var slice []singleLayer

	for _, v := range []int{10, 11, 12, 13} {
		r := singleLayer{num: &v} // want "exporting a pointer for the loop variable v"
		slice = append(slice, r)
	}
	println(`slice expecting "10, 11, 12, 13" but "13, 13, 13, 13"`)
	for _, p := range slice {
		printp(p.num)
	}
}

func printp(p *int) {
	println(*p)
}
