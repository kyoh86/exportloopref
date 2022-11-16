package main

// Moving this map to main.go fixes nil pointer deference
var globalMapInDifferentFile = map[int]*MyStruct{}
