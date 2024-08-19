package main

import (
    "fmt"
    "sort"
)

func main() {
    names:= []string{"name1", "name2", "name3"}
    get_names(names)
    sort_names(names)
    for i := 0; i < len(names); i++ {
        fmt.Println(names[i])
    }
}

func get_names(n []string) (nu []string) {
    for i:= 0; i < len(n); i++ {
        fmt.Println("Enter name ", i + 1)
        fmt.Scanln(&n[i])
        fmt.Println()
    }
    return
}

func sort_names(n []string) (ns []string){
    sort.Strings(n)
    return
}
