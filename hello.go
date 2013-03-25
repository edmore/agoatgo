package main

import (
	"fmt"
	"math"
	"strings"
	"time"
)

type Person struct {
  Name string
  Surname string
}

func (p *Person) setSurname(s string){
	p.Surname = s
}

func (p *Person) getName() string {
	return p.Name
}

func add1(i int) int {
	return i + 1
}

func add1ch(i int, c chan int){
	c <- i + 1
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

func day_of_week() time.Weekday {
	return time.Now().Weekday()
}

func main() {
	p := Person{Name : "Edmore"}
        q := new(Person)
	q.Name = "Tu"
	r := &p
        r.Name = "Tu"
	r.setSurname("Moyo2")

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
        fmt.Println(day_of_week())
	fmt.Println(r)
	fmt.Println(r.getName())
	f := fibonacci()

	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}

	c := make(chan int)
	go add1ch(3, c)
        go add1ch(5, c)
	fmt.Println(<-c,<-c)

	// buffered channel
	ch := make(chan int, 2)
	ch <- 1
	ch <- 2 // doesnt block
	fmt.Println(<-ch)
	ch <- 3 // would have blocked had we not received
	fmt.Println(<-ch)
	fmt.Println(<-ch)

	messages := make(chan string)
        // closure - run in parallel so there is no blocking
        go func(){
		fmt.Println("hello")
		messages <- "done"
	}()
	fmt.Println("Do something while you wait ...")
        fmt.Println(<-messages)
}
