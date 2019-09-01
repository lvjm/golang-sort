package main

import "fmt"

func bubbleSort(arr []int) {
	flag := true // 如果一趟比较中没有交换则认为排序完成
	n := len(arr)
	for i := 0; i < n-1; i++ {
		flag = true
		for j := 0; j < n-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
				flag = false
			}
			continue
		}
		if flag {
			break
		}
	}
}

func main() {
	arr := []int{3, 5, 7, 2, 1, 9, 0}
	bubbleSort(arr)
	fmt.Println(arr)
}
