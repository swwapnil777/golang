package main

import "fmt"

func main() {

	//var nums [4]int

	// nums[0] = 1

	// nums[1] = 2

	// nums[2] = 3

	// nums[3] = 4

	// for i := 0; i < len(nums); i++ {
	// 	fmt.Println(nums[i])
	// }

	nums1 := [4]int{1, 2, 3, 4}

	for i := 0; i < len(nums1); i++ {

		fmt.Println(nums1[i])
	}

}
