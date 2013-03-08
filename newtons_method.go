package main

import "fmt"

func Sqrt(x float64) float64 {
    guess := x/2.0

    for y:=0; y < 11; y++ {
        guess = (guess + x/guess)/2
    }
    return guess
}

func main() {
    fmt.Println(Sqrt(36))
}
