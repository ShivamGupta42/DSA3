package easy

import "testing"

func TestTwoSum(t *testing.T) {
	res := TwoSum([]int{2, 7, 11, 15}, 9)

	if res[0] != 0 || res[1] == 2 {
		t.Fail()
	}
}
