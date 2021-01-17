package main //包，表明代码所在的

import (
	"fmt"
	"os"
)

func main(){
	fmt.Println(os.Args)
	if (len(os.Args) == 1) {
		fmt.Println("111")
	} else {
		fmt.Println("222")
	}
	fmt.Println("Hello World")
	os.Exit(-1)
}
