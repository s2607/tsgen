package main

import "fmt"
import "bufio"
import "os"

func main() {

	fmt.Printf("hello world")
	reader := bufio.NewReader(os.Stdin)
	text, e := reader.ReadString('\n')
	if e == nil {
		fmt.Println(text)
	} else {
		fmt.Println(e)
	}
}
