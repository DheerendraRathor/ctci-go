package utils

func Max(a, b int) int {
    if a > b {
        return a
    } else {
        return b
    }
}

func Abs(a int) int {
    if a < 0 {
        return -a
    }

    return a
}
