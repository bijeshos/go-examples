package main

import (
	"fmt"
	"os"
)

func main() {

	fmt.Println("started program")
	fmt.Println("about to exit successfully")
	os.Exit(0) //0 is success codes, non-zero is considered as error
	//os.Exit(1) //non-zero is considered as error

}
