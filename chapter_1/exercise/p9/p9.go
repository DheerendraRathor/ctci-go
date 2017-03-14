package p9

import "strings"

func IsRotation(input, test string) bool {
    doubleInput := input + input
    if strings.Contains(doubleInput, test) {
        return true
    }

    return false
}
