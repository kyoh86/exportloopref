package main

type Container struct {
	Pointer *Inner
}

type Inner struct {
	Number int
}

func main() {
	var array [4]*int
	var slice []*int
	var ref *int
	var str struct{ x *int }

	target := []Container{{&Inner{10}}, {&Inner{11}}, {&Inner{12}}, {&Inner{13}}}

	// access to unsafe member
	println("loop expecting 10, 11, 12, 13")
	for i, p := range target {
		printp(&p.Pointer.Number)
		slice = append(slice, &p.Pointer.Number)
		array[i] = &p.Pointer.Number
		if i%2 == 0 {
			ref = &p.Pointer.Number
			str.x = &p.Pointer.Number
		}
	}

	println(`slice expecting "10, 11, 12, 13"`)
	for _, p := range slice {
		printp(p)
	}
	println(`array expecting "10, 11, 12, 13"`)
	for _, p := range array {
		printp(p)
	}
	println(`captured value expecting "12"`)
	printp(ref)
}

func printp(p *int) {
	println(*p)
}
