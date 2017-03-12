package exercise

import "testing"

type TestStr struct {
    str string
    isUnique bool
}

var strings []TestStr = []TestStr{
    {"", true},
    {"Hello", false},
    {"Bye, World", true},
    {"ʥቕॵཚچ", true},
    {"ʥቕॵཚچچ", false},
}

func TestIsUnique(t *testing.T) {
    for _, str := range strings {
        if IsUnique(str.str) != str.isUnique {
            t.Error("string", str.str, "is expected isUnique =", str.isUnique, "but found otherwise")
        }
    }
}

func TestIsUniqueWithoutDS(t *testing.T) {
    for _, str := range strings {
        if IsUnique(str.str) != str.isUnique {
            t.Error("string", str.str, "is expected isUnique =", str.isUnique, "but found otherwise")
        }
    }
}