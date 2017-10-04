package main

import "fmt"

func TestClosuresCanBeAssignedToVariablesAndCalledExplicitly(t *T) {
	addOne := func(n int) int { return n + 1 }
	t.AssertEquals(11, addOne(10))
}

func MakeOrder(order string) func(qty int) string {
	return func(qty int) string { return fmt.Sprintf("%d", qty) + " " + order + "s" }
}

func TestAccessingClosureViaAssignment(t *T) {
	sausages := MakeOrder("sausage")
	eggs := MakeOrder("egg")

	t.AssertEquals("3 sausages", sausages(3))
	t.AssertEquals("2 eggs", eggs(2))
}

func TestAccessingClosureWithoutAssignment(t *T) {
	t.AssertEquals("39823 spams", MakeOrder("spam")(39823))
}
