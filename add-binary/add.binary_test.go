package addbinary

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestAddBinary(t *testing.T) {
	a := "11"
	b := "1"
	require.Equal(t, addBinary(a, b), "100", fmt.Sprintf("%s + %s = 100", a, b))
	a = "1110001"
	b = "110100100"
	require.Equal(t, addBinary(a, b), "1000010101", fmt.Sprintf("%s + %s = 1000010101", a, b))
}
