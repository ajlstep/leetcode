package plusone

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRemoveDuplicates(t *testing.T) {
	a := []int{9, 9}
	r := []int{1, 0, 0}
	require.Equal(t, plusOne(a), r, fmt.Sprintf("%v = %v", a, r))
	a = []int{4, 3, 2, 1}
	r = []int{4, 3, 2, 2}
	require.Equal(t, plusOne(a), r, fmt.Sprintf("%v = %v", a, r))
}
