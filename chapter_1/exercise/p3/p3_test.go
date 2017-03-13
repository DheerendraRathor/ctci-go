package p3

import "testing"

type testStr struct {
    input string
    output string
    length int
}

var testStrings []testStr = []testStr {
    {"Mr John Smith    ", "Mr%20John%20Smith", 13},
    { "", "", 0},
    {"   ", "   ", 0},
    {"ʥቕ ॵཚچ  ", "ʥቕ%20ॵཚچ", 6},
}

func TestURLifyWithSliceCopy(t *testing.T) {
    for index, testString := range testStrings {
        result := URLifyWithSliceCopy(testString.input, testString.length)
        if result != testString.output {
            t.Error("Input index:", index, "Expected Result:", testString.output, "Got:", result)
        }
    }
}

func TestURLifyTwoPass(t *testing.T) {
    for index, testString := range testStrings {
        result := URLifyTwoPass(testString.input, testString.length)
        if result != testString.output {
            t.Error("Input index:", index, "Expected Result:", testString.output, "Got:", result)
        }
    }
}