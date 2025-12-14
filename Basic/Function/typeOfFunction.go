// Function Types
//  1.standard or name function
//  2. Anonymous function
//  3.Funciton expression or assign function on variable
//  4. Higher order function or firstclass function
//  5. Callback function
//  6. Variatic function
//  7. Init function --> you can not call this, computer call this auto
//  8. Closure --closue over
//  9. Defer function ---> last in first out
//  10. Receiver function
//  11. IIEE -- Immediately involked function expression

// --------------------------------
// 1. standard or name
// this functin has name that's standard or name
//  func add(){
// function body
//  }

// ---------------------------
// Init function ----computer call it
// here init will execute first then main will
/*
package main

import "fmt"

func main(){
	fmt.Println("hellow init function")
}

func init(){
	fmt.Println("I execute before main , so don't say hlw")
}
*/

// -----------------------
// Anynomous function && IIEF funciton
/*
package main
import "fmt"
func main(){
	// anynomous function and in go it has to be call
      func(a int, b int){
		 c:=a+b
		 fmt.Println(c)
	  }(2,3) //IIEF --immediate call
}
*/

//	------------------
//
// function expression or Assign function
/*
package main

import "fmt"
func main(){
	// here if i store it in a variable then error will gone. and it's a expression then.
	// here also we can not call/invoke a function expressoin before defie
	add(3,4)  //will not work, has to be under defien in locally but globally will work
     add:= func(a int, b int){
		 c:=a+b
		 fmt.Println(c)
	  }
	  add(3,4)
}
*/

//   *** Descreate Math => Functioal programming => first order function and higher order function come.
// -----logic---------
//   1. Object (Animal, people , car etc)  --expample---Tutul is a student
//   2. Property (color, student) ----Apple is red
//   3. Relation ----Tutul is taller then me

// First order logic--->
//   1. All customer must be pay their pizza bills.

//  Higher order logic--->
//     rules: Any rules that apply to customers must also apply to tutul
//  example:  All customer must pay tips our waiters.

// --------------------**-----------
// First order function is our all simple function that work simple things. higher order function work with complex things.
//   first order functions are:
// 1. standard function or name function
//  2. anonnymous function
//  3. IIEF
//  4. function expression

// ** Higer order function or first class function
// why first class? --> when we assign a value in var then it's first class citize. in go we can also assign a function in a var. like a:=10, here it's first class .
//
//	  rules: 1. parameter must be receive a function---parameter=>function
//	 2. function return
//	3. both
//
// Example
package main

import "fmt"

/*
// get parameter as funciton
func processOperation(a int, b int, op func(p int, q int)){
	op(a,b)
}

func add(a int, b int){
	c:=a+b
	fmt.Println(c)
}

func main(){
	processOperation(2,3,add)
}
*/
//  return function
	func call() func (x int, y int){
		return add
	}

	func add(a int, b int){
	    c:=a+b
		fmt.Println(c)
	}

	func main(){
		sum:=call()
		sum(2,3)
	}