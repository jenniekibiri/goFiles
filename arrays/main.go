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
fmt.Printf("users slice %v",usersSlice)
}