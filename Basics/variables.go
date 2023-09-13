package main

import "fmt"

// package level var
var c, python, java bool


// k := 3 { we can't delacre this syntax outside the function. This is availiable only inside the function.}

func main(){

	// function level var
	var i int
	fmt.Println(i, c, python, java)

	k := "Omkar"
	fmt.Println(k)

}