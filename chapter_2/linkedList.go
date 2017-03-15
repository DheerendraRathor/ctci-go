package chapter_2

import (
    "bytes"
    "fmt"
)

type Node struct {
    Value interface{}
    Next *Node
    Prev *Node
}

type LinkedList struct {
    Head *Node
}

func (l *LinkedList) AddNode(val interface{}) {
    if l.Head == nil {
        l.Head = &Node{Value: val}
        return
    }

    currentNode := l.Head
    for currentNode.Next != nil {
        currentNode = currentNode.Next
    }

    currentNode.Next = &Node{Value: val, Prev: currentNode}
}

func (l *LinkedList) AddNodesBySlice(values []interface{}) {
    if len(values) == 0 {
        return
    }

    if l.Head == nil {
        l.Head = new(Node)
        l.Head.Value = values[0]
    }

    lastNode := l.Head
    for lastNode.Next != nil {
        lastNode = lastNode.Next
    }

    for i := 1; i < len(values); i++ {
        temp := &Node{Value: values[i], Prev: lastNode}
        lastNode.Next = temp
        lastNode = temp
    }
}

func (l *LinkedList) String() string {

    if l.Head == nil {
        return ""
    }

    var buffer bytes.Buffer

    currentNode := l.Head
    for currentNode != nil {
        buffer.WriteString(fmt.Sprintf("%v", currentNode.Value))
        currentNode = currentNode.Next
    }

    return buffer.String()
}

// Caveat: Node must be part of list. It should not be a rogue node otherwise things will be messy
func (l *LinkedList) DeleteNodeByRef(n *Node)  {
    if n == nil {
        return
    }

    if n.Prev != nil {
        n.Prev.Next = n.Next
        if n.Next != nil {
            n.Next.Prev = n.Prev
        }
    } else {        // It means that node is head of list
        l.Head = n.Next
        if n.Next != nil {
            n.Next.Prev = nil
        }
    }
}

func (l *LinkedList) Equals(l1 *LinkedList) bool {
    if l.Head == nil && l1.Head == nil {
        return true
    }

    currentLNode := l.Head
    currentL1Node := l1.Head

    for true {
        if currentLNode == nil && currentL1Node == nil {
            return true
        }

        if currentLNode != nil &&
            currentL1Node != nil &&
            currentLNode.Value != currentL1Node.Value {
            return false
        } else if currentLNode == nil ||
            currentL1Node == nil {
            return false
        }

        currentLNode = currentLNode.Next
        currentL1Node = currentL1Node.Next
    }

    return true
}

func SliceToLinkedList(values []interface{}) *LinkedList {
    list := new(LinkedList)
    list.AddNodesBySlice(values)
    return list
}
