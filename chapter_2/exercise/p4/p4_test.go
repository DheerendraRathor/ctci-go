package p4

import (
    "github.com/DheerendraRathor/ctci-go/collections"
    "testing"
)

type  testStruct struct{
    input *collections.LinkedList
    pivot int
    output *collections.LinkedList
}

var testData = []testStruct {
    {
        collections.SliceToLinkedList([]interface{}{}),
        2,
        collections.SliceToLinkedList([]interface{}{}),
    },
    {
        collections.SliceToLinkedList([]interface{}{1, 4, 5, 2, 3, -1, 0}),
        2,
        collections.SliceToLinkedList([]interface{}{1, -1, 0, 4, 5, 2, 3}),
    },
}

func TestPartitionOfList(t *testing.T) {
    for index, data := range testData {
        actualResult := PartitionOfList(data.input, data.pivot)
        if !actualResult.Equals(data.output) {
            t.Errorf("For index %d, got output: %s", index, actualResult.String())
        }
    }
}
