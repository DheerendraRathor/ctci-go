package p1

import (
    "testing"
    "github.com/stretchr/testify/assert"
)

func TestInitializeTSA(t *testing.T) {
    tsa := InitializeTSA(5)

    if tsa.firstEndIndex != -1 {
        t.Error("First end index should be -1")
    }
    if tsa.secondEndIndex != 5 {
        t.Error("Second end index should be 5")
    }
    if tsa.thirdStartIndex != 2 {
        t.Error("Third start index should be 2")
    }
    if tsa.thirdEndIndex != 2 {
        t.Error("Third end index should be 2 as well")
    }
}

func TestThreeStacksArray_PushToFirst(t *testing.T) {
    tsa := InitializeTSA(10)
    tsa.PushToFirst(10)
    tsa.PushToFirst(20)

    if tsa.firstEndIndex != 1 {
        t.Error("First end index should be 1")
    }
    if tsa.array[0] != 10 && tsa.array[1] != 20 {
        t.Error("Invalid elements found in underlying array")
    }
}

func TestThreeStacksArray_PushToSecond(t *testing.T) {
    tsa := InitializeTSA(10)
    tsa.PushToSecond(10)
    tsa.PushToSecond(20)

    if tsa.secondEndIndex != 8 {
        t.Error("Second end index should be 8")
    }
    if tsa.array[9] != 10 && tsa.array[8] != 20 {
        t.Error("Invalid elements found in underlying array")
    }
}

func TestThreeStacksArray_PushToThird(t *testing.T) {
    tsa := InitializeTSA(10)
    tsa.PushToThird(10)
    tsa.PushToThird(20)

    if tsa.thirdEndIndex != 7 {
        t.Error("Third end index should be 7")
    }
    if tsa.thirdStartIndex != 5 {
        t.Error("Third start index should b1 5")
    }
    if tsa.array[5] != 10 && tsa.array[6] != 20 {
        t.Error("Invalid elements found in underlying array")
    }
}

func TestThreeStacksArray_PopAt(t *testing.T) {
    tsa := InitializeTSA(10)

    tsa.PushToFirst(10)
    tsa.PushToSecond(20)
    tsa.PushToThird(30)

    firstTop, err := tsa.PopAt(1)
    if err != nil {
        t.Error("Got err as null")
    }
    if firstTop != 10 {
        t.Error("Top at first is not 10")
    }

    _, err = tsa.PopAt(1)
    if err == nil {
        t.Error("I should have got error here. First stack")
    }

    //---------------------------------------------
    secondTop, err := tsa.PopAt(2)
    if err != nil {
        t.Error("Got err as null")
    }
    if secondTop != 20 {
        t.Error("Top at second is not 20")
    }

    _, err = tsa.PopAt(2)
    if err == nil {
        t.Error("I should have got error here. Second stack")
    }

    //-----------------------------------------
    thirdTop, err := tsa.PopAt(3)
    if err != nil {
        t.Error("Got err as null")
    }
    if thirdTop != 30 {
        t.Error("Top at third is not 30")
    }

    _, err = tsa.PopAt(3)
    if err == nil {
        t.Error("I should have got error here. Third stack")
    }

    _, err = tsa.PopAt(4)
    if err == nil {
        t.Error("Trying to access invalid stack but got no err")
    }
}

func TestThreeStacksArray_PushToFirst2(t *testing.T) {
    tsa := InitializeTSA(5)
    tsa.PushToThird(1)

    assert.Equal(t, tsa.array[2], 1)

    tsa.PushToFirst(2)
    tsa.PushToFirst(3)


    assert.Equal(t, tsa.array[0], 2)
    assert.Equal(t, tsa.array[1], 3)

    tsa.PushToFirst(4)

    assert.Equal(t, tsa.array[2], 4)
    assert.Equal(t, tsa.array[3], 1)
}

func TestThreeStacksArray_PushToSecond2(t *testing.T) {
    tsa := InitializeTSA(5)
    tsa.PushToThird(1)

    assert.Equal(t, tsa.array[2], 1)

    tsa.PushToSecond(2)
    tsa.PushToSecond(3)


    assert.Equal(t, tsa.array[4], 2)
    assert.Equal(t, tsa.array[3], 3)

    tsa.PushToSecond(4)

    assert.Equal(t, tsa.array[2], 4)
    assert.Equal(t, tsa.array[1], 1)

}

func TestThreeStacksArray_PushToThird2(t *testing.T) {
    tsa := InitializeTSA(5)
    tsa.PushToThird(1)
    tsa.PushToThird(2)
    tsa.PushToThird(3)
    tsa.PushToThird(4)

    assert.Equal(t, 1, tsa.array[1])
    assert.Equal(t, 2, tsa.array[2])
    assert.Equal(t, 3, tsa.array[3])
    assert.Equal(t, 4, tsa.array[4])

    tsa.PushToThird(5)

    assert.Equal(t, 1, tsa.array[0])
    assert.Equal(t, 2, tsa.array[1])
    assert.Equal(t, 3, tsa.array[2])
    assert.Equal(t, 4, tsa.array[3])
    assert.Equal(t, 5, tsa.array[4])
}

func TestThreeStacksArray_Panic(t *testing.T) {
    tsa := InitializeTSA(5)
    tsa.PushToFirst(1)
    tsa.PushToSecond(2)
    tsa.PushToFirst(3)
    tsa.PushToSecond(4)
    tsa.PushToThird(5)

    assert.Equal(t, 1, tsa.array[0])
    assert.Equal(t, 3, tsa.array[1])
    assert.Equal(t, 5, tsa.array[2])
    assert.Equal(t, 4, tsa.array[3])
    assert.Equal(t, 2, tsa.array[4])

    defer func() {
        if r := recover(); r == nil {
            t.Error("Waited for panic. Didn't get any. Said life")
        }
    }()

    tsa.PushToFirst(6)

}
