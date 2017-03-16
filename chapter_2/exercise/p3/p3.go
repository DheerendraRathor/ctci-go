package p3

import "github.com/DheerendraRathor/ctci-go/chapter_2"

func DeleteMiddleNode(list *chapter_2.LinkedList, node *chapter_2.Node)  {
    list.DeleteNodeByRef(node)
}
