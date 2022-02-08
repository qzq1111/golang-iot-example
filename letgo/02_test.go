package letgo

import (
	"fmt"
	"testing"
)

func TestAddTwoNumbers(t *testing.T) {

	t.Run("case1", func(t *testing.T) {
		d := AddTwoNumbers(
			&ListNode{2, &ListNode{4, &ListNode{3, nil}}},
			&ListNode{5, &ListNode{6, &ListNode{4, nil}}})
		fmt.Printf("%v", d)
	})

}
