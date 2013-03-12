package main

import (
	"fmt"
	"math"
)

func add1(i int) int {
	return i + 1
}

func pi() float32 {
	return math.Pi
}

func languages() string {
	lang := "Ruby"
	return lang
}

func looper(max int) int {
	sum := 0
	for i := 0; i <= max; i++ {
		if i != 1 {
			sum += i
		}
	}
	return sum
}

type Person struct {
  Name string
  Surname string
}

func main() {
	p := Person{Name : "Edmore"}
        q := new(Person)
	q.Name = "Tu"
	r := &p
        r.Name = "Tu"
	s := []int{2,4,5,6,7}

   	fmt.Println("Add 1 : ", add1(2))
	fmt.Println("Value of Pi : ", pi())
	fmt.Println("Fav OO language : ", languages())
	fmt.Println("Sum is : ", looper(3))
	fmt.Println("Look Ma a struct changed by a pointer : ", p.Name)
	fmt.Println("Look Ma using the new keyword: ", q.Name)
	fmt.Println("s: ", s[1:4])

}
