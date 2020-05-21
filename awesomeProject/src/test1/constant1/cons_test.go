package constant1

import "testing"

const (
	Monday = 1 + iota
	Tuesday
	Wednesday
	Thursday
)

func TestConst(t *testing.T) {
	t.Log(Wednesday, Thursday)
	t.Log(Monday)

}
