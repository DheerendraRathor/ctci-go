package chapter_2

import "testing"

func TestLinkedList_AddNode(t *testing.T) {
    list := new(LinkedList)
    for i := 0; i < 4; i++  {
        list.AddNode(i)
    }

    stringVal := list.String()
    if "0123" != stringVal {
        t.Errorf("Expected: 0123, Got: %v", stringVal)
    }
}

func TestLinkedList_AddNodesBySlice(t *testing.T) {
    list := new(LinkedList)

    list.AddNodesBySlice([]interface{}{1, 2, 3, 4, "Hi"})
    stringVal := list.String()
    if "1234Hi" != stringVal {
        t.Errorf("Expected: 1234Hi, Got: %v", stringVal)
    }
}

func TestLinkedList_DeleteNodeByRef(t *testing.T) {
    list := new(LinkedList)
    list.AddNodesBySlice([]interface{}{1, 2, 3, 4, "Hi"})

    list.DeleteNodeByRef(list.Head.Next)
    stringVal := list.String()
    if "134Hi" != stringVal {
        t.Errorf("Expected: 134Hi, Got: %v", stringVal)
    }
}
