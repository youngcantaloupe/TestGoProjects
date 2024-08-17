package main

import (
    "fmt"
    "encoding/json"
    "os"
    "bufio"
    "strings"
)

type List struct {
    Item []string 
}

func (l List) display() {
    for i, item := range l.Item {
        fmt.Printf("%d: %s\n", i + 1, item)
    }
}

func (l *List) modify(user_input string) {
    l.Item = append(l.Item, user_input)
}

func (l *List) delete() {

    var delete_index int
    
    for i, v := range l.Item{
        fmt.Printf("Index: %v, Value: %v \n", i, v)
    }

    fmt.Print("Enter index to delete: ")
    fmt.Scanln(&delete_index)

    if delete_index < 0 || delete_index >= len(l.Item) {
        fmt.Println("Invalid index, try again.")
    } else {
        l.Item = append(l.Item[:delete_index], l.Item[delete_index + 1:]...)
    }
}

func (l *List) load() {
    data, err := os.ReadFile("list.json")

    if err != nil {
        fmt.Println("No list loaded")
        return
    }

    err = json.Unmarshal(data, &l)
    if err != nil {
        panic(err)
    }
}

func (l *List) save() {

    data, err := json.MarshalIndent(l, "", "   ")
    if err != nil {
        panic(err)
    }

    err = os.WriteFile("list.json", data, 0644)
    if err != nil {
        panic(err)
    }
}

func main() {
    var (
        list List
        menu_selec string
        running bool = true
    )

    reader := bufio.NewReader(os.Stdin)

    list.load()

    for running {
        fmt.Println("--------------------------------")
        fmt.Println("|             MENU             |")
        fmt.Println("|  Add | List | Delete | Quit  |")
        fmt.Println("--------------------------------")
        fmt.Print("Select: ")
        fmt.Scanln(&menu_selec)

        switch menu_selec {
        case "add":
            fmt.Print("Add: ")
            user_input, _ := reader.ReadString('\n')
            user_input = strings.TrimSuffix(user_input, "\n")
            list.modify(user_input)
        case "list":
            list.display()
        case "delete":
            list.delete()
        case "quit":
            running = false
        default:
            fmt.Println("Unknown command, try lowercase")
        }
    }
    list.save()
}

