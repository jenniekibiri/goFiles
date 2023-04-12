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




}