package pointers

import (
	"fmt"
)

//PointerPlay1 shows pointer functionality
func PointerPlay1() {
	var a int = 42
	var b *int = &a
	fmt.Println(a, b)
	fmt.Println(a, *b) //Dereference b pointer
}

//PointerPlay2 ...
func PointerPlay2() {
	a := [3]int{1, 2, 3}
	b := &a[0]
	c := &a[1]
	fmt.Printf("%v %p %p\n", a, b, c)
	//Note: Pointer Arithmetic is NOT(!!!!) allowed for simplicity
	//Comment: Can use package "unsafe" to be allowed to use package arithmetic among other stuff
}

type myStruct struct {
	foo int
}

//PointerPlay3 ...
func PointerPlay3() {
	var ms *myStruct
	ms = &myStruct{foo: 42}
	fmt.Println(ms)

	var ms2 *myStruct
	fmt.Println(ms2)
	//It will be initialized to the "zero" value of pointer, nil
	ms2 = new(myStruct) //Cant initialize with this syntax
	fmt.Println(ms2)
	(*ms).foo = 42
	fmt.Println((*ms).foo)
	//Well this is great but the syntax is ugly AF
	//The compiler can help!
	ms.foo = 77
	fmt.Println(ms.foo)
	//This is a bit crazy, the compiler understands that we want the underlying field not a part of a pointer implicitly!

}
