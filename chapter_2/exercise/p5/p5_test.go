package p5

import (
    "github.com/DheerendraRathor/ctci-go/collections"
    "testing"
)

type testStruct struct {
    list1 collections.LinkedList
    list2 collections.LinkedList
    reverseOutput collections.LinkedList
    forwardOutput collections.LinkedList
}

var testData = []testStruct {
    {
        *collections.SliceToLinkedList([]interface{}{}),
        *collections.SliceToLinkedList([]interface{}{}),
        *collections.SliceToLinkedList([]interface{}{}),
        *collections.SliceToLinkedList([]interface{}{}),
    },
    {
        *collections.SliceToLinkedList([]interface{}{1}),
        *collections.SliceToLinkedList([]interface{}{}),
        *collections.SliceToLinkedList([]interface{}{1}),
        *collections.SliceToLinkedList([]interface{}{1}),
    },
    {
        *collections.SliceToLinkedList([]interface{}{1, 2}),
        *collections.SliceToLinkedList([]interface{}{3, 4, 5}),
        *collections.SliceToLinkedList([]interface{}{4, 6, 5}),
        *collections.SliceToLinkedList([]interface{}{3, 5, 7}),
    },
    {
        *collections.SliceToLinkedList([]interface{}{1, 2}),
        *collections.SliceToLinkedList([]interface{}{9, 7, 5}),
        *collections.SliceToLinkedList([]interface{}{0, 0, 6}),
        *collections.SliceToLinkedList([]interface{}{9, 8, 7}),
    },
    {
        *collections.SliceToLinkedList([]interface{}{1, 2}),
        *collections.SliceToLinkedList([]interface{}{9, 7}),
        *collections.SliceToLinkedList([]interface{}{0, 0, 1}),
        *collections.SliceToLinkedList([]interface{}{1, 0, 9}),
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
