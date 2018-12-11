package main
import "fmt"

type (
	TZ int,
	test string
)

func main() {
	var a, b TZ, t test = 3, 4, "abc"
	c := a + b
	fmt.Printf("c has the value %d", c)
}