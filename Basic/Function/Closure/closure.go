package main

import "fmt"

const a = 10
var p = 100

func outer() func(){ //this return show function
   money:=100
   age:=30

   fmt.Println("Age: ", age)

   show :=func(){ //here closure work
	money = money+p+age
	fmt.Println(money)
   }
   return show
}

func call(){
	incr1 := outer()
	incr1()
	incr1()

	incr2 := outer()
	incr2()
	incr2()
}
func main(){
	call()

	fmt.Println("finish, good bye")
}

func init(){
	fmt.Println("i am first work")
}


/*
  -----How closure work here-------
  1. it work like the go work. i manually build it first then run it using ./closure.exe
  
  2. here line 21 when call outer(), it return a function, show(), here after return we know this outer() will remove from stack memory. but here, outer() return show(). so when we will try to execute show, we need money and show(). but outer is remove already so how can we get that. here comes closure. go compiler know that , that's why go store money and show() in heap memory for further work. here work escape analytics in compiler. so whenever we call now incr1(), it get the reference of outer or get show() and money in heap memory.
  3. in line 25 it works same for outer(). here also save money and show() in heap memory not previous one. that's how heap increase dynamically. 
*/