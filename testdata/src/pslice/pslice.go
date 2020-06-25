package main

func main() {
	var pslice []*int
	var ppslice []**int

	println("loop expecting 10, 11, 12, 13")
	for _, p := range []int{10, 11, 12, 13} {
		p := p
		pslice = append(pslice, &p) // not a diagnostic (fixed p)
	}

	for _, p := range pslice {
		ppslice = append(ppslice, &p) // want "exporting a pointer for the loop variable p"
	}

	println(`ppslice expecting "10, 11, 12, 13" but "13, 13, 13, 13"`)
	for _, p := range ppslice {
		printp(*p)
	}
}

func printp(p *int) {
	println(*p)
}
