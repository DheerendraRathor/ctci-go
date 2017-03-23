package p5

import "github.com/DheerendraRathor/ctci-go/collections"

func SortStack(stack *collections.Stack) {

    // We don't sort nil stacks!
    if stack == nil {
        return
    }

    // Will contain sorted stack in desc order
    tempStack := collections.Stack{}

    for !stack.IsEmpty() {
        val, _ := stack.Pop()
        if tempStack.IsEmpty() {
            tempStack.Push(val)
        } else {

            // While top value in temp stack is larger than val, move value from temp stack to main stack
            for true {
                tempVal, err := tempStack.Peek()

                // Top value in temporary stack is larger than value we want to push
                if err == nil && tempVal.(int) > val.(int) {
                    tempStack.Pop()
                    stack.Push(tempVal)
                } else {
                    break
                }
            }

            tempStack.Push(val)
        }
    }

    // Now temp stack is sorted in desc order, just reverse it to main stack
    for !tempStack.IsEmpty() {
        val, _ := tempStack.Pop()
        stack.Push(val)
    }
}
