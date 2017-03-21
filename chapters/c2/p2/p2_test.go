package p2

import (
    "github.com/DheerendraRathor/ctci-go/collections"
    "testing"
)

type testStruct struct {
    list *collections.LinkedList
    k int32
    output interface{}
}

var testData []testStruct = []testStruct {
    {
        collections.SliceToLinkedList([]interface{}{1, 2, 3, "Hi"}),
        2,
        3,
    },
    {
        collections.SliceToLinkedList([]interface{}{"Go", "is", "Awesome"}),
        1,
        "Awesome",
    },
}

func TestGetKthFromLast(t *testing.T) {
    for index, data := range testData {
        actualResult := GetKthFromLast(data.list, data.k)
        if actualResult != data.output {
            t.Errorf("Index: %d Expected Result: %v Actual Result: %v", index, data.output, actualResult)
        }
    }
}

func TestGetKthFromLast2_Panic(t *testing.T) {
    defer func() {
        if r := recover(); r == nil {
            t.Error("Panic not recovered")
        }
    }()

    list := collections.SliceToLinkedList([]interface{}{1, 2})
    GetKthFromLast(list, 3)
}

func TestGetKthFromLast2_Panic2(t *testing.T) {
    defer func() {
        if r := recover(); r == nil {
            t.Error("Panic not recovered")
        }
    }()

    list := collections.SliceToLinkedList([]interface{}{1, 2})
    GetKthFromLast(list, -1)
}
