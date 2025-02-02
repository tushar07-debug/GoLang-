package main

import "fmt"

func main() {
	// nums := []int{6, 7, 8}
	// for i := 0; i < len(nums); i++ {
	// 	fmt.Println(nums[i])
	// }


	// sum:=0
	// for _,num :=range nums{
	// 	// fmt.Println(num)
	// 	sum=sum+num
	// }
	// fmt.Println(sum)

	
	// for i,num :=range nums{
	// 	fmt.Println(num,i)
	// 	// sum=sum+num
	// }
	// fmt.Println(sum)

	// m:=map[string]string{"fname":"tushar","lname":"singh"}
	// for k,v := range m{
	// 	fmt.Println(k,v)
	// }

	for i,c := range "tushar"{
		fmt.Println(i,string(c))
	}
}