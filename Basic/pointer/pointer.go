package main

import "fmt"

func main(){
	// pointer or address of memory(ram)
	x:=20 // x=20
    fmt.Println(x)
	// p := address of x but in programming for address of we use & (emparsand)
	//  *p  => * means value of address
	p := &x
	
	*p = 30  // x = 30

	fmt.Println(x)
}