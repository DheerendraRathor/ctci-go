package trees

import (
    "github.com/DheerendraRathor/ctci-go/utils"
)

type BinaryTree struct {
    Value interface{}
    Left *BinaryTree
    Right *BinaryTree
    height int
    isHeightPopulated bool          // This flag will be true once height is populated for all nodes in tree.
}

func (bt *BinaryTree) Height() int {
    if bt.isHeightPopulated {
        return bt.height
    }

    leftHeight, rightHeight := 0, 0
    if bt.Left != nil {
        leftHeight = bt.Left.Height()
    }
    if bt.Right != nil {
        rightHeight = bt.Right.Height()
    }

    bt.height = 1 + utils.Max(leftHeight, rightHeight)

    bt.isHeightPopulated = true
    return bt.height
}

func (bt *BinaryTree) IsComplete() bool {

    isLeftComplete, isRightComplete := true, true
    leftHeight, rightHeight := 0, 0
    if bt.Left != nil {
        isLeftComplete = bt.Left.IsComplete()
        leftHeight = bt.Left.Height()
    }
    if bt.Right != nil {
        isRightComplete = bt.Right.IsComplete()
        rightHeight = bt.Right.Height()
    }

    return isLeftComplete && isRightComplete && (leftHeight == rightHeight || leftHeight == rightHeight + 1)
}

func (bt *BinaryTree) IsFull() bool {
    if bt.Left == nil && bt.Right == nil {
        return true
    }

    if bt.Left != nil && bt.Right != nil {
        return bt.Left.IsFull() && bt.Right.IsFull()
    }

    return false
}

func (bt *BinaryTree) IsPerfect() bool {
    return isPerfectHelper(bt, bt.Height())
}

func (bt *BinaryTree) PreOrderTraversal() []*BinaryTree {
    var output []*BinaryTree
    preOrderTraversalHelper(bt, &output)
    return output
}

func (bt *BinaryTree) InOrderTraversal() []*BinaryTree {
    var output []*BinaryTree
    inOrderTraversalHelper(bt, &output)
    return output
}

func (bt *BinaryTree) PostOrderTraversal() []*BinaryTree {
    var output []*BinaryTree
    postOrderTraversalHelper(bt, &output)
    return output
}

func isPerfectHelper(bt *BinaryTree, depth int) bool {

    if bt == nil {
        if depth == 0 {
            return true
        } else {
            return false
        }
    }

    return isPerfectHelper(bt.Left, depth - 1) && isPerfectHelper(bt.Right, depth - 1)
}

func preOrderTraversalHelper(bt *BinaryTree, output *[]*BinaryTree) {
    if bt == nil {
        return
    }
    *output = append(*output, bt)
    preOrderTraversalHelper(bt.Left, output)
    preOrderTraversalHelper(bt.Right, output)
}

func inOrderTraversalHelper(bt *BinaryTree, output *[]*BinaryTree) {
    if bt == nil {
        return
    }
    inOrderTraversalHelper(bt.Left, output)
    *output = append(*output, bt)
    inOrderTraversalHelper(bt.Right, output)
}

func postOrderTraversalHelper(bt *BinaryTree, output *[]*BinaryTree) {
    if bt == nil {
        return
    }
    postOrderTraversalHelper(bt.Left, output)
    postOrderTraversalHelper(bt.Right, output)
    *output = append(*output, bt)
}
