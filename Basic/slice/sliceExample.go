package main

import "fmt"

func changeSlice(a []int) []int {
	a[0]=10
	a = append(a, 11)
	return a
}
func main(){
	x:=[]int {1,2,3,4,5}
	x = append(x, 6)
	x = append(x, 7)
	a:=x[4:]
	y:=changeSlice(a)

	fmt.Println(x)
	fmt.Println(y)
}

/*
 1. it start with main(). then x has declire and assign --> [1,2,3,4,5] -->place stack(9-13)(img)
 2. then we append 6 in x, so here x firstly call append so it will take space on stack memory. when x prev has-> ptr(9), len(5),cap(5) although our len and cap is equal we need more space to append.now it will coppy whole in heap memory. but now below 1024 it will 2x space, so cap will be len*2=>10, so now x->ptr(20),len(6), cap(10)
 3. again do same thing but now we don't need to create  or coppy this on heap , cause we have enought cap, just append it. so now x-->ptr(20), len(6), cap(10)
 4. now we slice x on a:=[4:] means 4 to cap. now ptr(24), len(3), cap(6)
 5. we call changeSlice() and change idex value so a = [10,6,7,11] it's y =a;
 so  x == [1 2 3 4 5 6 7], y == [10 6 7 11]
*/ 