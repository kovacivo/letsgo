package main

import (
	"fmt"
	"time"
	"math/rand"
)

func main(){
	fmt.Println("hello world")
	fmt.Println("Today is: ", time.Now())
	fmt.Println("Some random number between 1 - 100 is: ", rand.Intn(100))
}
