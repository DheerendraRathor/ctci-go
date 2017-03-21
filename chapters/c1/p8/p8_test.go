package p8

import (
    "testing"
    "github.com/DheerendraRathor/ctci-go/utils"
)

type testStruct struct {
    input [][]int32
    output [][]int32
}

var testData []testStruct= []testStruct {
    {
        input: [][]int32{},
        output: [][]int32{},
    },
    {
        input: [][]int32{{0}},
        output: [][]int32{{0}},
    },
    {
        input: [][]int32{{1}},
        output: [][]int32{{1}},
    },
    {
        input: [][]int32{{1, 2, 3}, {1, 4, 0}},
        output: [][]int32{{1, 2, 0}, {0, 0, 0}},
    },
    {
        input: [][]int32{{1, 0, 3}, {1, 4, 3}},
        output: [][]int32{{0, 0, 0}, {1, 0, 3}},
    },
}

func TestZeroMatrix(t *testing.T) {
    for _, data := range testData {
        originalInput := utils.Copy2DSlice(data.input)
        actualResult := ZeroMatrix(data.input)
        if !utils.AreInt32MatrixEqual(actualResult, data.output) {
            t.Errorf("Input: %v, output:%v, Expected output:%v", originalInput, actualResult, data.output)
        }
    }
}