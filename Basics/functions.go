package main

import "fmt"

func main(){
	fmt.Println("The addtion is - ", add(1, 12))
}

func add(x int, y int) int {
	return x + y 
}