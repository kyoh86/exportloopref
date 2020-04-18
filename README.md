# exportloopref

An analyzer that finds exporting pointers for loop variables.

[![Go Report Card](https://goreportcard.com/badge/github.com/kyoh86/exportloopref)](https://goreportcard.com/report/github.com/kyoh86/exportloopref)
[![Coverage Status](https://img.shields.io/codecov/c/github/kyoh86/exportloopref.svg)](https://codecov.io/gh/kyoh86/exportloopref)

## What's this?

Sample problem code from: https://github.com/kyoh86/scopelint/blob/master/example/readme.go

```
6  values := []string{"a", "b", "c"}
7  var funcs []func()
8  for _, val := range values {
9  	funcs = append(funcs, func() {
10 		fmt.Println(val)
11 	})
12 }
13 for _, f := range funcs {
14 	f()
15 }
16 /*output:
17 c
18 c
19 c
20 */
21 var copies []*string
22 for _, val := range values {
23 	copies = append(copies, &val)
24 }
25 /*(in copies)
26 &"c"
27 &"c"
28 &"c"
29 */
```

In Go, the `val` variable in the above loops is actually a single variable.
So in many case (like the above), using it makes for us annoying bugs.

You can find them with `scopelint`, and fix it.

```
$ scopelint ./example/readme.go
example/readme.go:10:16: Using the variable on range scope "val" in function literal
example/readme.go:23:28: Using a reference for the variable on range scope "val"
Found 2 lint problems; failing.
```

(Fixed sample):

```go
values := []string{"a", "b", "c"}
var funcs []func()
for _, val := range values {
  val := val // pin!
	funcs = append(funcs, func() {
		fmt.Println(val)
	})
}
for _, f := range funcs {
	f()
}
var copies []*string
for _, val := range values {
  val := val // pin!
	copies = append(copies, &val)
}
```

## Install

```
go get github.com/kyoh86/exportloopref
```

## Usage

```
exportloopref --help
```

# LICENSE

[![MIT License](http://img.shields.io/badge/license-MIT-blue.svg)](http://www.opensource.org/licenses/MIT)

This is distributed under the [MIT License](http://www.opensource.org/licenses/MIT).
