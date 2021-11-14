package util

import (
	"testing"
)

// Capitalises the first character of a string
func TestCapitaliseEmpty(t *testing.T) {
	str := Capitalise("")
	if len(str) != 0 {
		t.Fatalf(`Capitalise("") = %q, %v, want 0, error`, "", "Length not 0")
	}

}

func TestCapitaliseSmallCase(t *testing.T) {
	testStr := "abc"
	desiredStr := "Abc"
	str := Capitalise(testStr)
	if str != desiredStr {
		t.Fatalf(`Capitalise("%v") = %q, %v, want %v, error`, testStr, desiredStr,
			"string not capitalized", desiredStr)
	}

}
