package p4

import (
    "github.com/DheerendraRathor/ctci-go/chapter_2"
    "testing"
)

type  testStruct struct{
    input *chapter_2.LinkedList
    pivot int
    output *chapter_2.LinkedList
}

var testData = []testStruct {
    {
        chapter_2.SliceToLinkedList([]interface{}{}),
        2,
        chapter_2.SliceToLinkedList([]interface{}{}),
    },
    {
        chapter_2.SliceToLinkedList([]interface{}{1, 4, 5, 2, 3, -1, 0}),
        2,
        chapter_2.SliceToLinkedList([]interface{}{1, -1, 0, 4, 5, 2, 3}),
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
