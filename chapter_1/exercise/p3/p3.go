package p3


func URLifyWithSliceCopy(str string, trueLength int) string {
    runes := []rune(str) // rune conversion is done to maintain utf-8 internationalization
    var index int32 = 0
    for i := 0; i < trueLength;i++ {

        if runes[index] != 0x20 {  // Space check
            index++
        } else {
            copy(runes[index+3:], runes[index+1:])
            copy(runes[index:index+3], []rune("%20"))
            index += 3
        }
    }
    return string(runes)
}

func URLifyTwoPass(str string, trueLength int) string {
    runes := []rune(str)
    spacesCount := 0
    for i := 0; i < trueLength; i++ {
        if runes[i] == 0x20 {
            spacesCount++
        }
    }

    lastIndex := trueLength + spacesCount * 2;
    for i := trueLength - 1; i >= 0; i-- {
        if runes[i] != 0x20 {
            runes[lastIndex-1] = runes[i]
            lastIndex--
        } else {
            copy(runes[lastIndex-3:lastIndex], []rune("%20"))
            lastIndex -= 3
        }
    }
    return string(runes)
}
