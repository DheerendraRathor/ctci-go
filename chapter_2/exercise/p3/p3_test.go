package p3

import (
    "github.com/DheerendraRathor/ctci-go/collections"
    "testing"
)

func TestDeleteMiddleNode(t *testing.T) {
    var list1 = collections.SliceToLinkedList([]interface{}{1, 2, 3, 4})
    DeleteMiddleNode(list1, list1.Head.Next.Next)     // Delete 3
    if "124" != list1.String() {
        t.Errorf("Expected output: 124. Got: %s", list1.String())
    }
}
