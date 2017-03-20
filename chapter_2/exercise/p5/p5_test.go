package p5

import (
    "github.com/DheerendraRathor/ctci-go/chapter_2"
    "testing"
)

type testStruct struct {
    list1 chapter_2.LinkedList
    list2 chapter_2.LinkedList
    reverseOutput chapter_2.LinkedList
    forwardOutput chapter_2.LinkedList
}

var testData = []testStruct {
    {
        *chapter_2.SliceToLinkedList([]interface{}{}),
        *chapter_2.SliceToLinkedList([]interface{}{}),
        *chapter_2.SliceToLinkedList([]interface{}{}),
        *chapter_2.SliceToLinkedList([]interface{}{}),
    },
    {
        *chapter_2.SliceToLinkedList([]interface{}{1}),
        *chapter_2.SliceToLinkedList([]interface{}{}),
        *chapter_2.SliceToLinkedList([]interface{}{1}),
        *chapter_2.SliceToLinkedList([]interface{}{1}),
    },
    {
        *chapter_2.SliceToLinkedList([]interface{}{1, 2}),
        *chapter_2.SliceToLinkedList([]interface{}{3, 4, 5}),
        *chapter_2.SliceToLinkedList([]interface{}{4, 6, 5}),
        *chapter_2.SliceToLinkedList([]interface{}{3, 5, 7}),
    },
    {
        *chapter_2.SliceToLinkedList([]interface{}{1, 2}),
        *chapter_2.SliceToLinkedList([]interface{}{9, 7, 5}),
        *chapter_2.SliceToLinkedList([]interface{}{0, 0, 6}),
        *chapter_2.SliceToLinkedList([]interface{}{9, 8, 7}),
    },
    {
        *chapter_2.SliceToLinkedList([]interface{}{1, 2}),
        *chapter_2.SliceToLinkedList([]interface{}{9, 7}),
        *chapter_2.SliceToLinkedList([]interface{}{0, 0, 1}),
        *chapter_2.SliceToLinkedList([]interface{}{1, 0, 9}),
    },
}

func TestReverseOrderLinkedListAddition(t *testing.T) {
    for index, data := range testData {
        actualOutput := ReverseOrderLinkedListAddition(data.list1, data.list2)
        if !actualOutput.Equals(&data.reverseOutput) {
            t.Errorf("Index: %d failed", index)
        }
    }
}

func TestForwardOrderLinkedListAddition(t *testing.T) {
    for index, data := range testData {
        actualOutput := ForwardOrderLinkedListAddition(data.list1, data.list2)
        if !actualOutput.Equals(&data.forwardOutput) {
            t.Errorf("Index: %d failed", index)
        }
    }
}
