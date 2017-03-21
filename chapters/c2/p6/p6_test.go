package p6

import (
    "github.com/DheerendraRathor/ctci-go/collections"
    "testing"
)

type testStruct struct {
    input *collections.LinkedList
    isPalindrome bool
}

var testData = []testStruct {
    {
        collections.SliceToLinkedList([]interface{}{}),
        true,
    },
    {
        collections.SliceToLinkedList([]interface{}{1}),
        true,
    },
    {
        collections.SliceToLinkedList([]interface{}{1, 2, 1}),
        true,
    },
    {
        collections.SliceToLinkedList([]interface{}{1, 2, 2, 1}),
        true,
    },
    {
        collections.SliceToLinkedList([]interface{}{1, 2, 3}),
        false,
    },
    {
        collections.SliceToLinkedList([]interface{}{1, 2, 3, 1}),
        false,
    },
}

func TestIsLinkedListPalindrome(t *testing.T) {
    for index, data := range testData {
        actualResult := IsLinkedListPalindrome(*data.input)
        if actualResult != data.isPalindrome {
            t.Errorf("For index: %d, Expected Resutl: %t, Got Result : %t", index, data.isPalindrome,
                actualResult)
        }
    }
}
