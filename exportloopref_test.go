package exportloopref_test

import (
	"testing"

	"github.com/kyoh86/exportloopref"
	"golang.org/x/tools/go/analysis/analysistest"
)

func Test(t *testing.T) {
	testdata := analysistest.TestData()
	analysistest.Run(t, testdata, exportloopref.Analyzer, "simple")
	analysistest.Run(t, testdata, exportloopref.Analyzer, "struct")
	analysistest.Run(t, testdata, exportloopref.Analyzer, "complex")
	analysistest.Run(t, testdata, exportloopref.Analyzer, "fixed")
}
