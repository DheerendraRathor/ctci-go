package p6

import "github.com/DheerendraRathor/ctci-go/chapter_2"

func IsLinkedListPalindrome(list chapter_2.LinkedList) bool {

    if list.Head == nil {
        return true
    }

    lastNode := list.Last()
    startNode := list.Head
    isPalindrome := true

    // For odd size list, startNode and lastNode will become same at middle.
    // For even size list, both will cross each other and just after crossing startNode will be the nextNode of lastNode
    for lastNode != startNode && lastNode.Next != startNode{
        if lastNode.Value != startNode.Value {
            isPalindrome = false
            break
        }

        startNode = startNode.Next
        lastNode = lastNode.Prev
    }

    return isPalindrome
}
