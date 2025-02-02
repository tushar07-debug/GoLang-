package main

import "fmt"

func main() {
	//maps->hash,objects,dict
	//creating map
	m := make(map[string]string)
	//setting an element
	m["name"] = "golang"
	m["area"] = "backend"
	//get an elment
	fmt.Println(m["name"], m["area"])
	//imp: if key does not exist in the map then it return zero

	//delete
	delete(m, "name")
	fmt.Println(m)

	//clear
	clear(m)
	fmt.Println(m)

	// map without usign make
	// m := map[string]int{"name": 60, "area": 40}
	// fmt.Println(m)
}
