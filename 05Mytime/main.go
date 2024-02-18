package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

func main() {
	fmt.Println("Welcome to maths in golang")
	myRandomNum, _ := rand.Int(rand.Reader, big.NewInt(5))
	fmt.Println(myRandomNum)

}
