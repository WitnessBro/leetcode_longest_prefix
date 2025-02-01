package leetcodelongestprefix

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	var input = []string{"flower", "flow", "flight"}
	var expectedOutput = "fl"
	var k = longestCommonPrefix(input)

	assert.Equal(t, expectedOutput, k)
}
func longestCommonPrefix(strs []string) string {
	if len(strs) == 1 {
		return strs[0]
	}
	if len(strs) == 0 {
		return ""
	}
	first := strs[0]
	if len(first) == 0 {
		return ""
	}
	index := 0
	for i := 0; i < len(first); i++ {
		for _, val := range strs {
			if len(val) == 0 {
				return ""
			}
			if i > len(val)-1 || val[i] != first[i] {
				return first[:i]
			}
			index = i
		}
	}
	return first[:index+1]
}
