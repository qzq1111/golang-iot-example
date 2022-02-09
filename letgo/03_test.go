package letgo

import (
	"fmt"
	"testing"
)

func TestLengthOfLongestSubstring(t *testing.T) {

	t.Run("case1", func(t *testing.T) {

		fmt.Println(LengthOfLongestSubstring("abcabcbb"))
	})
}
