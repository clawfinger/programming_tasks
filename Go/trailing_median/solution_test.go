package trailingmedian

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPartition(t *testing.T) {
	type test struct {
		name             string
		input            []int32
		expectedPivotPos int
	}
	tests := []test{
		{
			name:             "sorted",
			input:            []int32{1, 2, 3, 4},
			expectedPivotPos: 3,
		},
		{
			name:             "reverse sorted",
			input:            []int32{4, 3, 2, 1},
			expectedPivotPos: 0,
		},
		{
			name:             "greatest pivot",
			input:            []int32{4, 5, 2, 1, 8, 3, 9},
			expectedPivotPos: 6,
		},
		{
			name:             "lowest pivot",
			input:            []int32{4, 5, 2, 9, 8, 3, 1},
			expectedPivotPos: 0,
		},
		{
			name:             "hackerrank 0",
			input:            []int32{2, 3, 4, 2, 3, 6, 8, 4, 5},
			expectedPivotPos: 6,
		},
		{
			name:             "two",
			input:            []int32{-1, -3},
			expectedPivotPos: 0,
		},
		{
			name:             "two 2",
			input:            []int32{0, 2},
			expectedPivotPos: 1,
		},
		{
			name:             "one",
			input:            []int32{0},
			expectedPivotPos: 0,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			res := partition(test.input, 0, len(test.input)-1)
			require.Equal(t, test.expectedPivotPos, res)
		})
	}

}

func TestMedianPos(t *testing.T) {
	type test struct {
		name              string
		input             []int32
		expectedMedianPos int
	}
	tests := []test{
		{
			name:              "sorted",
			input:             []int32{1, 2, 3, 4, 5},
			expectedMedianPos: 2,
		},
		{
			name:              "random",
			input:             []int32{1, 2, 3, 9, 4, 0, 6},
			expectedMedianPos: 3,
		},
		{
			name:              "hackerrank 0",
			input:             []int32{2, 3, 4, 3, 6},
			expectedMedianPos: 2,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			res := getMedianPos(test.input, 0, len(test.input)-1)
			require.Equal(t, test.expectedMedianPos, res)
		})
	}
}

func TestMdeian(t *testing.T) {
	type test struct {
		name           string
		input          []int32
		expectedMedian float32
	}
	tests := []test{
		{
			name:           "sorted odd",
			input:          []int32{1, 2, 3, 4, 5},
			expectedMedian: 3,
		},
		{
			name:           "sorted even",
			input:          []int32{1, 2, 3, 4},
			expectedMedian: 2.5,
		},
		{
			name:           "random",
			input:          []int32{1, 2, 3, 9, 4, 0, 6},
			expectedMedian: 3,
		},
		{
			name:           "hackerrank 0",
			input:          []int32{2, 3, 4, 2, 3, 6, 8, 4, 5},
			expectedMedian: 4,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			res := GetMedianWithSelect(test.input)
			require.Equal(t, test.expectedMedian, res)
		})
	}
}

func TestSlidingWindowMedian(t *testing.T) {
	type test struct {
		name       string
		input      []int32
		windowSize int32
		res        int32
	}
	tests := []test{
		{
			name:       "hackerrank 0",
			input:      []int32{2, 3, 4, 2, 3, 6, 8, 4, 5},
			windowSize: 5,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			res := ActivityNotifications(test.input, test.windowSize)
			require.Equal(t, test.res, res)
		})
	}
}

func TestSlidingWindowMedianWithArra(t *testing.T) {
	type test struct {
		name       string
		input      []int32
		windowSize int32
		res        int32
		medians    []float64
	}
	tests := []test{
		// {
		// 	name:       "leetcode",
		// 	input:      []int32{1, 3, -1, -3, 5, 3, 6, 7},
		// 	windowSize: 3,
		// 	medians:    []float64{1.00000, -1.00000, -1.00000, 3.00000, 5.00000, 6.00000},
		// },
		// {
		// 	name:       "leetcode 2",
		// 	input:      []int32{1, 2},
		// 	windowSize: 2,
		// 	medians:    []float64{1.50000},
		// },
		// {
		// 	name:       "leetcode 3",
		// 	input:      []int32{5, 2, 2, 7, 3, 7, 9, 0, 2, 3},
		// 	windowSize: 9,
		// 	medians:    []float64{1.50000},
		// },
		{
			name:       "leetcode 4",
			input:      []int32{7, 1, 8, 7},
			windowSize: 9,
			medians:    []float64{7.0},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			_, medians := ActivityNotificationsWithArray(test.input, test.windowSize)
			require.Equal(t, test.medians, medians)

		})
	}
}
