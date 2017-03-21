package p6

import (
    "strconv"
    "fmt"
)

func Compressor(input string) string {
    length := len(input)
    println(length, input)
    if length <= 1 {
        return input
    }

    // Since we're dealing with a..z only, there is no need to handle unicode
    output := make([]byte, length)

    var err error

    sliceUtilized := 0
    currentCount := 1
    currentChar := input[0]
    for i := 1; i < length; i++ {
        if input[i] == currentChar {
            currentCount += 1
        } else {

            output, sliceUtilized, err = compressorHelper(output, currentChar, currentCount, sliceUtilized, length)

            if err != nil {
                return input
            }

            currentChar = input[i]
            currentCount = 1
        }
    }

    output, sliceUtilized, err = compressorHelper(output, currentChar, currentCount, sliceUtilized, length)
    if err != nil {
        return input
    }

    return string(output[:sliceUtilized])
}

func compressorHelper(output []byte, char byte, count int, utilization int, inputLength int) ([]byte, int, error)  {
    intToBytes := []byte(strconv.Itoa(count))

    expectedUtilization := utilization + len(intToBytes) + 1   // Number of letters in currentCount + 1 for char appended
    if expectedUtilization >= inputLength {
        return output, utilization, fmt.Errorf("Crossed inputLength limit")
    }

    output[utilization] = char
    utilization += 1
    copy(output[utilization:], intToBytes[:])
    utilization += len(intToBytes)
    return output, utilization, nil
}