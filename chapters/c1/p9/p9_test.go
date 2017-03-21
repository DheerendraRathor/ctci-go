package p9

import "testing"

type testStruct struct {
    input string
    test string
    result bool
}

var testData []testStruct = []testStruct {
    {"", "", true},
    {"a", "a", true},
    {"a", "b", false},
    {"ab", "ba", true},
    {"abc", "bac", false},
    {"abc", "cab", true},
}

func TestIsRotation(t *testing.T) {
    for index, data := range testData {
        if IsRotation(data.input, data.test) != data.result {
            t.Errorf("Index: %d, Input: \"%s\", Test string: \"%s\", Expected Result: %t", index, data.input,
                data.test, data.result)
        }
    }
}