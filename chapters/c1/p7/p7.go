package p7

import "fmt"

func RotateMatrix(matrix [][]int32) [][]int32 {
    if len(matrix) == 0 {
        return matrix
    }

    rows, columns := len(matrix), len(matrix[0])
    if rows != columns {
        panic(fmt.Sprintf("Not a matrix. Rows:%d Columns:%d\n", rows, columns))
    }

    n := rows

    for level := 0; level < n/2; level++ {
        row := level
        column := n - 1 - level
        for i := row; i < column; i++ {
            // Rotation logic
            temp := matrix[row][i]
            matrix[row][i] = matrix[i][column]
            matrix[i][column] = matrix[column][n-1-i]
            matrix[column][n-1-i] = matrix[n-1-i][row]
            matrix[n-1-i][row] = temp
        }
    }

    return matrix
}
