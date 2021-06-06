package palindrome_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	. "github.com/rizkyalviandra/go-pair/palindrome"
)

func TestIsPalindrome_True(t *testing.T) {
	assert.True(t, IsPalindrome("kodok"))
	assert.True(t, IsPalindrome("katak"))
	assert.True(t, IsPalindrome("aba"))
	assert.True(t, IsPalindrome("121"))
}

func TestIsPalindrome_False(t *testing.T) {
	assert.False(t, IsPalindrome("kita"))
	assert.False(t, IsPalindrome("ab"))
	assert.False(t, IsPalindrome("halo"))
	assert.False(t, IsPalindrome("hello"))
}