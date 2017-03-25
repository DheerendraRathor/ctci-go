package graphs

import (
    "testing"
    "github.com/stretchr/testify/assert"
)

func TestGraph_AddNode(t *testing.T) {
    g := New()
    g.AddNode("a")
    g.AddNode("b")
    g.AddNode("c")

    var doesNodeExists bool
    doesNodeExists = g.DoesNodeExists("a")
    assert.True(t, doesNodeExists)

    doesNodeExists = g.DoesNodeExists("d")
    assert.False(t, doesNodeExists)

    doesNodeExists = g.DoesNodeExists("b")
    assert.True(t, doesNodeExists)

    // Assert edges
    assert.NotNil(t, g.edges["a"])
    assert.NotNil(t, g.edges["b"])
    assert.NotNil(t, g.edges["c"])
    assert.Nil(t, g.edges["d"])

    // Assert reverse edges
    assert.NotNil(t, g.reverseEdges["a"])
    assert.NotNil(t, g.reverseEdges["b"])
    assert.NotNil(t, g.reverseEdges["c"])
    assert.Nil(t, g.reverseEdges["d"])
}

func TestGraph_AddEdge(t *testing.T) {
    g := New()
    g.AddNode("a")

    edges, ok := g.GetEdges("a")
    assert.True(t, ok)
    assert.Equal(t, 0, len(edges))

    g.AddEdge("a", "b", 1)
    edges, ok = g.GetEdges("a")
    assert.True(t, ok)
    assert.Equal(t, 1, len(edges))
    assert.Equal(t, 1, len(g.reverseEdges["b"]))

    edges, ok = g.GetEdges("b")
    assert.True(t, ok)
    assert.Equal(t, 0, len(edges))

    edges, ok = g.GetEdges("c")
    assert.False(t, ok)

    g.AddEdge("c", "a", 2)
    edges, ok = g.GetEdges("a")
    assert.True(t, ok)
    assert.Equal(t, 1, len(edges))
    assert.Equal(t, 1, len(g.reverseEdges["a"]))
}

func TestGraph_RemoveNode(t *testing.T) {
    g := New()

    g.RemoveNode("a")       // Trying to delete non-existent node

    g.AddNode("a")
    g.AddNode("b")
    g.AddNode("c")
    g.AddNode("d")
    g.AddNode("e")

    g.AddEdge("a", "b", 1)
    g.AddEdge("b", "c", 1)
    g.AddEdge("c", "d", 1)
    g.AddEdge("b", "d", 1)
    g.AddEdge("e", "b", 1)
    g.AddEdge("c", "b", 1)

    edges, ok := g.GetEdges("b")
    assert.True(t, ok)
    assert.Equal(t, 2, len(edges))

    assert.Equal(t, 3, len(g.reverseEdges["b"]))

    g.RemoveNode("b")

    doesNodeExists := g.DoesNodeExists("b")
    assert.False(t, doesNodeExists)

    edges, ok = g.GetEdges("b")
    assert.False(t, ok)

    assert.Equal(t, 0, len(g.reverseEdges["b"]))
    assert.Equal(t, 0, g.edges["a"]["b"])
    assert.Equal(t, 0, g.edges["c"]["b"])
    assert.Equal(t, 0, g.edges["e"]["b"])
}

func TestGraph_RemoveEdge(t *testing.T) {
    g := New()

    g.AddNode("a")
    g.AddNode("b")
    g.AddNode("c")
    g.AddNode("d")
    g.AddNode("e")

    g.AddEdge("a", "b", 1)
    g.AddEdge("b", "c", 1)
    g.AddEdge("c", "d", 1)
    g.AddEdge("b", "d", 1)
    g.AddEdge("e", "b", 1)
    g.AddEdge("c", "b", 1)

    edges, ok := g.GetEdges("b")
    assert.True(t, ok)
    assert.Equal(t, 2, len(edges))
    assert.Equal(t, 1, len(g.reverseEdges["c"]))

    g.RemoveEdge("b", "c")

    edges, ok = g.GetEdges("b")
    assert.True(t, ok)
    assert.Equal(t, 1, len(edges))
    assert.Equal(t, 0, len(g.reverseEdges["c"]))
}
