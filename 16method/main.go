package main

import "fmt"

func main() {
	fmt.Println("Structs in golang")
	// no inheritance in golang; No super or parent

	pri := User{"Prince", "agarwal@go.dev", true, 16}
	fmt.Println(pri)
	fmt.Printf("prince details are: %+v\n", pri)
	fmt.Printf("Name is %v and email is %v.\n", pri.Name, pri.Email)
	pri.GetStatus()
	pri.NewMail()
	fmt.Printf("Name is %v and email is %v.\n", pri.Name, pri.Email)

}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func (u User) GetStatus() {
	fmt.Println("Is user active: ", u.Status)
}

func (u User) NewMail() {
	u.Email = "test@go.dev"
	fmt.Println("Email of this user is: ", u.Email)
}
