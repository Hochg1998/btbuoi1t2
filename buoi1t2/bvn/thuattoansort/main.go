package main

import "fmt"

func bubbleSort(arr []int) []int {
	for i := 0; i < len(arr); i++ {
		isSorted := true
		for j := 1; j < len(arr); j++ {
			if arr[j] < arr[j-1] {
				arr[j], arr[j-1] = arr[j-1], arr[j]
				isSorted = false
			}
		}
		if isSorted {
			return arr
		}
	}
	return arr
}

func mergeSort(arr []int) []int {
	lenghtArr := len(arr)
	if lenghtArr == 1 {
		return arr
	}
	//chia đôi arr
	left := make([]int, 0, int(lenghtArr/2))
	right := make([]int, 0, lenghtArr-int(lenghtArr/2))

	left = arr[:int(lenghtArr/2)]
	right = arr[int(lenghtArr/2):]

	return merge(mergeSort(left), mergeSort(right))

}

func merge(left, right []int) []int {
	result := make([]int, 0, len(left)+len(right))

	for len(left) > 0 && len(right) > 0 {
		if left[0] > right[0] {
			result = append(result, right[0])
			right = right[1:]
		} else {
			result = append(result, left[0])
			left = left[1:]
		}
	}
	if len(left) > len(right) {
		for i := 0; i < len(left); i++ {
			result = append(result, left[i])
		}
	} else {
		for i := 0; i < len(right); i++ {
			result = append(result, right[i])
		}
	}

	return result
}

func quickSort(arr []int, low int, high int) []int {

	if low >= high {
		return arr
	}
	p := partition(arr, low, high)

	quickSort(arr, low, p-1)
	quickSort(arr, p+1, high)

	return arr
}

func partition(arr []int, low int, high int) int {
	pivot := arr[high]
	i := low - 1
	for j := low; j < high; j++ {
		if arr[j] < pivot {
			i++
			arr[i], arr[j] = arr[j], arr[i]
		}
	}
	arr[i+1], arr[high] = arr[high], arr[i+1]
	return i + 1
}

func main() {
	fmt.Println(bubbleSort([]int{12, 33, 65, 27, 11, 13}))
	fmt.Println(mergeSort([]int{1, 3, 5, 2, 4, 8, 9}))
	fmt.Println(quickSort([]int{1, 3, 5, 2, 4, 8, 9}, 0, 6))

}
