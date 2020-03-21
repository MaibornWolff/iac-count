package ansible

import (
	"testing"

	input "github.com/MaibornWolff/iac-count/pkg/input"
)

func TestRloc(t *testing.T) {
	tests := map[string]metricTest{
		"successful rloc Calculation": {
			path:       "test/data/main.yml",
			content:    input.ReadFileToString("test/data/main.yml"),
			calculator: RlocCalculator{},
			output:     26,
		},
	}

	runMetricTest(t, tests)
}
