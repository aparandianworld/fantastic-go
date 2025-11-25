package main

import (
	"fmt"
)

type Role string

const (
	Viewer    Role = "viewer"
	Developer Role = "developer"
	Admin     Role = "admin"
)

type User struct {
	Login string
	Role  Role
}

func promote(u *User, r Role) {

	u.Role = r
}

func main() {
	aaron := User{
		Login: "aaron",
		Role:  Developer,
	}

	fmt.Println(aaron)

	promote(&aaron, Admin)

	fmt.Println(aaron)
}
