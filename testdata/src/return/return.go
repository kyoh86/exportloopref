package main

func main() {
	result := sub()
	println(*result)
}

func sub() (result *int) {
	haystack := []int{1, 2}
	for _, hay := range haystack {
		result = &hay // Returns just below, the loop does NOT set unexpected value to result
		return
	}
	return
}
