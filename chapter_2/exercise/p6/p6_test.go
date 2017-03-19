package p6

import (
    "github.com/DheerendraRathor/ctci-go/chapter_2"
    "testing"
)

type testStruct struct {
    input *chapter_2.LinkedList
    isPalindrome bool
}

var testData = []testStruct {
    {
        chapter_2.SliceToLinkedList([]interface{}{}),
        true,
    },
    {
        chapter_2.SliceToLinkedList([]interface{}{1}),
        true,
    },
    {
        chapter_2.SliceToLinkedList([]interface{}{1, 2, 1}),
        true,
    },
    {
        chapter_2.SliceToLinkedList([]interface{}{1, 2, 2, 1}),
        true,
    },
    {
        chapter_2.SliceToLinkedList([]interface{}{1, 2, 3}),
        false,
    },
    {
        chapter_2.SliceToLinkedList([]interface{}{1, 2, 3, 1}),
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
