package flowcontrol

import (
	"fmt"
	"log"
	"net/http"
)

//PanicPlay1 shows general panic behaviour with no code to handle it
func PanicPlay1() {
	a, b := 1, 0
	ans := a / b
	fmt.Println(ans)
}

//PanicPlay2 shows how panic can be used to show errors.
func PanicPlay2() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello Go!"))
	})
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err.Error())
	}
}

//PanicPlay3 shows how to recover panics
func PanicPlay3() {
	fmt.Println("start")
	defer fmt.Println("This was deffered")
	panic("Something bad happened")
	fmt.Println("end")

	//From this we see that we first execute main function, then deffered functions, then panic and then we handle return value
	//Therefore, any deffered calls to close resources for example will still close even though the program panics

}

//PanicPlay4 shows logging of errors after panic by using defer amd how the rest of the program in main still continues
func PanicPlay4() {
	fmt.Println("start")

	defer func() {
		if err := recover(); err != nil {
			log.Println("Error:", err)
			//With only the line above, the function above this in the call stack
			//(In our case main), will continue under the assumption that we have handeled the error
			//If we did NOT handle the error, we should retrow it here
			//Uncomment the line below to change so that the application panics again, this time killing the main.
			//panic(err) //Rethrow the panic inside the handler because the error wasd not handled. (Can also generate new panic if more suitable)
		}
	}() //The last () makes the anonymous function execute
	panic("Something bad happened")
	fmt.Println("end")
}
