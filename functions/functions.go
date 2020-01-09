package functions

import (
	"fmt"
)

func sayGreeting(greeting, name string) {
	fmt.Println(greeting, name)
}

//Variametic parameter (MUST be at end) creates slice.
func sum(msg string, values ...int) int {
	fmt.Println(values)
	result := 0
	for _, v := range values {
		result += v
	}
	fmt.Println(msg, result)
	return result
}

func sumReturnPointer(values ...int) *int {
	fmt.Println(values)
	result := 0
	for _, v := range values {
		result += v
	}
	return &result //Automatically puts the result on the heap because we return a pointer.
}

//Named returns
func sumV2(values ...int) (result int) {
	fmt.Println(values)
	for _, v := range values {
		result += v
	}
	return
	//This method is not used often because it is a bit hard to read, especially in long functions
}

func divide(a, b float64) (float64, error) {
	if b == 0.0 {
		return 0.0, fmt.Errorf("Cannot divide by zero")
	}
	return a / b, nil //Nil error when nothing goes wrong
	//Will return Inf if b is 0! (No crash unless we choose to panic)

}

func divideUsage() {
	d, err := divide(5.0, 0.0)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(d)
}

//PlayFunctions tests function functionality
func PlayFunctions() {
	divideUsage()

	//Introducing anonymous function
	func() {
		fmt.Println("Hello Go!")
	}() //Need () to immideately invoke function, needed or compiler wont know what to do with it.

	for i := 0; i < 5; i++ {
		func(i int) { //Pass the i vaariable to pass it into the function, to get correct behaviour with asynchronous operation
			fmt.Println(i)
		}(i)
	}

	//Defining a function as a variable we freely can pass around in our application
	var f func() = func() {
		fmt.Println("Hello Go!")
	}
	f()

	/*
		var divide func(float64, float64) (float64, error)
		divide = func(a, b float64) (float64, error) {
			if b == 0.0 {
				return 0.0, fmt.Errorf("Cannot divide by zero")
			}
			return a / b, nil
		}
		d, err := divide(5.0, 0.0)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(d)
	*/
	//Note: When functions are declared as variables, make sure to not use them before they are declared as that will not work.

	g := greeter{
		greeting: "Hello",
		name:     "Go",
	}
	g.greet()
	fmt.Println(g.name)
	g.greetV2()
	fmt.Println(g.name)
}

type greeter struct {
	greeting string
	name     string
}

type counter int

//Method: Special type of function that executes in copntext of a type.
//Declearing method, this will be a function for the type greeter
//g is value reciever -> gets copy of type
func (g greeter) greet() {
	fmt.Println(g.greeting, g.name)
	g.name = "Changed?"
	//Showing that is a copy fo the greeter object that is imported.
}

//g is pointer reciever -> gets pointer to type.
func (g *greeter) greetV2() {
	fmt.Println(g.greeting, g.name)
	g.name = "Changed!"
}
