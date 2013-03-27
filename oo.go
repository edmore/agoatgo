package main

import(
"fmt"
)

type Person struct {
	Name, Surname string
}

func (p Person) getName() string {
	return p.Name
}

func (p *Person) setName(s string) {
	p.Name = s
}

func main() {
	p := &Person{Name : "Edmore"}
	fmt.Println(p.getName())
	p.setName("Tu")
	fmt.Println(p.getName())
}