package collections

import (
    "testing"
    "github.com/stretchr/testify/assert"
)

func TestBinaryTree_Height(t *testing.T) {
    bt := new(BinaryTree)
    bt.Left = new(BinaryTree)
    bt.Right = new(BinaryTree)

    height := bt.Height()
    assert.Equal(t, 2, height)
    assert.True(t, bt.isHeightPopulated)

    assert.Equal(t, 2, bt.Height())
}

func TestBinaryTree_IsComplete(t *testing.T) {
    bt1 := new(BinaryTree)
    assert.True(t, bt1.IsComplete())

    bt2 := new(BinaryTree)
    bt2.Right = new(BinaryTree)
    assert.False(t, bt2.IsComplete())

    bt3 := new(BinaryTree)
    bt3.Left = new(BinaryTree)
    assert.True(t, bt3.IsComplete())

    bt4 := new(BinaryTree)
    bt4.Left = new(BinaryTree)
    bt4.Right = new(BinaryTree)
    assert.True(t, bt4.IsComplete())
}

func TestBinaryTree_IsFull(t *testing.T) {
    bt1 := new(BinaryTree)
    assert.True(t, bt1.IsFull())

    bt2 := new(BinaryTree)
    bt2.Right = new(BinaryTree)
    assert.False(t, bt2.IsFull())

    bt3 := new(BinaryTree)
    bt3.Left = new(BinaryTree)
    assert.False(t, bt3.IsFull())

    bt4 := new(BinaryTree)
    bt4.Left = new(BinaryTree)
    bt4.Right = new(BinaryTree)
    assert.True(t, bt4.IsFull())

    bt5 := new(BinaryTree)
    bt5.Left = new(BinaryTree)
    bt5.Right = new(BinaryTree)
    bt5.Right.Left = new(BinaryTree)
    bt5.Right.Right = new(BinaryTree)
    assert.True(t, bt5.IsFull())
}

func TestBinaryTree_IsPerfect(t *testing.T) {
    bt1 := new(BinaryTree)
    assert.True(t, bt1.IsPerfect())

    bt2 := new(BinaryTree)
    bt2.Right = new(BinaryTree)
    assert.False(t, bt2.IsPerfect())

    bt3 := new(BinaryTree)
    bt3.Left = new(BinaryTree)
    assert.False(t, bt3.IsPerfect())

    bt4 := new(BinaryTree)
    bt4.Left = new(BinaryTree)
    bt4.Right = new(BinaryTree)
    assert.True(t, bt4.IsPerfect())

    bt5 := new(BinaryTree)
    bt5.Left = new(BinaryTree)
    bt5.Right = new(BinaryTree)
    bt5.Right.Left = new(BinaryTree)
    bt5.Right.Right = new(BinaryTree)
    assert.False(t, bt5.IsPerfect())
}

func TestBinaryTree_PreOrderTraversal(t *testing.T) {
    bt1 := BinaryTree{Value:1}
    output := bt1.PreOrderTraversal()
    validateBinaryTreeTraversal(t, output, []interface{}{1})

    bt2 := BinaryTree{Value: 1}
    bt2.Right = &BinaryTree{Value: 2}
    output = bt2.PreOrderTraversal()
    validateBinaryTreeTraversal(t, output, []interface{}{1, 2})

    bt3 := BinaryTree{Value: 1}
    bt3.Left = &BinaryTree{Value: 2}
    output = bt3.PreOrderTraversal()
    validateBinaryTreeTraversal(t, output, []interface{}{1, 2})

    bt4 := BinaryTree{Value: 1}
    bt4.Left = &BinaryTree{Value: 2}
    bt4.Right = &BinaryTree{Value: 3}
    output = bt4.PreOrderTraversal()
    validateBinaryTreeTraversal(t, output, []interface{}{1, 2, 3})

    bt5 := BinaryTree{Value: 1}
    bt5.Left = &BinaryTree{Value: 2}
    bt5.Right = &BinaryTree{Value: 3}
    bt5.Right.Left = &BinaryTree{Value: 4}
    bt5.Right.Right = &BinaryTree{Value: 5}
    output = bt5.PreOrderTraversal()
    validateBinaryTreeTraversal(t, output, []interface{}{1, 2, 3, 4, 5})
}

func TestBinaryTree_InOrderTraversal(t *testing.T) {
    bt1 := BinaryTree{Value:1}
    output := bt1.InOrderTraversal()
    validateBinaryTreeTraversal(t, output, []interface{}{1})

    bt2 := BinaryTree{Value: 1}
    bt2.Right = &BinaryTree{Value: 2}
    output = bt2.InOrderTraversal()
    validateBinaryTreeTraversal(t, output, []interface{}{1, 2})

    bt3 := BinaryTree{Value: 1}
    bt3.Left = &BinaryTree{Value: 2}
    output = bt3.InOrderTraversal()
    validateBinaryTreeTraversal(t, output, []interface{}{2, 1})

    bt4 := BinaryTree{Value: 1}
    bt4.Left = &BinaryTree{Value: 2}
    bt4.Right = &BinaryTree{Value: 3}
    output = bt4.InOrderTraversal()
    validateBinaryTreeTraversal(t, output, []interface{}{2, 1, 3})

    bt5 := BinaryTree{Value: 1}
    bt5.Left = &BinaryTree{Value: 2}
    bt5.Right = &BinaryTree{Value: 3}
    bt5.Right.Left = &BinaryTree{Value: 4}
    bt5.Right.Right = &BinaryTree{Value: 5}
    output = bt5.InOrderTraversal()
    validateBinaryTreeTraversal(t, output, []interface{}{2, 1, 4, 3, 5})
}

func TestBinaryTree_PostOrderTraversal(t *testing.T) {
    bt1 := BinaryTree{Value:1}
    output := bt1.PostOrderTraversal()
    validateBinaryTreeTraversal(t, output, []interface{}{1})

    bt2 := BinaryTree{Value: 1}
    bt2.Right = &BinaryTree{Value: 2}
    output = bt2.PostOrderTraversal()
    validateBinaryTreeTraversal(t, output, []interface{}{2, 1})

    bt3 := BinaryTree{Value: 1}
    bt3.Left = &BinaryTree{Value: 2}
    output = bt3.PostOrderTraversal()
    validateBinaryTreeTraversal(t, output, []interface{}{2, 1})

    bt4 := BinaryTree{Value: 1}
    bt4.Left = &BinaryTree{Value: 2}
    bt4.Right = &BinaryTree{Value: 3}
    output = bt4.PostOrderTraversal()
    validateBinaryTreeTraversal(t, output, []interface{}{2, 3, 1})

    bt5 := BinaryTree{Value: 1}
    bt5.Left = &BinaryTree{Value: 2}
    bt5.Right = &BinaryTree{Value: 3}
    bt5.Right.Left = &BinaryTree{Value: 4}
    bt5.Right.Right = &BinaryTree{Value: 5}
    output = bt5.PostOrderTraversal()
    validateBinaryTreeTraversal(t, output, []interface{}{2, 4, 5, 3, 1})
}

func validateBinaryTreeTraversal(t *testing.T, output []*BinaryTree, expectedOutput []interface{}) {
    assert.Equal(t, len(output), len(expectedOutput))
    for i := 0; i < len(output); i++ {
        assert.Equal(t, expectedOutput[i], output[i].Value, "Mismatch at index %d", i)
    }
}
