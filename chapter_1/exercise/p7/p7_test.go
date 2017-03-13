package p7

import "testing"

type testStruct struct {
    input [][]int32
    output [][]int32
}

var testData []testStruct = []testStruct {
    {
        input: [][]int32{},
        output: [][]int32{},
    },
    {
        input: [][]int32{{0}},
        output: [][]int32{{0}},
    },
    {
        input: [][]int32{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}},
        output: [][]int32{{3, 6, 9}, {2, 5, 8}, {1, 4, 7}},
    },
    {
        input: [][]int32{{1, 2}, {3, 4}},
        output: [][]int32{{2, 4}, {1, 3}},
    },
}

func TestRotateMatrix(t *testing.T) {
    for _, data := range testData {
        originalInputSlice := copySlice(data.input)
        output := RotateMatrix(originalInputSlice)
        if !areMatrixEqual(output, data.output) {
            t.Errorf("Input: %v output:%v Expected output:%v\n", data.input, output, data.output)
        }
    }
}

func copySlice(m [][]int32) [][]int32 {
    newSlice := make([][]int32, len(m))
    for i := 0; i < len(m); i++ {
        newSlice[i] = make([]int32, len(m[i]))
    }

    for i := 0; i < len(m); i++ {
        for j := 0; j < len(m[0]); j++ {
            newSlice[i][j] = m[i][j]
        }
    }

    return newSlice
}

func areMatrixEqual(m1, m2 [][]int32)  bool {
    if len(m1) != len(m2) {
        return false
    }

    if len(m1) == 0 {
        return true
    }

    if len(m1[0]) != len(m2[0]) {
        return false
    }

    for i := 0; i< len(m1); i++ {
        for j := 0; j < len(m1[0]); j++ {
            if m1[i][j] != m2[i][j] {
                return false
            }
        }
    }

    return true
}

func TestRotateMatrixWithPanic(t *testing.T) {
    testMatrix := [][]int32 {
        {0, 0},
    }

    defer func() {
        if r := recover(); r == nil {
            t.Error("Didn't caught panic. Sad :(")
        }
    }()

    RotateMatrix(testMatrix)
}
