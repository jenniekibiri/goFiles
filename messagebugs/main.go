package main

import "fmt"
 
func main(){
	//variables
 count :=5

//use normal variable definition for undefined var
var message string
if count >5 {
	message = "greater then 5"
	println(message)
}else{
	 message = "Not greater than 5"
	println(message)
}
	fmt.Println("hello")
}