package collections

type Queue struct {
    first *Node
    last *Node
    length int
}

func (q *Queue) IsEmpty() bool {
    return q.first == nil
}

func (q *Queue) Add(item interface{}) {
    newNode := &Node{Value: item}

    if q.first == nil {
        q.first = newNode
        q.last = newNode
    } else {
        q.last.Next = newNode
        q.last = newNode
    }

    q.length++
}

func (q *Queue) Remove() (interface{}, *DataStructureEmptyError) {
    if q.IsEmpty() {
        return nil, &DataStructureEmptyError{}
    }

    topNode := q.first

    // Only one element present in Queue
    if q.first == q.last {
        q.first = nil
        q.last = nil
    } else {
        q.first = q.first.Next
    }

    q.length--
    return topNode.Value, nil
}

func (q *Queue) Peek() (interface{}, *DataStructureEmptyError) {
    if q.IsEmpty() {
        return nil, &DataStructureEmptyError{}
    }

    return q.first.Value, nil
}

func (q *Queue) Length() int {
    return q.length
}
