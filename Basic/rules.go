// Package rules-------------------------------
// 1. One folder = one package
// 2. All .go files in a folder MUST have same package name
// 3. package name is written manually
// 4. Folder name SHOULD match package name
// 5. func main() runs ONLY in package main
// 6. Capitalized names = exported (public)
// 7. lowercase names = unexported (private)
// 8. Library packages do NOT need main()

// ---------IMPORTANT-------
// 1.If i create other func then i will give the folder name as a package name;
// 2. But if there also contain the main func. then must be package name will be main;
// 3. Go can't run any function without main func. means it will run throw main func.
// --------->
// package main

// import "fmt"

// func Myfun(){
// 	fmt.Print("hlw")
// }

// func main(){
// 	Myfun()
// }

// -------------->
package main

import "fmt"

func main(){
	var name string
	fmt.Println("Enter your name:--")
    fmt.Scanln(&name);
	fmt.Println("your name is: ",name)
}

