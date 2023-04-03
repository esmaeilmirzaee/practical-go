package main

import (
	"fmt"
	"github.com/esmaeilmirzaee/practical_go/datatypes/organization"
)

func main() {
	var p organization.Identifiable = &organization.Person{}
	var p2 = organization.NewPerson("John", "Doe")
	println(p2.FullName())
	println(p.ID())

	handler := "@esmaeilmirzaee"
	if err := p2.SetHandler(handler); err != nil {
		fmt.Printf("%s", err.Error())
	}
	println("Twitter handler: ", p2.TwitterHandler())
}
