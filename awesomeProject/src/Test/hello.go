package main

import (
	// 需在此处添加代码。[1]
	"flag"
	"fmt"
)

var name string

func init() {
	// 需在此处添加代码。[2]
	flag.StringVar(&name, "name", "everyone", "the greeting")
}

func main() {
	// 需在此处添加代码。[3]
	flag.Parse()
	fmt.Printf("Hello, %s!\n", name)
}
