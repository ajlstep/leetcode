package linkedlistcycle

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPolindrome(t *testing.T) {
	s := "PAYPALISHIRING"
	require.Equal(t, fmt.Sprintf("%p", &s), &s, "sts")
}
