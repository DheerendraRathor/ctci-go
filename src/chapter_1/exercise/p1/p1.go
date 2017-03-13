package p1

import (
    "../../../utils"
)

func IsUnique(str string) bool {
    charMap := make(map[rune]bool)
    for _, char := range str {
        if charMap[char] {
            return false
        }
        charMap[char] = true
    }
    return true
}

func IsUniqueWithoutDS(str string) bool  {
    sortedStr := utils.SortString(str)
    isUnique := true
    for i := 1; i < len(sortedStr); i++ {
        if sortedStr[i] == sortedStr[i-1] {
            isUnique = false
            break
        }
    }
    return isUnique
}