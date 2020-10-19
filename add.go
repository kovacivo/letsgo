package main

import(		"fmt"
)

func add(x, y float32) float32 {
	return x + y
}

func main(){
	number1, number2 := float32(5.6), float32(9.1);
	fmt.Println("Adding two float numbers: 5.6 and 9.1 yields: ",add(number1,number2))
}
