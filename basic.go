package main

import (
	"fmt"
	"math"
	"math/cmplx"
)

func variable()  {
	a,b,c,d :=3,5,"hello",true
	fmt.Println(a,b,c,d)
}

func euler()  {
	//c := 3+4i
	//println(cmplx.Abs(c))//复数，输出绝对值
	fmt.Printf("%.3f\n",cmplx.Exp(1i * math.Pi) +1)
}

func Fconst()  {
	const (
		filename 	= "test.log"
		key 		= "360"
	)
	fmt.Println(filename,key)
}

func enums()  {
	const (
		cpp = iota
		java
		_
		python
		golang
	)
	println(cpp,java,python,golang)

	const  (
		b = 1<<(10 * iota) //自增值，公式复用
		kb
		mb
		gb
		tb
		pb
	)
	println(b,kb,mb,gb,tb,pb)
}
func main() {
	variable()
	euler()
	Fconst()
	enums()
	println(1<<20)
}