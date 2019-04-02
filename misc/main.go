package main

import "C"
import (
	"fmt"

	_ "misc/db"
)

func main() {

}

//export gotest
func gotest(s *C.char) {
	defer func() {
		fmt.Println("i am in defer")
	}()
	fmt.Println("this is go ", C.GoString(s))
}
