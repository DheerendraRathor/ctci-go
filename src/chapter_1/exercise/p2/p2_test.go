package p2

import "testing"

type permTestStr struct {
    str1 string
    str2 string
    isPerm bool
}

var permTestStrings []permTestStr = []permTestStr {
    {"", "", true},
    {"", "a", false},
    {"Hello", "oHell", true},
    {"Hello", "hola", false},
    {"Hello", "World", false},
    {"ʥቕॵཚچ", "ቕॵཚچʥ", true},
    {"ʥቕॵཚچ", "ʥॵཚچ", false},
}

func TestArePermutations(t *testing.T) {
    for _, testData := range permTestStrings {
        if ArePermutations(testData.str1, testData.str2) != testData.isPerm {
            t.Error("Strings", testData.str1, "and", testData.str2, "are supposed to be isPerm =",
                testData.isPerm, "but found otherwise")
        }
    }
}