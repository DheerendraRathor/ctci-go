package p5

import "github.com/DheerendraRathor/ctci-go/utils"

func IsOneAway(input, output string) bool {
    isOneAway := true

    inputRunes, outputRunes := []rune(input), []rune(output)

    inputLength := len(inputRunes)
    outputLength := len(outputRunes)
    if inputLength == outputLength - 1 {
        for i := 0; i < inputLength; i++ {
            if inputRunes[i] != outputRunes[i] {
                isOneAway = utils.AreRuneSlicesEqual(inputRunes[i:], outputRunes[i+1:])
                break
            }
        }
    } else if inputLength == outputLength {
        for i := 0; i < inputLength; i++ {
            if inputRunes[i] != outputRunes[i] {
                isOneAway = utils.AreRuneSlicesEqual(inputRunes[i+1:], outputRunes[i+1:])
                break
            }
        }
    } else if inputLength == outputLength + 1 {
        inputRunes, outputRunes = outputRunes, inputRunes
        for i := 0; i < outputLength; i++ {
            if inputRunes[i] != outputRunes[i] {
                isOneAway = utils.AreRuneSlicesEqual(inputRunes[i:], outputRunes[i+1:])
                break
            }
        }
    } else {
        isOneAway = false
    }

    return isOneAway
}
