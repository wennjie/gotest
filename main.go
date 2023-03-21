package main

import (
	"awesomeProject/test"
	"awesomeProject/test/cmd"
	"fmt"
)

func main() {
	test.Name("main-name-test-git3")
	//我实现了一个用户管理功能
	test.Name("main-name-test-git2")
	test.Age()
	//我实现了另外一个功能
	fmt.Println("1111")
	cmd.CmdTest()
	println()
}
