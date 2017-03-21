package collections

type DataStructureEmptyError struct {}

func (s DataStructureEmptyError) Error() string {
    return "Data structure you're trying to access is empty!"
}
