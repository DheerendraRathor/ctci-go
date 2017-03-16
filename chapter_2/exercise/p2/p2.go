package p2

import (
    "github.com/DheerendraRathor/ctci-go/chapter_2"
    "fmt"
)

// Since problem asked about single linked list we won't be using Prev to access anything
// Return k th element from last for LinkedList list
// k should be greater than or equal to 1
func GetKthFromLast(list *chapter_2.LinkedList, k int32) interface{} {
    totalCount := LengthOfList(list)
    kthFromLast, err := GetIthElement(list, totalCount-k )
    if err != nil {
        panic(fmt.Sprintf("Got Err :%v", err))
    }
    return kthFromLast
}

func LengthOfList(list *chapter_2.LinkedList) int32 {
    var count int32 = 0
    currentNode := list.Head
    for currentNode != nil {
        count++
        currentNode = currentNode.Next
    }

    return count
}

// Return i th element from a LinkedList list
// In this case counting starts with 1
func GetIthElement(list *chapter_2.LinkedList, i int32) (interface{}, error) {

    if i < 0 {
        return -1, fmt.Errorf("i is negative\n")
    }

    var count int32 = 0
    currentNode := list.Head
    for count < i {
        if currentNode == nil {
            return -1, fmt.Errorf("i is less than length of list\n")
        }
        count++
        currentNode = currentNode.Next
    }

    return currentNode.Value, nil
}
