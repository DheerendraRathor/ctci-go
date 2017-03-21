package p8

func ZeroMatrix(matrix [][]int32) [][]int32 {
    if len(matrix) == 0 {
        return matrix
    }

    firstRowHasZero := false
    firstColumnHasZero := false

    m, n := len(matrix), len(matrix[0])
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            if matrix[i][j] == 0 {
                // Mark top row and left most column 0
                matrix[0][j] = 0
                matrix[i][0] = 0

                if i == 0 {
                    firstRowHasZero = true
                }
                if j == 0 {
                    firstColumnHasZero = true
                }
            }
        }
    }

    // If first value in row is 0 then make all values in that row 0
    for i := 1; i < m; i++ {
        if matrix[i][0] == 0 {
            for j := 0; j < n; j++ {
                matrix[i][j] = 0
            }
        }
    }

    // If first value of column is 0 then make all values in that column 0
    for i := 1; i < n; i++ {
        if matrix[0][i] == 0 {
            for j := 0; j < m; j++ {
                matrix[j][i] = 0
            }
        }
    }

    if firstRowHasZero {
        for i := 0; i < n; i++ {
            matrix[0][i] = 0
        }
    }

    if firstColumnHasZero {
        for i := 0; i < m; i++ {
            matrix[i][0] = 0
        }
    }

    return matrix
}
