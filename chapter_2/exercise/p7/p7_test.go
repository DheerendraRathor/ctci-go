package p7

import (
    "github.com/DheerendraRathor/ctci-go/collections"
    "testing"
)

var list1 = *collections.SliceToLinkedList([]interface{}{1, 2, 3, 4, 5})
var list2 = *collections.SliceToLinkedList([]interface{}{1, 2})

var list3 = *collections.SliceToLinkedList([]interface{}{1, 2, 9})

var list5 = *collections.SliceToLinkedList([]interface{}{})
var list6 = *collections.SliceToLinkedList([]interface{}{})

func TestGetIntersectionOfLists(t *testing.T) {
    emptyListIntersection := GetIntersectionOfLists(list5, list6)
    if emptyListIntersection != nil {
        t.Error("Intersection of empty list is not null")
    }

    emptyNonEmptyIntersection := GetIntersectionOfLists(list3, list5)
    if emptyNonEmptyIntersection != nil {
        t.Error("Intersection of empty and non empty list is not null")
    }

    nonIntersectingListIntersection := GetIntersectionOfLists(list1, list2)
    if nonIntersectingListIntersection != nil {
        t.Error("Intersection of non intersecting list is not null")
    }

    list3.Last().Next = list1.Head.Next.Next

    intersection := GetIntersectionOfLists(list1, list3)
    if intersection != list1.Head.Next.Next {
        t.Errorf("Intersection is not what I expected. Expected: %v, got: %v", list1.Head.Next.Next,
            intersection)
    }

}
