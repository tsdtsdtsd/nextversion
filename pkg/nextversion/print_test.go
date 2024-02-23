package nextversion_test

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tsdtsdtsd/nextversion/pkg/nextversion"
)

type printTestcase struct {
	Format        string
	VersionResult *nextversion.Result
	Expected      string
}

func TestPrint(t *testing.T) {

	testTable := []printTestcase{
		// Format = simple
		{
			Format: "simple",
			VersionResult: &nextversion.Result{
				CurrentVersion:          "v0.1.1",
				CurrentVersionStrict:    "0.1.1",
				HasCurrentVersion:       true,
				NextVersion:             "v0.1.1",
				NextVersionStrict:       "0.1.1",
				HasNextVersion:          false,
				PrereleaseVersion:       "v0.1.1-rc+main.6f2f133",
				PrereleaseVersionStrict: "0.1.1-rc+main.6f2f133",
			},
			Expected: `CURRENT=v0.1.1
CURRENT_STRICT=0.1.1
HAS_CURRENT=true
NEXT=v0.1.1
NEXT_STRICT=0.1.1
HAS_NEXT=false
PRERELEASE=v0.1.1-rc+main.6f2f133
PRERELEASE_STRICT=0.1.1-rc+main.6f2f133
`,
		},
		{
			Format: "json",
			VersionResult: &nextversion.Result{
				CurrentVersion:          "v0.1.1",
				CurrentVersionStrict:    "0.1.1",
				HasCurrentVersion:       true,
				NextVersion:             "v0.1.1",
				NextVersionStrict:       "0.1.1",
				HasNextVersion:          false,
				PrereleaseVersion:       "v0.1.1-rc+main.6f2f133",
				PrereleaseVersionStrict: "0.1.1-rc+main.6f2f133",
			},
			Expected: `{"current":"v0.1.1","current-strict":"0.1.1","has-current":true,"next":"v0.1.1","next-strict":"0.1.1","has-next":false,"prerelease":"v0.1.1-rc+main.6f2f133","prerelease-strict":"0.1.1-rc+main.6f2f133"}`,
		},
	}

	var out bytes.Buffer

	for _, testCase := range testTable {
		out.Reset()
		nextversion.Print(&out, testCase.VersionResult, testCase.Format)
		assert.Equal(t, testCase.Expected, out.String())
	}
}
