package flowcontrol

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

//DeferPlay1 displays defer functionality
func DeferPlay1() {
	fmt.Println("start")
	defer fmt.Println("middle") //Defer does NOT but function at end of main, it executes it after main is done but before main returns
	fmt.Println("end")
}

//DeferPlay2 displays the LIFO order of defer
func DeferPlay2() {
	defer fmt.Println("start")
	defer fmt.Println("middle")
	defer fmt.Println("end")
}

//DeferPlay3 shows real usage of defer
func DeferPlay3() {
	res, err := http.Get("http://www.google.com/robots.txt")
	if err != nil {
		//Error with import
		log.Fatal(err)
	}
	//We close before read to not forget it, this is the most common usecase for defer
	//Warning: If has a lot of resources, A LOT of files will be open for a long time!

	defer res.Body.Close()
	robots, err := ioutil.ReadAll(res.Body)
	if err != nil {
		//Error with read process
		log.Fatal(err)
	}
	fmt.Printf("%s", robots)
}

//DeferPlay4 shows that deferred function takes argument at the time the defer is called, not when the function is executed
func DeferPlay4() {
	a := "start"
	defer fmt.Println(a)
	a = "end"

}
