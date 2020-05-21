package fib

// 菲薄纳妾

import "testing"

func TestFib(t *testing.T) {
	var a int = 1
	var b int = 1
	t.Log(a)
	t.Log(b)
	for i := 0; i < 5; i++ {
		c := a + b
		t.Log(c)
		a, b = b, c

	}
}
