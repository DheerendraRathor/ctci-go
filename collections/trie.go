package collections

// Trie for unicode chars
// Value stores the value corresponding to string for e.g. it can store phone number for a name
// Children map is a map of char to next node. Since we can have many unicode runes, keeping an array is not good
//      For leaf nodes, Children map will be empty
// isTerminal will be true if a word ends there
// parent is pointer to parent trie of current trie. For root node it will be nil
// currentChar stores rune which is mapped to this node in current word
type Trie struct {
    Value interface{}
    Children map[rune]*Trie
    isTerminal bool
    parent *Trie
    currentChar rune
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
            child.parent = currentTrie
            child.currentChar = char
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

// Silently remove string from trie if present. It won't throw error if it is not present
func Remove(trie *Trie, word string) *Trie {
    if trie == nil {
        return nil
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

    if isWordPresent {      // Cleanup node if required
        currentTrie.isTerminal = false
        currentTrie.Value = nil

        parent := currentTrie.parent
        for parent != nil && !currentTrie.isTerminal{
            if len(currentTrie.Children) == 0 { // No more string present. Better delete this node
                delete(currentTrie.parent.Children, currentTrie.currentChar)
                currentTrie = parent
                parent = parent.parent
            } else {
                break
            }
        }
    }

    return trie
}
