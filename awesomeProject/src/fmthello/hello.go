package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("hello world" + os.Args[0])
	fmt.Print(os.Args)
}
