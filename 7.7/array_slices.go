package main
import "fmt"

// slice 就是一个数组的引用 所以多个切片能共享数据 而且其长度不能超过数组
// slice 不需要指针指向 因为他本身就是个指针
// var slice1 []int = arr1[2:5] 其中 索引范围是 [2:5)
func main() {
	var arr1 [6]int
	var slice1 []int = arr1[2:5]

	for i := 0; i < len(arr1); i++ {
			arr1[i] = i
	}

	for i := 0; i < len(slice1); i++ {
		fmt.Printf("Slice at %d is %d\n", i, slice1[i])
	}

	fmt.Printf("The length of arr1 is %d\n", len(arr1))
	fmt.Printf("The length of slice1 is %d\n", len(slice1))
	fmt.Printf("The capacity of slice1 is %d\n", cap(slice1))

	slice1 = slice1[0:4]
	for i := 0; i < len(slice1); i++ {
		fmt.Printf("Slice at %d is %d\n", i, slice1[i])
	}

	fmt.Printf("The length of slice1 is %d\n", len(slice1))
	fmt.Printf("The capacity of slice1 is %d\n", cap(slice1))
}