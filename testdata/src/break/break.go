package main

func main() {
	var result *int
	list := []int{1, 2}
	for _, p := range list {
		result = &p // Broken just below, the loop does NOT set unexpected value to result
		break
	}

	for _, p := range list {
		result = &p // want "exporting a pointer for the loop variable p"
		if *result == 1 {
			break // Static analyzer can't judge the loop breaks always.
		}
	}

	for _, p := range list {
		result = &p // want "exporting a pointer for the loop variable p"
		for {
			break // it does not break `for _, p := range list`
		}
	}
}
