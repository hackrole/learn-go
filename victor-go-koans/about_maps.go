package main

func TestCreatingMap(t *T) {
	var timeZone = map[string]int{
		"UTC": 0 * 60 * 60,
		"EST": -5 * 60 * 60,
		"CST": -6 * 60 * 60,
		"MST": -7 * 60 * 60,
		"PST": -8 * 60 * 60,
	}
	t.AssertTrue(5 == len(timeZone))
	t.AssertTrue(0 == timeZone["UTC"])
}
