package main

func main() {
	//var a int
	var b int32
	//a = 15
	//b = a + a   错误  因为go不允许不同类型之间的混合使用
	b = b + 5		//正确  因为go允许常量之间的混合使用
}