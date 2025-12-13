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
package main

import "fmt"

func main(){
	fmt.Println("hellow init function")
}

func init(){
	fmt.Println("I execute before main , so don't say hlw")
}