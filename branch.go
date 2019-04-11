package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	const filename  = "src/helloword/abc.txt" //路径一定要在gopath下
	
	if contents,err := ioutil.ReadFile(filename); err !=nil{
		fmt.Println(err)
	} else {
		fmt.Printf("%s\n",contents)
	}


}
