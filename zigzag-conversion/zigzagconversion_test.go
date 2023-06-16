package zigzagconversion

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPolindrome(t *testing.T) {
	s := "PAYPALISHIRING"
	numRows := 4
	r := "PINALSIGYAHRPI"
	require.Equal(t, convert(s, numRows), r, fmt.Sprintf("rows:%d|%s => %s", numRows, s, r))
}
