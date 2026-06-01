package cards

// FavoriteCards returns a slice with the cards 2, 6 and 9 in that order.
func FavoriteCards() []int {
	var cards []int = []int{2,6,9}
    return cards
}

// GetItem retrieves an item from a slice at given position.
// If the index is out of range, we want it to return -1.
func GetItem(slice []int, index int) int {
	if index>=len(slice) || index<0{
        return -1
    }else{
        return slice[index]
    }
}

// SetItem writes an item to a slice at given position overwriting an existing value.
// If the index is out of range the value needs to be appended.
func SetItem(slice []int, index, value int) []int {
	if index >= len(slice) || index < 0{
        slice = append(slice, value)
    }else{
        slice[index] = value
    }
    return slice
}

// PrependItems adds an arbitrary number of values at the front of a slice.
func PrependItems(slice []int, values ...int) []int {
	var new_slice []int
    for i := range values{
        new_slice = append(new_slice, values[i])
    }
    for i := range slice{
        new_slice = append(new_slice, slice[i])
    }
    return new_slice
}

// RemoveItem removes an item from a slice by modifying the existing slice.
func RemoveItem(slice []int, index int) []int {
	if index >= len(slice) || index < 0{
        return slice
    }else{
        if index == 0{
            slice = slice[1:]
        }else if index == len(slice)-1{
            slice = slice[0:len(slice)-1]
        }else{
            slice = append(slice[0:index],slice[index+1:]...)
        }
    }
    return slice
}
