package excelsheet

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPolindrome(t *testing.T) {
	s := 35
	r := "A"
	require.Equal(t, convertToTitle(s), r, "asd")
}
