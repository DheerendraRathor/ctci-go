package p7

import "github.com/DheerendraRathor/ctci-go/chapter_2"

// Caveats:
// If their is actually an intersection, then end node will be the same for both lists
// Since we're dealing with singly linked lists, we'll be using only Next to solve this issue
func GetIntersectionOfLists(list1 chapter_2.LinkedList, list2 chapter_2.LinkedList) *chapter_2.Node {

    // Last is implemented only by using Next
    list1LastNode := list1.Last()
    list2LastNode := list2.Last()

    if list1LastNode != list2LastNode || list1LastNode == nil || list2LastNode == nil{
        return nil
    }

    list1Length := list1.Length()
    list2Length := list2.Length()

    listToFastForward := list1
    otherList := list2
    stepsToForward := list1Length - list2Length
    if list2Length > list1Length {
        listToFastForward = list2
        otherList = list1
        stepsToForward = list2Length - list1Length
    }

    ffListCurrentNode := listToFastForward.Head
    otherListCurrentNode := otherList.Head
    for i := 0; i < stepsToForward; i++ {
        ffListCurrentNode = ffListCurrentNode.Next
    }

    var intersection *chapter_2.Node
    for true {
        if ffListCurrentNode == otherListCurrentNode {
            intersection = ffListCurrentNode
            break
        }
        ffListCurrentNode = ffListCurrentNode.Next
        otherListCurrentNode = otherListCurrentNode.Next
    }

    return intersection
}
