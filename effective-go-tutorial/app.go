/*
Greets you with hello message.
This is an example package for the go doc tutorial: https://golang.org/doc/effective_go
*/
package main

import (
	"fmt"
)

//A structure representing a kind person
type KindPerson struct {
	name   string
	age    int
	phrase string
}

//A structure representing a boring person which has nothing to say
type BoringPerson struct {
	name string
	age  int
}

//Variables can be grouped to show relationships
//A group of good people
var (
	eddie           KindPerson
	meddie          KindPerson
	concertLocation string
)

//A group of bad people
var (
	eran         BoringPerson
	homeLocation string
)

//Say function is used to make kind person speak
func (p KindPerson) say() {
	fmt.Printf("%s of the age of %d says \"%s\"\n", p.name, p.age, p.phrase)
}

func main() {
	p := KindPerson{"Jack", 19, "Hi!"}
	p.say()
}
