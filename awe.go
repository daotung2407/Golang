package main

import (
	"bufio"
	"fmt"
 	"os"
)
// định nghĩa hàm

func main ()  {
	 fmt.Println("mời bạn nhập tên: ");
	 var reader *bufio.Reader = bufio.NewReader(os.Stdin)

	 name, _ := reader.ReadString('\n')
	fmt.Println("xin chào: -> " + name)
}