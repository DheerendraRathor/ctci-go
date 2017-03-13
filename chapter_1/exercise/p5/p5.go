package p5

func IsOneAway(input, output string) bool {
    isOneAway := true

    inputRunes, outputRunes := []rune(input), []rune(output)

    inputLength := len(inputRunes)
    outputLength := len(outputRunes)
    if inputLength == outputLength - 1 {
        for i := 0; i < inputLength; i++ {
            if inputRunes[i] != outputRunes[i] {
                isOneAway = areEqual(inputRunes[i:], outputRunes[i+1:])
                break
            }
        }
    } else if inputLength == outputLength {
        for i := 0; i < inputLength; i++ {
            if inputRunes[i] != outputRunes[i] {
                isOneAway = areEqual(inputRunes[i+1:], outputRunes[i+1:])
                break
            }
        }
    } else if inputLength == outputLength + 1 {
        inputRunes, outputRunes = outputRunes, inputRunes
        for i := 0; i < outputLength; i++ {
            if inputRunes[i] != outputRunes[i] {
                isOneAway = areEqual(inputRunes[i:], outputRunes[i+1:])
                break
            }
        }
    } else {
        isOneAway = false
    }

    return isOneAway
}

func areEqual(r1, r2 []rune) bool {
    if len(r1) != len(r2) {
        return false
    }

    for index, char := range r1 {
        if char != r2[index] {
            return false
        }
    }

    return true
}