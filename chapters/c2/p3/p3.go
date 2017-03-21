package p3

import "github.com/DheerendraRathor/ctci-go/collections"

func DeleteMiddleNode(list *collections.LinkedList, node *collections.Node)  {
    list.DeleteNodeByRef(node)
}
