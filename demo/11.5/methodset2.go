package main

import (
		"fmt"
)

type Lener interface {
		Len() int
}

type Appender interface {
		Append(int)
}

type List []int

func (l List) Len() int {
		return len(l)
}

func (l *List) Append(val int) {
		*l = append(*l, val)
}

func CountInto(a Appender, start, end int) {
		for i := start; i <= end; i++ {
				a.Append(i)
		}
}


func LongEnough(l Lener) bool {
		return l.Len()*10 > 42
}

func main() {
		var lst List

		// CountInto(lst, 1, 10)	不可行 因为Append方法是定义在List指针上

		if LongEnough(lst) {
				fmt.Printf("- lst is long enough\n")
		}

		plst := new(List)
		CountInto(plst, 1, 10)
		if LongEnough(plst) {					//Len方法虽然定义在对象上 可是指针会自动解引用
				fmt.Printf("- plst is long enough\n")
		}
}