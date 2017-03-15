package p1

import (
    "testing"
    "github.com/DheerendraRathor/ctci-go/chapter_2"
    "fmt"
)

type testStruct struct {
    input *chapter_2.LinkedList
    output *chapter_2.LinkedList
}

var testData []testStruct = []testStruct {
    {
        chapter_2.SliceToLinkedList([]interface{}{}),
        chapter_2.SliceToLinkedList([]interface{}{}),
    },
    {
        chapter_2.SliceToLinkedList([]interface{}{1}),
        chapter_2.SliceToLinkedList([]interface{}{1}),
    },
    {
        chapter_2.SliceToLinkedList([]interface{}{1, 2, 1, 6, 3, 4, 2, 3, 6}),
        chapter_2.SliceToLinkedList([]interface{}{1, 2, 6, 3, 4}),
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
