package main

func main() {
	// int list

	var intArray [4]*int
	var intSlice []*int
	var intRef *int
	var intStr struct{ x *int }

	println("loop expecting 10, 11, 12, 13 and get also")
	for i, p := range []int{10, 11, 12, 13} {
		printp(&p)                      // not a diagnostic
		intSlice = append(intSlice, &p) // want "exporting a pointer for the loop variable p"
		intArray[i] = &p                // want "exporting a pointer for the loop variable p"
		if i%2 == 0 {
			intRef = &p   // want "exporting a pointer for the loop variable p"
			intStr.x = &p // want "exporting a pointer for the loop variable p"
		}
		var vStr struct{ x *int }
		var vArray [4]*int
		var v *int
		if i%2 == 0 {
			v = &p         // not a diagnostic (x is inner variable)
			vArray[1] = &p // not a diagnostic (x is inner variable)
			vStr.x = &p
		}
		_ = v
	}

	println(`slice expecting "10, 11, 12, 13" but "13, 13, 13, 13" will be gotten`)
	for _, p := range intSlice {
		printp(p)
	}
	println(`array expecting "10, 11, 12, 13" but "13, 13, 13, 13" will be gotten`)
	for _, p := range intArray {
		printp(p)
	}
	println(`captured value expecting "12" but "13" will be gotten`)
	printp(intRef)

	// single-layer members

	type singleLayer struct {
		num int
	}

	var singleArray [4]*int
	var singleSlice []*int
	var singleRef *int

	println("loop expecting 10, 11, 12, 13 and get also")
	for i, p := range []singleLayer{{10}, {11}, {12}, {13}} {
		printp(&p.num)                            // not a diagnostic
		singleSlice = append(singleSlice, &p.num) // want "exporting a pointer for the loop variable p"
		singleArray[i] = &p.num                   // want "exporting a pointer for the loop variable p"
		if i%2 == 0 {
			singleRef = &p.num // want "exporting a pointer for the loop variable p"
		}
	}

	println(`slice expecting "10, 11, 12, 13" but "13, 13, 13, 13" will be gotten`)
	for _, p := range singleSlice {
		printp(p)
	}
	println(`array expecting "10, 11, 12, 13" but "13, 13, 13, 13" will be gotten`)
	for _, p := range singleArray {
		printp(p)
	}
	println(`captured value expecting "12" but "13" will be gotten`)
	printp(singleRef)

	// complex
	var cmpRef struct {
		x struct {
			y []map[string]*int
			z []map[string]*int
		}
	}
	for i, p := range []int{10, 11, 12, 13} {
		if i%2 == 0 {
			cmpRef = struct {
				x struct {
					y []map[string]*int
					z []map[string]*int
				}
			}{
				x: struct {
					y []map[string]*int
					z []map[string]*int
				}{
					y: []map[string]*int{{
						"baz": &p, // want "exporting a pointer for the loop variable p"
					}},
				},
			}
			cmpRef.x.y = append(cmpRef.x.y, map[string]*int{"baz": &p}) // want "exporting a pointer for the loop variable p"
			cmpRef.x.z = append(
				cmpRef.x.z,
				append(
					[]map[string]*int{{"baz": &p}}, // want "exporting a pointer for the loop variable p"
					map[string]*int{"baz": &p},     // want "exporting a pointer for the loop variable p"
					map[string]*int{},
				)...,
			)
		}
	}

	println(`captured value expecting "12" but "13" will be gotten`)
	printp(cmpRef.x.y[0]["baz"])
	printp(cmpRef.x.y[1]["baz"])
	printp(cmpRef.x.z[0]["baz"])
	printp(cmpRef.x.z[1]["baz"])

}

func printp(p *int) {
	println(*p)
}
