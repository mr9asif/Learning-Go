package main

import "fmt"

const a=10
var p=100

func call(){
	add:=func(x int, y int){
		z:=x+y
		fmt.Println(z)
	}
	add(2,3)
	add(p,a)
}
func main(){
	call()
	fmt.Println(a)
}

func init(){
	fmt.Println("hello i am first")
}


/*
 phase: 1. compile phase
        2. execution phase

		1. compile---> 01101001010101010010
		2. execution---> line by line

		when we write on terminal-->
		go run main.go --> it create .main compile file and -> auto run ./main
		if we want--we can do it-->
		go run build main.go--> build .main compile file--> then we have to manually run it
		./main in terminal.
*/

/*
  ---Compile phase------
  **Create a Code Segment***it store in binary compile file first
      a = 10
	  call = func(){..}
	  add = func(){..}
	  main = func(){..}
	  init = func(){..}
	  
	  ------Execution phase---- now code segment, global , stack and heap all create on ram
	  after do ./main then it will execute and semulation will start
	  **when a function run it alocate a memory in stack memory, it only runtime


	  -------How it work internally------
	   1.we know it execute line by lie
	   2. so first 'a' will store code segement and 'p' will store data segment.
	   3. then init() execute first , if not will go and find main() if main() not available ,get error
	   4. so after init() , in main(), line 17 we have call(),,it will find first on main , couldn't find then go code segment, find it then for execure call() create a space on stack memory.
	   4. then call has add(), it doesn't has in call but has on code segment, so it will use it as a reference. only call() can execute it.
	   5. ge line 13 and for add() also alocate stack memory space. and execute after it will remove again call add() line 14 then execute and remove. after that call() will be remove from stack memory. but they will remain on code segment. then execute main() and program finish all remove.


*/