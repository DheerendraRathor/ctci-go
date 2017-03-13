package p2

import (
    "utils"
)

func ArePermutations(str1, str2 string) bool  {
    if len(str1) != len(str2) {
        return false
    }
    sortedStr1, sortedStr2 := utils.SortString(str1), utils.SortString(str2)
    for index, _ := range sortedStr1 {
        if sortedStr1[index] != sortedStr2[index] {
            return false
        }
    }
    return true
}
