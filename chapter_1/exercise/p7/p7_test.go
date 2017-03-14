package p7

import (
    "testing"
    "github.com/DheerendraRathor/ctci-go/utils"
)

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
        originalInputSlice := utils.Copy2DSlice(data.input)
        output := RotateMatrix(originalInputSlice)
        if !utils.AreInt32MatrixEqual(output, data.output) {
            t.Errorf("Input: %v output:%v Expected output:%v\n", data.input, output, data.output)
        }
    }
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
