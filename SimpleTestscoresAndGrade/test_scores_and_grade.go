package main

import "fmt"

func main() {
    var (
        avg float64
        letter_grade string
    )
    scores := [3]float64{}
    scores = get_scores()
    avg = average(scores)
    letter_grade = determine_grade(avg)
    fmt.Printf("Grade: %s\n", letter_grade)
}
func get_scores() (scores [3]float64) {
    for i := 0; i < len(scores); i++ {
    fmt.Println("Enter test score", i)
        fmt.Scanln(&scores[i])
    } 
    return
}

func average(scores [3]float64) (avg float64){
    total := 0.0
    for i:= 0; i < len(scores); i++ {
       total += scores[i] 
    }
    avg = total / float64(len(scores))
    return
}

func determine_grade(avg float64) (letter_grade string){
    if avg >= 90 && avg <= 100 {
        letter_grade = "A"
    }
    if avg >= 80 && avg < 90 {
        letter_grade = "B"
    }
    if avg >= 70 && avg < 80 {
        letter_grade = "C"
    }
    if avg >= 60 && avg < 70 {
        letter_grade = "D"
    }
    if avg >= 0 && avg < 60 {
        letter_grade = "F"
    }
    return
}
