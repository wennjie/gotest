package main

import (
	"awesomeProject/test"
	"awesomeProject/test/cmd"
	"fmt"
)

func main() {
	test.Name("main-name")
	test.Age()
	fmt.Println("1111")
	cmd.CmdTest()
	println()
}
