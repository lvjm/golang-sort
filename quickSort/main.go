package main

import "fmt"

func quickSort(values []int, l int, r int) {
	if l >= r {
		return
	}

	pivot := values[l]
	i := l + 1

	for j := l; j <= r; j++ {
		if pivot > values[j] {
			values[i], values[j] = values[j], values[i]
			i++
		}
	}

	values[l], values[i-1] = values[i-1], pivot

	quickSort(values, l, i-2)
	quickSort(values, i, r)
}

func main() {
	arr := []int{1, 8, 4, 3, 7, 9, 2, 7, 6, 1, 8}
	quickSort(arr, 0, len(arr)-1)
	fmt.Println(arr)
}
