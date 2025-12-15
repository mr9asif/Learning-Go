package main

import (
	"fmt"
)

	 func num(numbers *[3]int){
       fmt.Println(numbers)
	 }

func main(){
	// pointer or address of memory(ram)
	x:=20 // x=20
    // fmt.Println(x)
	// p := address of x but in programming for address of we use & (emparsand)
	//  *p  => * means value of address
	p := &x
	
	*p = 30  // x = 30

	// fmt.Println(x)
 

	arr := [3]int {1,2,3}  
	s := arr[1:2]
       fmt.Println(s) 
	// fmt.Println(arr)  // in go we didn't need any loop
	num(&arr)  //now using empasant and star, it didn't coppy like 3 million data, just coppy 1 address of that array. also don't need to 3 one. 
}