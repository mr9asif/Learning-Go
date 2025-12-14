/*
 1. struct funtion
 2. receiver function uses here
*/

package main

import "fmt"

type User struct{  //here User is type like string and others
	Name string
	age int
}

func printUserDetails(usr User){
	fmt.Println(usr)
	
}

// reciver function
// it will only work for custom type
func (usr User) printDetails(){
	fmt.Println(usr)
}

// different example
func (usr User) call(a int){
	fmt.Println(usr.Name)
	fmt.Println(a)
}

func main(){
	user1 := User{ //instance or object
		Name:"Asif",
		age:23,
	}

	// user2:=User{
	// 	Name:"Babu",
	// 	age:22,
	// }
	
	// printUserDetails(user1)
	// printUserDetails(user2)

	// user1.printDetails()  //user reciver function
	// user2.printDetails()

	user1.call(2)
}



/*
   ----simulation----
** code segment**
  1. User = type User struct(..)
  2. printDetailt()
  3. main()
*/