package main
import "fmt"

func main() {
	a()
}

func function1() {
		fmt.Printf("In function1 at the top\n")
		defer function2()
		fmt.Printf("In function1 at the bottom!\n")
}

func function2() {
	fmt.Printf("Function2: Deferred until the end of the calling function")
}

func a() {
	i := 0
	defer fmt.Println(i)
	i++
	fmt.Println("test: ", i)
	return
}