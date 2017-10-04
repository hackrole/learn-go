package main

func TestEmptyPointerIsNil(t *T) {
	var p *int
	t.AssertTrue(nil == p) //Put nil here instead of Intp__
}
