package main

import "fmt"

func QuickSort(nums []int, low int, high int) []int {
	if low < high {
		pi := partition(nums, low, high)
		nums = QuickSort(nums, low, pi-1)
		nums = QuickSort(nums, pi+1, high)
	}
	return nums
}

func partition(nums []int, low int, high int) int {
	pivot := nums[high]
	i := low - 1
	for j := low; j <= high; j++ {
		if nums[j] < pivot {
			i++
			nums[i], nums[j] = nums[j], nums[i]
			fmt.Println(nums)
		}
	}
	nums[i+1], nums[high] = nums[high], nums[i+1]
	fmt.Println(nums)
	return i + 1
}

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 8, 9}
	nums = QuickSort(nums, 0, len(nums)-1)
	fmt.Print(nums)
}
