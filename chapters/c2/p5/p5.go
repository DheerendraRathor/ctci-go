package p5

import "github.com/DheerendraRathor/ctci-go/collections"

func ReverseOrderLinkedListAddition(list1 collections.LinkedList, list2 collections.LinkedList) *collections.LinkedList {
    outputList := new(collections.LinkedList)

    currentList1Node := list1.Head
    currentList2Node := list2.Head
    carry := 0

    for true {
        if currentList1Node == nil && currentList2Node == nil {
            break
        }

        sum := carry
        if currentList1Node != nil {
            sum += currentList1Node.Value.(int)
            currentList1Node = currentList1Node.Next
        }

        if currentList2Node != nil {
            sum += currentList2Node.Value.(int)
            currentList2Node = currentList2Node.Next
        }

        carry = sum / 10
        outputList.AddNode(sum % 10)
    }

    if carry != 0 {
        outputList.AddNode(carry)
    }

    return outputList
}

func ForwardOrderLinkedListAddition(list1 collections.LinkedList, list2 collections.LinkedList) *collections.LinkedList {
    outputList := new(collections.LinkedList)

    currentList1Node := list1.Last()
    currentList2Node := list2.Last()
    carry := 0

    for true {
        if currentList1Node == nil && currentList2Node == nil {
            break
        }

        sum := carry
        if currentList1Node != nil {
            sum += currentList1Node.Value.(int)
            currentList1Node = currentList1Node.Prev
        }

        if currentList2Node != nil {
            sum += currentList2Node.Value.(int)
            currentList2Node = currentList2Node.Prev
        }

        carry = sum / 10
        outputList.AddNodeToFront(sum % 10)
    }

    if carry != 0 {
        outputList.AddNodeToFront(carry)
    }

    return outputList
}
