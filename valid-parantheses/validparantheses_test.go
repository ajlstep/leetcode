package validpatantheses

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestIsValid(t *testing.T) {
	require.Equal(t, isValid("()"), true, "()")
	require.Equal(t, isValid("()[]{}"), true, "()[]{}")
	require.Equal(t, isValid("(]"), false, "(]")
}
