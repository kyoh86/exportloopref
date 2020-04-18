package main

import (
	"github.com/kyoh86/exportloopref"
	"golang.org/x/tools/go/analysis/singlechecker"
)

func main() {
	singlechecker.Main(exportloopref.Analyzer)
}
