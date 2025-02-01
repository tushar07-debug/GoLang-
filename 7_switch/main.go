package main

import "fmt"

func main() {

	//simple switch
	// i := 8

	// switch i {
	// case 1:
	// 	fmt.Println("one")
	// 	// break
	// case 2:
	// 	fmt.Println("two")
	// 	// break
	// case 3:
	// 	fmt.Println("Three")
	// 	// break
	// case 4:
	// 	fmt.Println("Four")
	// 	// break
	// case 5:
	// 	fmt.Println("five")
	// 	// break;
	// default:
	// 	fmt.Println("others")

	//break adn default are optional

	// multiple condition switch

	// switch time.Now().Weekday(){
	// case time.Saturday,time.Sunday:
	// 	fmt.Println("it's weekend day")
	// default:
	// 	fmt.Println("it's working day")
	// }

	//type switch
	whoAmI := func(i interface{}) {
		switch t := i.(type) {
		case int:
			fmt.Println("it's a integer")
		case string:
			fmt.Println("it's a string")
		case bool:
			fmt.Println("it's a boolean")
		default:
			fmt.Println("it's a other type", t)
		}
	}

	whoAmI("golang")
	whoAmI(50)
	whoAmI(true)
	whoAmI(4.5)
}
