package sort

import "math/rand"

func QuickSort(arr []int) {
	qSort(arr, 0, len(arr)-1)
}

func qSort(arr []int, l, r int) {
	if l >= r {
		return
	}
	p := partition(arr, l, r)
	qSort(arr, l, p-1)
	qSort(arr, p+1, r)
}

func partition(arr []int, l, r int) int {
	if l == r {
		return l
	}
	// 增加随机性
	pIdx := l + rand.Intn(r-l)
	pivot := arr[pIdx]
	arr[r], arr[pIdx] = arr[pIdx], arr[r]
	idx := l - 1
	for i := l; i < r; i++ {
		if arr[i] < pivot {
			idx++
			arr[idx], arr[i] = arr[i], arr[idx]
		}
	}

	idx++
	arr[idx], arr[r] = arr[r], arr[idx]
	return idx
}

func QuickSort2(arr []int) {
	sort2(arr, 0, len(arr)-1)
}

func sort2(arr []int, l, h int) {
	if l >= h {
		return
	}
	p := partition2(arr, l, h)
	sort2(arr, l, p-1)
	sort2(arr, p+1, h)
}

func partition2(arr []int, l, h int) int {
	if l == h {
		return l
	}

	pIdx := l + rand.Intn(h-l)
	p := arr[pIdx]
	arr[l], arr[pIdx] = arr[pIdx], arr[l]

	i, j := l+1, h

	for {
		for arr[i] < p {
			if i == h {
				break
			}
			i++
		}

		for arr[j] > p {
			if j == l {
				break
			}
			j--
		}

		if i >= j {
			break
		}

		arr[i], arr[j] = arr[j], arr[i]
	}
	arr[j], arr[l] = arr[l], arr[j]
	return j
}
