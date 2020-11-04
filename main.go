package main

import (
"fmt"
	"os"
)

func main() {
	switch os.Args[1] {
		case "console":
			fmt.Println(os.Args[2:])
		default:
		fmt.Println("what??")
	}
}