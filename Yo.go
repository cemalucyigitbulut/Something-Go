package main

import (
	"fmt"
	"math"
	"math/rand"
)

func add(x int ,y int) int{
	return x+y
}

func add2(x,y int) int{
	return x+y
}

func swap(x,y string)(string,string){
	return x,y
}

var c,python,java bool

var h ,k int =1,2

func main(){
	fmt.Println("Hello , 世界")

	fmt.Println(math.Pi)

	randomNumber := rand.Intn(100)
	fmt.Println("Random number:", randomNumber)

	fmt.Println("My favorite number is", rand.Intn(100))

	fmt.Printf("Now you have %g problems.\n", math.Sqrt(7))

	fmt.Println(add(11,20))

	a,b := swap("hello","world")
	fmt.Println(a,b)

	var i int
	fmt.Println(i,c,python,java)

	var c,python,java = true ,false,"no!"
	fmt.Println(h,k,c,python,java)
}