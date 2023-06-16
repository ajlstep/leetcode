package climbingstars

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPolindrome(t *testing.T) {
	s := 35
	r := 14930352
	require.Equal(t, climbStairs(s), r, fmt.Sprintf("%d|%d", s, r))
}
