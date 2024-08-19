package main 

import (
    "fmt"
    "os"
    "encoding/json"
)

type Names struct {
    FirstName string
    LastName string
}

    
func (n *Names) load() {
    data, err := os.ReadFile("names.json")

    if err != nil {
        fmt.Println("No names loaded")
        return
    }

    err = json.Unmarshal(data, &n)
    if err != nil {
        panic(err)
    }
}

func (n *Names) save() {
    data, err := json.Marshal(n)
    if err != nil {
        panic(err)
    }

    err = os.WriteFile("names.json", data, 0644)
    if err != nil {
        panic(err)
    }
}

func (n *Names) modify(user_input string) {
  n.FirstName = user_input
}

func main() {

var names = []Names {
    {FirstName: "Porter", LastName: "Sampson"},
}
    //var user_input string
    for i := 0; i < len(names); i++ {
    names[i].load()
    }
    //fmt.Print("Add first name: ")
    //fmt.Scanln(&user_input)
    //names[0].modify(user_input)
    fmt.Println(names)
    names[0].save()
}
