package lucksaving

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestLuckSaving(t *testing.T) {
	type test struct {
		name  string
		input [][]int32
		luck  int32
		res   int32
	}
	tests := []test{
		{
			name:  "hackerrank 0",
			input: [][]int32{{5, 1}, {2, 1}, {1, 1}, {8, 1}, {10, 0}, {5, 0}},
			luck:  3,
			res:   29,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			res := LuckBalance(test.luck, test.input)
			require.Equal(t, test.res, res)
		})
	}
}
