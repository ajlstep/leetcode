package isisomorphic

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPolindrome(t *testing.T) {
	var a string
	var b string
	var r bool

	a = "abcdefghijklmnopqrstuvwxyzva"
	b = "abcdefghijklmnopqrstuvwxyzck"
	r = false
	require.Equal(t, isIsomorphic(a, b), r, fmt.Sprintf("%s|%s => %v", a, b, r))
	a = "egg"
	b = "add"
	r = true
	require.Equal(t, isIsomorphic(a, b), r, fmt.Sprintf("%s|%s => %v", a, b, r))
	a = "paper"
	b = "title"
	r = true
	require.Equal(t, isIsomorphic(a, b), r, fmt.Sprintf("%s|%s => %v", a, b, r))
}
