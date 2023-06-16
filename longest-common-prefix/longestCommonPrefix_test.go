package longestCommonPrefix

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestLongestCommonPrefix(t *testing.T) {
	require.Equal(t, longestCommonPrefix([]string{"cir", "car"}), "c", "test 4")
	require.Equal(t, longestCommonPrefix([]string{"ab", "a"}), "a", "test 3")
	require.Equal(t, longestCommonPrefix([]string{"flower", "flow", "flight"}), "fl", "test 1")
	require.Equal(t, longestCommonPrefix([]string{"dog", "racecar", "car"}), "", "test 2")
}
