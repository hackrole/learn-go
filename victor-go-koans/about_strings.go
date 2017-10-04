package main

import "strings"
import "fmt"

func TestDoubleQuotedStringsAreStrings(t *T) {
	str := "hello" //A string can be defined with a literal

	t.AssertEquals("hello", str)
}

func TestBackQuotedStringsAreStrings(t *T) {
	str := `hello\n
world`

	t.AssertEquals(13, len(str))
}

func TestPlusConcatenatesString(t *T) {
	str := "Hello " + "World"
	t.AssertEquals("Hello World", str)
}

func TestPlusWillNotModifyOriginalStrings(t *T) {
	hi := "Hello, "
	there := "world"
	str := hi + there
	t.AssertEquals("Hello, ", hi)
	t.AssertEquals("world", there)
	t.AssertEquals("Hello, world", str)
}

func TestPlusEqualsWillAppendToEndOfString(t *T) {
	hi := "Hello, "
	there := "world"
	hi += there
	t.AssertEquals("Hello, world", hi)
}

func TestPlusEqualsAlsoLeavesOriginalStringUnmodified(t *T) {
	original := "Hello, "
	hi := original
	there := "world"
	hi += there
	t.AssertEquals("Hello, ", original)
}

func TestUseSprintfToInterpolateVaribales(t *T) {
	value1 := "one"
	value2 := 2
	str := fmt.Sprintf("The values are %s and %d", value1, value2)
	t.AssertEquals("The values are one and 2", str)
}

func TestYouCanGetASubstringFromAString(t *T) {
	str := "Bacon, lettuce and tomato"
	t.AssertEquals("let", str[7:10])
}

func TestYouCanGetASingleCharacterFromAString(t *T) {
	str := "Bacon, lettuce and tomato"
	t.AssertTrue('a' == str[1])
}

func TestCharactersAreBytesActually(t *T) {
	t.AssertTrue('b' == 'a'+1)
}

func TestStringsCanBeSplit(t *T) {
	str := "Sausage Egg Cheese"
	words := strings.Split(str, " ")
	t.AssertEqualsStringSlices([]string{"Sausage", "Egg", "Cheese"}, words)
}

func TestStringsCanBeSplitWithDifferentPatterns(t *T) {
	str := "the,|;rain;in,spain"
	words := strings.Split(str, ",|;")

	t.AssertEqualsStringSlices([]string{"the", "rain;in,spain"}, words)
}
