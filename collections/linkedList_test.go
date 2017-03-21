package collections

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

func TestLinkedList_AddNodeToFront(t *testing.T) {
    list := new(LinkedList)

    for i := 0; i < 5; i++ {
        list.AddNodeToFront(i)
    }

    stringVal := list.String()
    if "43210" != stringVal {
        t.Errorf("Expected: 43210, Got: %v", stringVal)
    }
}

func TestLinkedList_AddNodesBySlice(t *testing.T) {
    list := new(LinkedList)

    list.AddNodesBySlice([]interface{}{})
    stringVal := list.String()
    if "" != stringVal {
        t.Errorf("Expected: \"\", Got: %v", stringVal)
    }

    list.AddNodesBySlice([]interface{}{0, 5, 6})
    stringVal = list.String()
    if "056" != stringVal {
        t.Errorf("Expected: 056, Got: %v", stringVal)
    }

    list.AddNodesBySlice([]interface{}{1, 2, 3, 4, "Hi"})
    stringVal = list.String()
    if "0561234Hi" != stringVal {
        t.Errorf("Expected: 0561234Hi, Got: %v", stringVal)
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

    list.DeleteNodeByRef(list.Head)
    stringVal = list.String()
    if "34Hi" != stringVal {
        t.Errorf("Expected: 134Hi, Got: %v", stringVal)
    }

    list.DeleteNodeByRef(nil)
    stringVal = list.String()
    if "34Hi" != stringVal {
        t.Errorf("Expected: 134Hi, Got: %v", stringVal)
    }
}

func TestLinkedList_Equals(t *testing.T) {
    list1 := new(LinkedList)
    list2 := new(LinkedList)

    if !list1.Equals(list2) {
        t.Error("List1 and List2 are not equal")
    }

    if !list2.Equals(list1) {
        t.Error("List2 and List1 are not equal")
    }

    sameList := []interface{}{0, 1, 2}
    list1.AddNodesBySlice(sameList)
    list2.AddNodesBySlice(sameList)

    if !list1.Equals(list2) {
        t.Error("List1 and List2 are not equal")
    }

    if !list2.Equals(list1) {
        t.Error("List2 is not equal to List1")
    }

    list1.AddNodesBySlice([]interface{}{4, 5})
    list2.AddNode(4)
    if list1.Equals(list2) {
        t.Error("List1 and List2 are equal, which should not be")
    }

    if list2.Equals(list1) {
        t.Error("List2 is equal to List1 which should not be")
    }

    list2.AddNode(6)
    if list1.Equals(list2) {
        t.Error("List1 and List2 are equal, which should not be")
    }

    if list2.Equals(list1) {
        t.Error("List2 is equal to List1 which should not be")
    }
}

func TestSliceToLinkedList(t *testing.T) {
    list := SliceToLinkedList([]interface{}{1, 2})
    stringVal := list.String()
    if stringVal != "12" {
        t.Errorf("Expected: 12, Got: %v", stringVal)
    }
}

func TestLinkedList_Last(t *testing.T) {
    nullList := SliceToLinkedList([]interface{}{})
    lastNode := nullList.Last()
    if lastNode != nil {
        t.Error("Expected last node to be nil, but it isn't")
    }

    headOnlyList := SliceToLinkedList([]interface{}{1})
    lastNode = headOnlyList.Last()
    if lastNode != headOnlyList.Head {
        t.Error("Last node is not head in case of headonly list")
    }

    someElementList := SliceToLinkedList([]interface{}{1, 2, 3})
    lastNode = someElementList.Last()
    if lastNode.Value != 3 {
        t.Error("Last node is not 3 in someElementList")
    }
}

func TestLinkedList_Length(t *testing.T) {
    list := new(LinkedList)
    if list.Length() != 0 {
        t.Error("Length of empty list is not 0")
    }

    list.AddNode(10)
    list.AddNode(20)
    if list.Length() != 2 {
        t.Error("Expected Length is not 2")
    }

    list.DeleteNodeByRef(list.Head)
    if list.Length() != 1 {
        t.Error("Expected Length is not 1")
    }
}
