package slices

func GetItem(slice []int, index int)int{
	if index >= len(slice) || index <0{
		return -1
	}
	return slice[index]
}


