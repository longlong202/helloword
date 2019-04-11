package main

import (
	"fmt"
)

func main() {
	fmt.Println("Helllo,world!")
	var s = "我爱China!"
	for i,ch := range []rune(s){
		fmt.Printf("(%d %c)",i,ch)
	}
	fmt.Println()
	x :=[]int{
		1, 2, 3,
		4, 5, 6,
	}
	for _,v:= range x{
		fmt.Println(v)
	}
}