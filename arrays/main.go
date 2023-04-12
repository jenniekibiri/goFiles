package main

import (
	"fmt"
)
func main (){

var users [20]string
users[0] = "ken"
users[1] = "charles"
users[8] = "james"
users[2] = "8"
fmt.Printf("users aryyy %v",users)
// fmt.Printf("users array length %v",len(users))
println(users[0])


//array are fixed length and cannot be changed
//slices are dynamic and can be changed
//slices are like arrays but they are dynamic
var usersSlice []string
usersSlice =append(usersSlice,"ken")
usersSlice =append(usersSlice,"charles")
fmt.Printf("users slice %v \n",usersSlice)
//loops 
//for loop
for { // this is like a while loop 
	var firstName string
	for _, user :=range usersSlice{
		fmt.Println( "user",user)
	}
	fmt.Println("enter your name")
	fmt.Scan(&firstName)
fmt.Println("hello",firstName)
}
//for loop

//maps like javacript objects 
//key value pairs
var usersMap map[string]string
usersMap = make(map[string]string)
usersMap["firstName"] = "ken"
usersMap["lastName"] = "charles"
fmt.Printf("users map %v \n",usersMap)
// or 

// functions  use the 
add(1,2) // usually outside the main function 
// you specify the return type after the function name

//conditionals 
//switch
switch "firstName" {
case "ken":
	fmt.Println("hello ken")
case "charles":
	fmt.Println("hello charles")
default:
	fmt.Println("hello default")
}

//structs
type User struct{
	FirstName string
	LastName string
	Age int
}
var userMap2 map[string]User
userMap2 = make(map[string]User)
userMap2["ken"] = User{
	FirstName: "ken",
	LastName: "charles",
	Age: 12,
}
fmt.Printf("users map %v \n",userMap2)
// usually a way to create a map with types  of different types
//concurrency
//you go  to separate threads
//async  w.g wait group 

//goroutines



}
func add(a int, b int) int{
	return a+b
}