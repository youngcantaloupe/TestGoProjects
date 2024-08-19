package main

import "fmt"

func main() {
    var num int 

    fmt.Println("Enter number: ")
    fmt.Scanln(&num)
    DeterminePrime(num)
}

func DeterminePrime(num int) {
    var is_div bool
    div := 2
    for div < num && !is_div {
        if num % div == 0 {
            is_div = true
        }
        div++
    }
    if is_div {
        fmt.Println("Not Prime")
    }
    if !is_div {
        fmt.Println("Prime")
    }
}
