package lib

import (
	"fmt"
	"testing"
)

func TestQuickSort(t *testing.T) {
	fmt.Println(QuickSort([]int{3, 7, 4, 5, 1, 2, 9}))
	fmt.Println(QuickSort([]int{1, 2, 3, 4, 5, 6, 7, 8, 9}))
	fmt.Println(QuickSort([]int{9, 8, 7, 6, 5, 4, 3, 2, 1}))
	fmt.Println(QuickSort([]int{9, 2, 7, 4, 8, 6, 3, 5, 1}))
}
