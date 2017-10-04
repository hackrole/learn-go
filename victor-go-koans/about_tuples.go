package main

func sample() (string, bool) {
	return "hello", true
}

func TestTupleAssignment(t *T) {
	result, errcode := sample()

	t.AssertTrue("hello" == result)
	t.AssertTrue(true == errcode)
}

func TestTupleAssignmentBlankIdentifier(t *T) {
	result, _ := sample()

	t.AssertTrue("hello" == result)
}

func TestSwapWithTuples(t *T) {
	a := 1
	b := 2
	a, b = b, a
	t.AssertTrue(2 == a)
	t.AssertTrue(1 == b)

}
