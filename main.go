package main

import (
	"fmt"

	"github.com/nomadcoders/nomadcoin/person"
)

func main() {
	nico := person.Person{}
	nico.SetDetails("nico", 21)
	fmt.Println("Maing 'nico'", nico)
}
