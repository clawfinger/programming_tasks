package validstring

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestValidString(t *testing.T) {
	type test struct {
		name  string
		input string
		res   string
	}
	tests := []test{
		// {
		// 	name:  "hackerrank 0",
		// 	input: "abcdefghhgfedecba",
		// 	res:   "YES",
		// },
		{
			name:  "hackerrank 1",
			input: "aabbcd",
			res:   "NO",
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			res := IsValid(test.input)
			require.Equal(t, test.res, res)
		})
	}
}
