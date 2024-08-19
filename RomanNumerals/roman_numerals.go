package main

import "fmt"

func main() {
    var (
    count_to int
    numeral string
    )
    fmt.Println("Enter a number to count to. (MAX = 10)")
    fmt.Scanln(&count_to)
   for count_to <= 0 || count_to > 10 { 
    fmt.Println("Error, number entered is not 1-10")
    fmt.Scanln(&count_to)
    }
    for i := 1; i <= count_to; i++ {
        switch i {
        case 1:
           numeral = "I"
        case 2:
           numeral = "II"
        case 3:
           numeral = "III"
        case 4:
           numeral = "IV"
        case 5:
           numeral = "V"
        case 6:
           numeral = "VI"
        case 7:
           numeral = "VII"
        case 8:
           numeral = "VIII"
        case 9:
           numeral = "IX"
        case 10:
           numeral = "X"
        default:
            numeral = "Problem"

        }
        if i < count_to {
            fmt.Print(numeral + ", ")
        }
        if i == count_to {
            fmt.Print(numeral + ". ")
        }
    }
    fmt.Println()
}
