package main

import "fmt"

func main() {
	fmt.Println("User value")

	prince := User{"prince", "agarwal.dev", true, 1}
	fmt.Println(prince)
	fmt.Printf("prince details are: %+v\n", prince)
	fmt.Printf("Name is %v and email is %v.", prince.Name, prince.email)

}

type User struct {
	Name  string
	email string
	male  bool
	id    int
}
