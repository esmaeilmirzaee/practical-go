package main

import (
	"github.com/esmaeilmirzaee/practical_go/datatypes/organization"
)

func main() {
	var p organization.Identifiable = &organization.Person{}
	println(p.ID())
}
