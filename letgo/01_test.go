package letgo

import (
	"reflect"
	"testing"
)

func TestTwoSum(t *testing.T) {

	t.Run("case1", func(t *testing.T) {
		d := TwoSum([]int{2, 7, 11, 15}, 9)
		if !reflect.DeepEqual(d, []int{0, 1}) {
			t.FailNow()
		}
	})

	t.Run("case2", func(t *testing.T) {
		d := TwoSum([]int{3, 2, 4}, 6)
		if !reflect.DeepEqual(d, []int{1, 2}) {
			t.FailNow()
		}
	})

}
