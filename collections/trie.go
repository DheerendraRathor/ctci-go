package collections


// Trie for unicode chars
type Trie struct {
    Value interface{}
    Children map[rune]*Trie
    isTerminal bool
}

func NewTrie() *Trie {
    return &Trie{
        Children: make(map[rune]*Trie),
        isTerminal: false,
    }
}

func Insert(trie *Trie, word string, value interface{}) *Trie {
    if trie == nil {
        trie = NewTrie()
    }

    currentTrie := trie
    for _, char := range word {
        child := currentTrie.Children[char]
        if child == nil {
            child = NewTrie()
            currentTrie.Children[char] = child
        }

        currentTrie = child
    }

    currentTrie.isTerminal = true
    currentTrie.Value = value

    return trie
}

func GetValueForString(trie *Trie, word string) (interface{}, bool) {
    if trie == nil {
        return nil, false
    }

    currentTrie := trie
    isWordPresent := true
    for _, char := range word {
        child := currentTrie.Children[char]
        if child == nil {
            isWordPresent = false
            break
        }

        currentTrie = child
    }

    isWordPresent = isWordPresent && currentTrie.isTerminal
    var val interface{} = nil
    if isWordPresent {
        val = currentTrie.Value
    }

    return val, isWordPresent
}
