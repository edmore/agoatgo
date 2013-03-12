package main

import (
	"fmt"
	"math"
	"strings"
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

func WordCount(s string) map[string]int {
	m := make(map[string]int)
	split := strings.Fields(s)

	for _, v := range split{
		m[v] = m[v] + 1
	}
	return m
}


func fibonacci() func() int {
	x, y, z := 0, 1, 0
	return func() int {
		z, x, y = x, y, x+y
		return z
	}
}

func main() {
	p := Person{Name : "Edmore"}
        q := new(Person)
	q.Name = "Tu"
	r := &p
        r.Name = "Tu"
	s := []int{2,4,5,6,7}
        m := make(map[string]int)
        m["Age"] = 30

   	fmt.Println("Add 1 : ", add1(2))
	fmt.Println("Value of Pi : ", pi())
	fmt.Println("Fav OO language : ", languages())
	fmt.Println("Sum is : ", looper(3))
	fmt.Println("Look Ma a struct changed by a pointer : ", p.Name)
	fmt.Println("Look Ma using the new keyword: ", q.Name)
	fmt.Println("s: ", s[1:4])
	fmt.Println("Age : ", m["Age"])
        fmt.Println(WordCount("I ate a donut. Then I ate another donut."))

	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
