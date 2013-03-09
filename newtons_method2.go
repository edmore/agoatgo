package main

import (
    "fmt"
)

func Sqrt(x float64) float64 {
    z := x/2.0
     for y:=0; y < 11; y++ {
        z = z - (((z*z) - x) / (2 * z))
    }
    return z
}

func main() {
    fmt.Println(Sqrt(16))
}