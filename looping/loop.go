package looping

import (
	"fmt"
)

//LoopPlay1 tests out different ways to use loops
func LoopPlay1() {
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	i := 0
	for i < 5 {
		i++
		fmt.Println(i)
	}

	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			continue //Similar to c/Python etc
		}
		fmt.Println(i)
	}
Loop: //Tag
	for i := 1; i <= 3; i++ {
		for j := 1; j <= 3; j++ {
			fmt.Println(i * j)
			if i*j >= 3 {
				break Loop //Break out of bboth inner and outer loop to the tag above
			}
		}
	}

	s := []int{1, 2, 3} //This is a SLICE, not an array as we have not defined its length in the[]
	for k, v := range s {
		fmt.Println(k, v)
	}

}
