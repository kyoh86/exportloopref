package main

func main() {
	var result *int
	haystack := []int{1, 2}
	for _, hay := range haystack {
		result = &hay // Broken just below, the loop does NOT set unexpected value to result
		break
	}
}
