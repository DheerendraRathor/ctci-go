package p4

import "testing"

type testStr struct {
    str string
    isPerm bool
}


var testStrings []testStr = []testStr {
    {"", true},
    {"a", true},
    {"abb", true},
    {"a abc", false},
    {"aab bcc", true},
    {"ʥቕ ʥቕ", true},
}

func TestIsPalindromePermutation(t *testing.T) {
    for index, testString := range testStrings {
        actualOutput := IsPalindromePermutation(testString.str)
        if actualOutput != testString.isPerm {
            t.Error("Index:", index, "String:", testString.str, "Expected Result:", testString.isPerm,
                "Actual Result:", actualOutput)
        }
    }
}