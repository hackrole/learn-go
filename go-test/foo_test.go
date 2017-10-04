package foo

import "testing"

func TestSomething(t *testing.T) {
	t.Fail()
}

func TestError(t *testing.T) {
	t.Error("bad mood.")
}
