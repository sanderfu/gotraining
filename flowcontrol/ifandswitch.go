package flowcontrol

import (
	"fmt"
)

//FlowPlay1 is some initial flow testing
func FlowPlay1() {
	wordCount := map[string]int{
		"Hello": 1,
		"Bye":   1,
		"The":   76,
		"LOL":   7,
	}

	if count, ok := wordCount["BY"]; ok {
		fmt.Printf("The word %s was mentioned %d time(s)\n", "By", count)
	} else if count, ok := wordCount["Bye"]; ok {
		fmt.Printf("The word %s was mentioned %d time(s)\n", "Bye", count)
	} else {
		fmt.Printf("Nothing to mention\n")
	}

	switch i := 2 + 3; i {
	case 1, 5, 10:
		fmt.Println("one, five or ten")
	default:
		fmt.Println("Another number")
	}
	i := 10
	switch {
	case i <= 10:
		fmt.Println("less than or equal to ten")
		fallthrough //To also go to next case, PS NO logic, so will cause wrong things to be said here.
	case i <= 20:
		fmt.Println("smaller or equal to twenty")
	default:
		fmt.Println("Larger than 20")
	}

	//Typeswitch
	var x interface{} = 1
	switch x.(type) {
	case int:
		fmt.Println("x is an int")
		//break //Break out of the switch early, should be wrapped in logcial test.
		fmt.Println("This will print too")
	case float64:
		fmt.Println("x is a float64")
	case [2]int:
		fmt.Println("x is an int array with 2 elements")
	default:
		fmt.Println("x is another type")
	}
}
