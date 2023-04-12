package main
import "fmt"
//use double quotes
//create the main fun
func main(){
	//define variable 
	//normal method
	var firstName = "ken"
	var familyName="charles"
	//shorthand method
	age :=12
	peanutAllergy:=true
	//print the values
	//const for not changing values

//print the values format the output with a new line 
fmt.Println(firstName)
fmt.Println(familyName)
fmt.Println(age)
fmt.Println(peanutAllergy)
//print does not format the output with a new line 
//print the values format the output with a new line  
// the printf doesnt not add a new line you have to add it manual
fmt.Printf("my name is %v %v and i am %v years old\n",firstName,familyName,age)
//gettting user input
 var username string; 
 fmt.Println("enter your name")
 //you use the & to get the address of the variable 
 fmt.Scan(&username)
 fmt.Println("hello",username)
}


