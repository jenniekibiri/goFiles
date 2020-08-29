
package main

import "fmt"
//fmt is?

//initialize the func
func main(){
	//define variables
	//shorthand
	a ,b :=5,10
	swap(&a,&b)
	//pointers ? var whose value is the address of another var
fmt.Println(a,b)
}
func swap(a *int,b *int)  {
	*a,*b=*b,*a

}

