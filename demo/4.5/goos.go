package main

import (
	"fmt"
	"runtime"
	"os"
)

//Println 是能自动使用格式化标识符%v对自负床进行格式化，还会在每个参数之间自动增加空格

func main() {
	// var a string = "abc"		//编译期会排除没有使用到的局部变量
	var goos string = runtime.GOOS
	fmt.Println("The operating system is", goos)
	path := os.Getenv("PATH")
	fmt.Println("Path is", path)
}