package p4

func IsPalindromePermutation(str string) bool {

    charMap := make(map[rune]int) // to count occurrence of chars excluding spaces
    for _, char := range str {
        if char != ' ' {
            charMap[char]++
        }
    }

    oddCount := 0
    isPerm := true
    for _, val := range charMap {
        if val % 2 == 1 {
            oddCount++
        }
        if oddCount > 1 {
            isPerm = false
            break
        }
    }

    return isPerm
}
