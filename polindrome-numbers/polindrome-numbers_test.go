package polindromenumbers

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPolindrome(t *testing.T) {
	require.Equal(t, isPalindrome(-252), false, "int = -252|true")
	require.Equal(t, isPalindrome(121), true, "int = 121|true")
	require.Equal(t, isPalindrome(1231), false, "int = 1231|false")
	require.Equal(t, isPalindrome(13531), true, "int = 13531|true")
}
