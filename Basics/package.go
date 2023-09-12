package main

import (
	"fmt"
	// "math"
	"math/rand"
)

func main(){
	fmt.Println("The random integer is " , rand.Intn(100))
}


// rand.Intn returns the random integer from the range (0, n). It panics of n <= 0