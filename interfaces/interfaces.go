package interfaces

import (
	"fmt"
)

type writer interface {
	//Interfaces dont describe data, they describe behaviour
	write([]byte) (int, error)
}

type consoleWriter struct{}

type incrementer interface {
	increment() int
}

//Creating a new type that is not a struct!
type intCounter int

//Implementing the interface incrementer for the IntCounter type
func (ic *intCounter) increment() int {
	*ic++
	return int(*ic)
}

//Implementing the writer interface for the ConsoleWriter type
func (cw consoleWriter) write(data []byte) (int, error) {
	n, err := fmt.Println(string(data))
	return n, err
}

//PlayInterfaces tests out interface functionality
func PlayInterfaces() {
	var w writer = consoleWriter{}
	w.write([]byte("Hello Go!")) //This line does not care what it is writing to!

	myInt := intCounter(0)
	var inc incrementer = &myInt
	for i := 0; i < 10; i++ {
		fmt.Println(inc.increment())
	}
}
