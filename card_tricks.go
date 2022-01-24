package cards

// GetItem retrieves an item from a slice at given position. The second return value indicates whether
// the given index exists in the slice or not.
func GetItem(slice []int, index int) (int, bool) {
	if index< len(slice) && index >=0 {
        return slice[index], true
    } else if index>=len(slice) || index<0{
        return 0,false
    } else {
    return 0,false
    }
}

// SetItem writes an item to a slice at given position overwriting an existing value.
// If the index is out of range the value needs to be appended.
func SetItem(slice []int, index, value int) []int {
	if index < len(slice) && index >= 0 {
        slice[index] = value
        return slice
    }
    slice = append(slice, value)
    return slice
}

// PrefilledSlice creates a slice of given length and prefills it with the given value.
func PrefilledSlice(value, length int) []int {
	if (length <= 0) {
        return nil
    }
	slice := make([]int, length)
    for idx := range slice {
        slice[idx] = value
    }
	return slice
}

// RemoveItem removes an item from a slice by modifying the existing slice.
func RemoveItem(slice []int, index int) []int {
    i:= index
    if index < len(slice) && index>=0 {
    slice1 := slice[:i]
    slice2 := slice[(i+1):]
    slice3 := append(slice1, slice2...)
    return slice3
    } else {
    return slice
    }
}
