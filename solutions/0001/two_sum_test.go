package two_sum

import (
	"testing"
	"reflect"
	"fmt"
)

var testCases = []struct {
	input  []int
	target int
	want   []int
}{
	{[]int{2, 7, 11, 15}, 9, []int{0, 1}},
	{[]int{3, 2, 4}, 6, []int{1, 2}},
	{[]int{3, 3}, 6, []int{0, 1}},
	{[]int{}, 999, []int{}},
}

func TestTwoSum(t *testing.T) {
	for i, tc := range testCases {
		testName := fmt.Sprintf("Test_%02d", i)
		t.Run(testName, func(t *testing.T) {
			got := twoSum(tc.input, tc.target)

			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("Test fail! | got: '%v', want: '%v'", tc.want, got)
			}
		})
	}
}
