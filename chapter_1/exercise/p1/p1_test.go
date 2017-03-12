package p1

import "testing"

type uniqueTestStr struct {
    str string
    isUnique bool
}

var uniqueTestStrs []uniqueTestStr = []uniqueTestStr{
    {"", true},
    {"Hello", false},
    {"Bye, World", true},
    {"ʥቕॵཚچ", true},
    {"ʥቕॵཚچچ", false},
}

func TestIsUnique(t *testing.T) {
    for _, str := range uniqueTestStrs {
        if IsUnique(str.str) != str.isUnique {
            t.Error("string", str.str, "is expected isUnique =", str.isUnique, "but found otherwise")
        }
    }
}

func TestIsUniqueWithoutDS(t *testing.T) {
    for _, str := range uniqueTestStrs {
        if IsUniqueWithoutDS(str.str) != str.isUnique {
            t.Error("string", str.str, "is expected isUnique =", str.isUnique, "but found otherwise")
        }
    }
}