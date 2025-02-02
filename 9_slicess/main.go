package main

import (
	"fmt"
	"slices"
)

//slice are dyanmic array
//most used construct in go
// useful methods we can add
func main() {
	//unitialized slice is nil
	var nums []int
	fmt.Println(nums)
	fmt.Println(nums == nil)
	fmt.Println(len(nums))

	//slice
	var num = make([]int, 2, 5) //capacity, initial capacity
	// var num = make([]int,0,5) we put it 0 awlays
	fmt.Println(cap(num)) //capacity ->max num can fit
	fmt.Println(num)
	fmt.Println(num == nil)

	num = append(num, 1)
	num = append(num, 2)
	num = append(num, 3)
	num = append(num, 4)
	num = append(num, 5)
	fmt.Println(num)
	fmt.Println(cap(num))

	//  num:=[]int{} another way to make slice

	//copy slice
	var numo = make([]int, 0, 5)
	numo= append(numo,2)
	var nums2 = make([]int, len(numo))
	numo = append(numo, 2)
	copy(nums2,numo)
	fmt.Println(numo, nums2)
	

	//slice operator
	var nu = []int {1,2,3}
	fmt.Println(nu[0:2]) //or 
	fmt.Println(nu[:1])

	//slice equal function
	var num3 = []int{1,2}
	var num4 = []int{1,2}
	fmt.Println(slices.Equal(num3,num4))

	var nums5 = [][]int{{1,2,3},{4,5,6}}
	fmt.Println(nums5)
}
