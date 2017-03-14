package utils

func Copy2DSlice(m [][]int32) [][]int32 {
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