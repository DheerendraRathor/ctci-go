package trees

import (
    "testing"
    "github.com/stretchr/testify/assert"
)

func TestTrie_Insert_NullTrie(t *testing.T) {
    val, ok := GetValueForString(nil, "Hello World")
    assert.Nil(t, val)
    assert.False(t, ok)

    trie := Insert(nil, "Hello World", true)
    assert.NotNil(t, trie)

    val, ok = GetValueForString(trie, "Hello World")
    assert.True(t, ok)
    assert.Equal(t, true, val)

    val, ok = GetValueForString(trie, "Hello")
    assert.False(t, ok)
    assert.Nil(t, val)
}

func TestTrie_Insert(t *testing.T) {
    trie := NewTrie()
    trie = Insert(trie, "Hello", 'H')
    trie = Insert(trie, "World", 'W')
    trie = Insert(trie, "From", 'F')
    trie = Insert(trie, "Golang", 'G')

    trie = Insert(trie, "Hell", 'h')

    var val interface{}
    var ok bool

    val, ok = GetValueForString(trie, "Hello World")
    assert.False(t, ok)
    assert.Nil(t, val)

    val, ok = GetValueForString(trie, "Hello")
    assert.True(t, ok)
    assert.Equal(t, 'H', val)

    val, ok = GetValueForString(trie, "World")
    assert.True(t, ok)
    assert.Equal(t, 'W', val)

    val, ok = GetValueForString(trie, "Hell!")
    assert.False(t, ok)
    assert.Nil(t, val)

    val, ok = GetValueForString(trie, "Hell")
    assert.True(t, ok)
    assert.Equal(t, 'h', val)

    val, ok = GetValueForString(trie, "Fro")
    assert.False(t, ok)
    assert.Nil(t, val)
}

func TestTrie_Remove(t *testing.T) {
    var trie *Trie
    var val interface{}
    var ok bool

    trie = Remove(nil, "Hello")
    assert.Nil(t, trie)

    trie = NewTrie()
    trie = Insert(trie, "Hello", 'H')
    trie = Insert(trie, "World", 'W')
    trie = Insert(trie, "From", 'F')
    trie = Insert(trie, "Golang", 'G')
    trie = Insert(trie, "Go", 'g')

    trie = Insert(trie, "Hell", 'h')

    trie = Remove(trie, "Hell")

    val, ok = GetValueForString(trie, "Hell")
    assert.False(t, ok)
    assert.Nil(t, val)

    val, ok = GetValueForString(trie, "Hello")
    assert.True(t, ok)
    assert.Equal(t, 'H', val)

    trie = Remove(trie, "Hello")
    trie = Remove(trie, "World")
    trie = Remove(trie, "From")
    trie = Remove(trie, "Golang")
    trie = Remove(trie, "Hello")

    assert.Equal(t, 1, len(trie.Children))

    val, ok = GetValueForString(trie, "Go")
    assert.True(t, ok)
    assert.Equal(t, 'g', val)
}
