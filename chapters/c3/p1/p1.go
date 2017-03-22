package p1

import (
    "github.com/DheerendraRathor/ctci-go/collections"
)

/*
Algorithm:
- Start first stack from start of array and grow towards end
- Start second stack from end of array and grow towards start
- Start third stack from middle of array and grow towards end

- If third stack and second stack overlaps then shift third stack towards start to provide equal spacing from both first
and second stack
- If third stack and first stack overlaps then shift third stack towards end to provide euqal spacing from both first
and second stack
*/

type ThreeStacksArray struct {
    array []int
    size int

    firstEndIndex int

    thirdStartIndex int
    thirdEndIndex int

    secondEndIndex int
}

func InitializeTSA(size int) *ThreeStacksArray {
    return &ThreeStacksArray{
        array: make([]int, size),
        size: size,
        firstEndIndex: -1,
        secondEndIndex: size,
        thirdStartIndex: size/2,
        thirdEndIndex: size/2,
    }
}

func (tsa *ThreeStacksArray) shiftThirdStackEqually() {
    distanceInFirstAndSecnod := tsa.secondEndIndex - tsa.firstEndIndex - 1
    elementsInThird := tsa.thirdEndIndex - tsa.thirdStartIndex

    remainingSpace := distanceInFirstAndSecnod - elementsInThird

    if remainingSpace == 0 {
        panic("No more space to move")
    }

    spaceInLeft := remainingSpace / 2
    newThirdStartIndex := tsa.firstEndIndex + spaceInLeft + 1
    newThirdEndIndex := newThirdStartIndex + elementsInThird

    copy(tsa.array[newThirdStartIndex:newThirdEndIndex], tsa.array[tsa.thirdStartIndex : tsa.thirdEndIndex])
    tsa.thirdStartIndex = newThirdStartIndex
    tsa.thirdEndIndex = newThirdEndIndex
}

func (tsa *ThreeStacksArray) PushToFirst(item int) {
    if tsa.thirdStartIndex == tsa.firstEndIndex + 1 {   // Next correct index for first stack is taken by third stack
        tsa.shiftThirdStackEqually()
    }

    tsa.firstEndIndex++
    tsa.array[tsa.firstEndIndex] = item
}

func (tsa *ThreeStacksArray) PushToSecond(item int) {
    if tsa.thirdEndIndex == tsa.secondEndIndex - 1 {
        tsa.shiftThirdStackEqually()
    }

    tsa.secondEndIndex--
    tsa.array[tsa.secondEndIndex] = item
}

func (tsa *ThreeStacksArray) PushToThird(item int) {
    if tsa.thirdEndIndex == tsa.secondEndIndex {
        tsa.shiftThirdStackEqually()
    }

    tsa.array[tsa.thirdEndIndex] = item
    tsa.thirdEndIndex++
}

//
func (tsa *ThreeStacksArray) PopAt(stack uint8) (int, *collections.DataStructureEmptyError) {
    item := 0
    except := &collections.DataStructureEmptyError{}
    switch stack {
    case 1:
        if tsa.firstEndIndex > -1 {
            item = tsa.array[tsa.firstEndIndex]
            except = nil
            tsa.firstEndIndex--
        }
    case 2:
        if tsa.secondEndIndex < tsa.size {
            item = tsa.array[tsa.secondEndIndex]
            except = nil
            tsa.secondEndIndex++
        }
    case 3:
        if tsa.thirdEndIndex > tsa.thirdStartIndex {
            item = tsa.array[tsa.thirdEndIndex-1]
            except = nil
            tsa.thirdEndIndex--
        }
    default:

    }

    return item, except
}
