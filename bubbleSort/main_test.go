package main

import "testing"

func TestBubbleSort(t *testing.T) {
	a := []int{1, 3, 7, 4, 9, 0}
	b := []int{0, 1, 3, 4, 7, 9}
	bubbleSort(a)
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			t.Errorf("sort error!")
		}
	}
}
