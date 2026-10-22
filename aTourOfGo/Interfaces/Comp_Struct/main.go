package main

import (
	"fmt"
)

type Address struct {
	Street, City, Country string
}

func (a Address) String() string {
	return fmt.Sprintf("%s, %s, %s", a.Street, a.City, a.Country)
}

type Contact struct {
	Name, Email string
	Address Address
}

func (c Contact) String() string {
	return fmt.Sprintf("%s <%s>\n%s", c.Name, c.Email, c.Address)
}

func main() {
	addr := Address{"123 Main St", "Springfield", "USA"}
	contact := Contact{"Bob", "bob@example.com", addr}

	fmt.Println(contact)
}