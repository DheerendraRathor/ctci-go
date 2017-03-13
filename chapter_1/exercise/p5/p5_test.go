package p5

import "testing"

type testStruct struct {
    input string
    output string
    isOneAway bool
}

var testData []testStruct = []testStruct {
    {"", "", true},
    {"", "a", true},
    {"a", "bc", false},
    {"pale", "ple", true},
    {"pales", "pale", true},
    {"pale", "bale", true},
    {"pale", "bake", false},
    {"plu", "plur", true},
    {"ʥቕॵཚچ", "ʥʥॵཚچ", true},
    {"ʥቕॵཚچ", "ʥچʥཚچ", false},
    {"abcd", "ab", false},
}

func TestIsOneAway(t *testing.T) {
    for index, data := range testData {
        actualResult := IsOneAway(data.input, data.output)
        if actualResult != data.isOneAway {
            t.Errorf("For index:%d input:\"%s\" output:\"%s\" expected result was %t but got %t", index,
                data.input, data.output, data.isOneAway, actualResult)
        }
    }
}
