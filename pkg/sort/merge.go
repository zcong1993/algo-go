package sort

func merge(arr []int, l, m, r int) {
	n1 := m - l + 1
	n2 := r - m

	arr1, arr2 := make([]int, n1), make([]int, n2)
	copy(arr1, arr[l:m+1])
	copy(arr2, arr[m+1:r+1])

	i, j, k := 0, 0, l

	for i < n1 && j < n2 {
		if arr1[i] < arr2[j] {
			arr[k] = arr1[i]
			k++
			i++
		} else {
			arr[k] = arr2[j]
			k++
			j++
		}
	}

	for i < n1 {
		arr[k] = arr1[i]
		k++
		i++
	}

	for j < n2 {
		arr[k] = arr2[j]
		k++
		j++
	}
}

func mergeSort(arr []int, l, r int) {
	if l >= r {
		return
	}
	m := l + (r-l)/2
	mergeSort(arr, l, m)
	mergeSort(arr, m+1, r)
	merge(arr, l, m, r)
}

func MergeSort(arr []int) {
	mergeSort(arr, 0, len(arr)-1)
}
