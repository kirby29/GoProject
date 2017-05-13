package main

import ("fmt"
		"time"
)
	
func main(){
	//var name = "Mr.Kirby"
	//name := "Mr.Kirby"
	fmt.Println("Hello",validate("hello", 5))
	fmt.Println("time", time.Now())
}
func swap(x,y string)(string,string){
	return y,x
}
func  validate(input string, number int) int{
	if input == "hello"{
		return 4 * number
	}
	return 0 * number
}