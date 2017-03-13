package p6

import "testing"

type testStruct struct {
    input string
    output string
}


var testData []testStruct = []testStruct {
    {"", ""},
    {"a", "a"},
    {"ab", "ab"},
    {"abb", "abb"},
    {"abbbbbc", "a1b5c1"},
    {"aabcccccaaa", "a2b1c5a3"},
}

func TestCompresser(t *testing.T) {
    for index, data := range testData {
        actualResult := Compressor(data.input)
        if actualResult != data.output {
            t.Errorf("Index: %d Input: %s Output: %s ExpectedOutput: %s", index, data.input, actualResult,
                data.output)
        }
    }
}