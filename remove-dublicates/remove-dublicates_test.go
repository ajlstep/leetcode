package removedublicates

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRemoveDuplicates(t *testing.T) {
	a := []int{1, 1, 2}
	r := 2
	require.Equal(t, removeDuplicates(a), r, fmt.Sprintf("%v = %d", a, r))
	a = []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	r = 5
	require.Equal(t, removeDuplicates(a), r, fmt.Sprintf("%v = %d", a, r))
}
