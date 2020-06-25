package exportloopref_test

import (
	"testing"

	"github.com/kyoh86/exportloopref"
	"golang.org/x/tools/go/analysis/analysistest"
)

func TestSimple(t *testing.T) {
	testdata := analysistest.TestData()
	analysistest.Run(t, testdata, exportloopref.Analyzer, "simple")
}

func TestStruct(t *testing.T) {
	testdata := analysistest.TestData()
	analysistest.Run(t, testdata, exportloopref.Analyzer, "struct")
}

func TestComplex(t *testing.T) {
	testdata := analysistest.TestData()
	analysistest.Run(t, testdata, exportloopref.Analyzer, "complex")
}

func TestFixed(t *testing.T) {
	testdata := analysistest.TestData()
	analysistest.Run(t, testdata, exportloopref.Analyzer, "fixed")
}

func TestPslice(t *testing.T) {
	testdata := analysistest.TestData()
	analysistest.Run(t, testdata, exportloopref.Analyzer, "pslice")
}

func TestIssue2(t *testing.T) {
	testdata := analysistest.TestData()
	analysistest.Run(t, testdata, exportloopref.Analyzer, "issue2")
}
