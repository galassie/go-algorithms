package sort

// QuickSort is the implementation of the Quick Sort algorithm
func QuickSort(collection []int) []int {
	arrayLength := len(collection)

	if arrayLength <= 1 {
		return collection
	}

	pivot := collection[arrayLength-1]
	var lesser []int
	var greater []int

	for i := 0; i < arrayLength-1; i++ {
		if collection[i] < pivot {
			lesser = append(lesser, collection[i])
		} else {
			greater = append(greater, collection[i])
		}
	}

	return append(append(QuickSort(lesser), pivot), QuickSort(greater)...)
}
