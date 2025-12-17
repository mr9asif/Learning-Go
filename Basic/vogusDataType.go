package main

import "fmt"

func main(){
	/*
	var a int8 = -128 //-128 -- 127
	var b int8 = 127

	var x uint8 = 255 // 0-255

	var y float32 = 32
	var z float64 = 32  //if our cpu 32bit then mini go os will take 2row of ram .
	*/
	m:='😅' //rune
	var j float32=10.59584
     fmt.Printf("%.2f\n", j) //fmt.Println does not support "formatting verbs" like %.2f
	 fmt.Printf("%T", j)
	fmt.Printf("%c", m)  //we have to use "%c" to print this character imogi

}