package sort

// SelectionSort is the implementation of the Selection Sort algorithm
func SelectionSort(collection []int) []int {
	arrayLength := len(collection)

	for i := 0; i < arrayLength-1; i++ {
		minIndex := i
		for j := i + 1; j < arrayLength; j++ {
			if collection[j] < collection[minIndex] {
				minIndex = j
			}
		}

		collection[i], collection[minIndex] = collection[minIndex], collection[i]
	}

	return collection
}
