package exportloopref_test

import (
	"testing"

	"github.com/kyoh86/exportloopref"
	"golang.org/x/tools/go/analysis/analysistest"
)

func Test(t *testing.T) {
	testdata := analysistest.TestData()
	analysistest.Run(t, testdata, exportloopref.Analyzer, "a")
}
