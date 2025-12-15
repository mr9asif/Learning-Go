/*
  ---slice-----
    slice has 2 things
	 1. pointer --- where it is start from
	 2. length ---- how much it take like from to  slice end , 1 : 4 -- len(3)
	 3. capasity --- from start to arr end
*/

package main

import (
	"fmt"
)

// ----variatic function----
 func print(numbers ...int){
		fmt.Println(numbers)
	 }

func main(){

	

	//  print(12,3,4,4)
	/*
	arr :=[5] string {"i" , "am", "asif", "khan", "pro"}
	fmt.Println(arr)

	s :=arr[1:4]
	fmt.Println(s) //slice--- [am asif khan]

	s2 :=s[2:3]
	fmt.Println(s2)  // [khan]

	*/

	// arr := []int {1,2,3} //slice literal

	// arr:=make([]int, 3, 5)  // it'a a way to create slice, here lenth-3, capasity-5

	// var arr []int //emty or nill arr

    // s :=append(arr, 1,2,3)   //here append 1,2,3 in this arr
	// fmt.Print(s)
	// how it works? here create a nill or emty arr[] on ram. then we append 1,2,3 so after nill it will create 3 el of array. and after it will give a name itself. but here 1,2,3 will store heap memory cause it will remove after return but we need it later.
}



/*
  1. slice from arr
  2. slice from slice
  3. slice literal
  4. create func with len, and cap
  5. empty slice or nill ---
  6. slice underlying array --- 1024 till increase new array 100% mean 2x. then 25% increase
*/ 