package main

type MyStruct struct {
	MyStructPtrField *int
}

func main() {
	localVal := 0
	arr := []MyStruct{{&localVal}}
	for _, p := range arr {
		t := *p.MyStructPtrField
		globalMapInDifferentFile[t] = &p // want "exporting a pointer for the loop variable p"
	}
