// Basic code
/* package main

import "fmt"
func main(){
	 var a int =20
	  a=40
	fmt.Println(a)
} */

// variable in go : how it works

package main

import "fmt"

func main() {

	// ===============================
	// 1. VARIABLE DECLARATION
	// ===============================

	// Using var (explicit type)
	var a int
	a = 10
	fmt.Println("a =", a)

	// Using var with value
	var b int = 20
	fmt.Println("b =", b)

	// Using var with type inference
	var c = 30 // Go infers int
	fmt.Println("c =", c)

	// Short variable declaration (MOST COMMON)
	d := 40
	fmt.Println("d =", d)

	// IMPORTANT:
	// := can ONLY be used inside functions
	// := declares + assigns


	// ===============================
	// 2. REASSIGNING A VALUE
	// ===============================

	d = 50 // ✅ reassignment is allowed
	fmt.Println("d after reassignment =", d)

	// ❌ ERROR (do NOT do this)
	// d := 60
	// error: no new variables on left side of :=


	// ===============================
	// 3. REDECLARATION RULES
	// ===============================

	x := 100
	fmt.Println("x =", x)

	// ❌ ERROR: redeclaration in same scope
	// x := 200

	// ✅ VALID: at least ONE new variable
	y, x := 300, 400
	fmt.Println("y =", y)
	fmt.Println("x =", x)


	// ===============================
	// 4. USING var AGAIN
	// ===============================

	var z int = 500
	fmt.Println("z =", z)

	// ❌ ERROR: redeclared variable
	// var z int


	// ===============================
	// 5. VARIABLE SCOPE (SHADOWING)
	// ===============================

	num := 10
	fmt.Println("num (outer) =", num)

	if true {
		num := 20 // NEW variable, different scope
		fmt.Println("num (inner) =", num)
	}

	fmt.Println("num (outer again) =", num)
}


// Rules for variable---------
// 1. :=  → declare + assign (ONLY inside functions)
// 2. =   → reassign value
// 3. You CANNOT redeclare a variable in the same scope
// 4. := works only if at least ONE variable is new
// 5. Same variable name is allowed in a different scope (shadowing)
