package main

import "fmt"

func welcome(){
	fmt.Println("Welcome to our application!🙂")
}

// get name and return name
func name(){
	var name string
	fmt.Println("What's your name: ")
	fmt.Scanln(&name)
	fmt.Println("Ok", name , "let's move next 🙂")
}

func add(){
	var num1, num2 int
	fmt.Println("Give number 1: ")
	fmt.Scanln(&num1)
	fmt.Println("Give number 2: ")
	fmt.Scanln(&num2)
	fmt.Println("your num1:", num1, "num2: ", num2)
	sum:=num1+num2
	fmt.Println("Sum is: ", sum)
}

func bye(){
	fmt.Println("Thank you. Bye 🙂")
}

func main(){
     welcome()
	 name()
	 add()
 bye()

}