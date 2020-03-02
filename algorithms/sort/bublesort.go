package sort

// BubleSort is the implementation of the Buble Sort algorithm
func BubleSort(collection []int) []int {
	arrayLength := len(collection)

	for maxLength := arrayLength; maxLength > 1; maxLength-- {
		hasSwapped := false
		for i := 0; i < maxLength-1; i++ {
			if collection[i] > collection[i+1] {
				collection[i], collection[i+1] = collection[i+1], collection[i]
				hasSwapped = true
			}
		}
		if !hasSwapped {
			break
		}
	}

	return collection
}
