package main

func main() {
	type singleLayer struct {
		num int
	}

	var array [4]*int
	var slice []*int
	var ref *int

	println("loop expecting 10, 11, 12, 13")
	for i, p := range []singleLayer{{10}, {11}, {12}, {13}} {
		printp(&p.num)                // not a diagnostic
		slice = append(slice, &p.num) // want "exporting a pointer for the loop variable p"
		array[i] = &p.num             // want "exporting a pointer for the loop variable p"
		if i%2 == 0 {
			ref = &p.num // want "exporting a pointer for the loop variable p"
		}
	}

	println(`slice expecting "10, 11, 12, 13" but "13, 13, 13, 13"`)
	for _, p := range slice {
		printp(p)
	}
	println(`array expecting "10, 11, 12, 13" but "13, 13, 13, 13"`)
	for _, p := range array {
		printp(p)
	}
	println(`captured value expecting "12" but "13"`)
	printp(ref)
}

func printp(p *int) {
	println(*p)
}
