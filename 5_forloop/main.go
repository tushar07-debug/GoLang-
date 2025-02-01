package main

import "fmt"

//only construct in go for looping
func main(){
	//while loop
	// i:=1
	// for i <=3{
	// 	fmt.Println(i)
	// 	i=i+1
	// }

	//infinite loop
	// for{
	// 	println("1")
	// }


	//classic for loop
	// for i := 0; i<3; i++{
	// 	//break
	// 	//continue
	// 	if i==2{
	// 		continue
	// 	}
	// 	fmt.Println(i)
	// }

	for i:= range 5{
		fmt.Println(i)
	}
}