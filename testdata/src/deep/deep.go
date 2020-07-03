package main

type Container struct {
	Inner
}

type Inner struct {
	Number int
}

func main() {
	var array [4]*int
	var slice []*int
	var ref *int
	var str struct{ x *int }

	var target = []Container{{Inner{10}}, {Inner{11}}, {Inner{12}}, {Inner{13}}}

	// access to unsafe member
	println("loop expecting 10, 11, 12, 13")
	for i, p := range target {
		printp(&p.Inner.Number)
		slice = append(slice, &p.Inner.Number) // want "exporting a pointer for the loop variable p"
		array[i] = &p.Inner.Number             // want "exporting a pointer for the loop variable p"
		if i%2 == 0 {
			ref = &p.Inner.Number   // want "exporting a pointer for the loop variable p"
			str.x = &p.Inner.Number // want "exporting a pointer for the loop variable p"
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
