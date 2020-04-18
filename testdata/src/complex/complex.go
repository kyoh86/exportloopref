package main

func main() {
	var ref struct {
		x struct {
			y []map[string]*int
			z []map[string]*int
		}
	}
	for i, p := range []int{10, 11, 12, 13} {
		if i%2 == 0 {
			ref = struct {
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
			ref.x.y = append(ref.x.y, map[string]*int{"baz": &p}) // want "exporting a pointer for the loop variable p"
			ref.x.z = append(
				ref.x.z,
				append(
					[]map[string]*int{{"baz": &p}}, // want "exporting a pointer for the loop variable p"
					map[string]*int{"baz": &p},     // want "exporting a pointer for the loop variable p"
					map[string]*int{},
				)...,
			)
		}
	}

	println(`captured value expecting "12" but "13"`)
	printp(ref.x.y[0]["baz"])
	printp(ref.x.y[1]["baz"])
	printp(ref.x.z[0]["baz"])
	printp(ref.x.z[1]["baz"])

}

func printp(p *int) {
	println(*p)
}
