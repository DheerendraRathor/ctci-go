package collections

import "testing"

func TestDataStructureEmptyError_Error(t *testing.T) {
    error := &DataStructureEmptyError{}
    if error.Error() != "Data structure you're trying to access is empty!" {
        t.Error("Error I got is not what I expected")
    }
}
