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

func (l *LinkedList) AddNode(val interface{}) *Node {
    if l.Head == nil {
        l.Head = &Node{Value: val}
        return l.Head
    }

    currentNode := l.Head
    for currentNode.Next != nil {
        currentNode = currentNode.Next
    }

    currentNode.Next = &Node{Value: val, Prev: currentNode}
    return currentNode.Next
}

func (l *LinkedList) AddNodeToFront(val interface{}) *Node {
    newNode := &Node{Value: val, Next: l.Head}

    if l.Head != nil {
        l.Head.Prev = newNode
    }

    l.Head = newNode

    return newNode
}

func (l *LinkedList) AddNodesBySlice(values []interface{}) {
    if len(values) == 0 {
        return
    }

    startIndex := 0
    if l.Head == nil {
        l.Head = new(Node)
        l.Head.Value = values[0]
        startIndex = 1
    }

    lastNode := l.Head
    for lastNode.Next != nil {
        lastNode = lastNode.Next
    }

    for i := startIndex; i < len(values); i++ {
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

    areEqual := true

    for true {
        if currentLNode == nil && currentL1Node == nil {
            areEqual = true
            break
        }

        if currentLNode != nil &&
            currentL1Node != nil &&
            currentLNode.Value != currentL1Node.Value {
            areEqual = false
            break
        } else if currentLNode == nil ||
            currentL1Node == nil {
            areEqual = false
            break
        }

        currentLNode = currentLNode.Next
        currentL1Node = currentL1Node.Next
    }

    return areEqual
}

// Returns last node of list l
// If head is nil then it returns nil.
func (l *LinkedList) Last() *Node {
    currentNode := l.Head
    lastNode := l.Head

    for currentNode != nil {
        lastNode = currentNode
        currentNode = currentNode.Next
    }
    return lastNode
}

func SliceToLinkedList(values []interface{}) *LinkedList {
    list := new(LinkedList)
    list.AddNodesBySlice(values)
    return list
}


