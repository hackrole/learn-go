package main

import "fmt"

func TestCreatingArray(t *T) {
	var array [10]int
	t.AssertTrue(10 == len(array)) //The length of array is part of its type.
}

func TestArraysAreValues(t *T) {
	array1 := [...]int{1, 2, 3}
	var array2 [3]int
	array2 = array1
	t.AssertTrue("[1 2 3]" == fmt.Sprintf("%v", array2))
	array1[0] = 2
	t.AssertTrue("[1 2 3]" == fmt.Sprintf("%v", array2))
}

func TestAccessingArrayElements(t *T) {
	a := [...]string{"peanut", "butter", "and", "jelly"}

	t.AssertTrue("peanut" == a[0])
	t.AssertTrue("jelly" == a[3])
}

/*
  Slices can slice arrays only within arrays' bounds.
  Inverted slice ranges are not allowed.
*/
func TestSlicingArrays(t *T) {
	a := [...]string{"peanut", "butter", "and", "jelly"}

	t.AssertTrue("[peanut]" == fmt.Sprintf("%v", a[0:1]))
	t.AssertTrue("[peanut butter]" == fmt.Sprintf("%v", a[0:2]))
	t.AssertTrue("[]" == fmt.Sprintf("%v", a[2:2]))
	t.AssertTrue("[and jelly]" == fmt.Sprintf("%v", a[2:]))
	t.AssertTrue("[peanut butter]" == fmt.Sprintf("%v", a[:2]))
}
