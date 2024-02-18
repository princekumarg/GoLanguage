package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "welcome to userinput"
	fmt.Println(welcome)
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Input The Rating")
	input, _ := reader.ReadString('\n')
	fmt.Println("Thank For rating", input)
	fmt.Printf("Type of Rating %T", input)
}
