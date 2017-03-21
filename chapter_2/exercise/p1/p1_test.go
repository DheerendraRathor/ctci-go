package p1

import (
    "testing"
    "github.com/DheerendraRathor/ctci-go/collections"
    "fmt"
)

type testStruct struct {
    input *collections.LinkedList
    output *collections.LinkedList
}

var testData []testStruct = []testStruct {
    {
        collections.SliceToLinkedList([]interface{}{}),
        collections.SliceToLinkedList([]interface{}{}),
    },
    {
        collections.SliceToLinkedList([]interface{}{1}),
        collections.SliceToLinkedList([]interface{}{1}),
    },
    {
        collections.SliceToLinkedList([]interface{}{1, 2, 1, 6, 3, 4, 2, 3, 6}),
        collections.SliceToLinkedList([]interface{}{1, 2, 6, 3, 4}),
    },
}

func TestRemoveDuplicates(t *testing.T) {
    for index, data := range testData {
        RemoveDuplicates(data.input)
        fmt.Printf("Index: %d\n", index)
        if !data.input.Equals(data.output) {
            t.Errorf("Index: %d, Actual Result: %s Expected Result: %s", index, data.input, data.output)
        }
    }
}
