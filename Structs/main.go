package main

import (
	"encoding/json"
	"fmt"
)

/*
   [*] Interface
   [*] Structs
   [*] Types,Type Assertions Switches

*/
type Person struct{
    Name string `json:"Name"`
    Age uint8   `json:"age"`

}
type Employee struct{
    Id uint8
    Person Person
    Salary int
}
type rect struct{
    Length int
    Width int

}
func main() {
    var speaker Speaker
    speaker = Dog{Name: "Buddy"}
    fmt.Println(speaker.Speak()) // Outputs: Woof!

    speaker = Human{Name: "Alice"}
    fmt.Println(speaker.Speak()) // Outputs: Hello!
    z := rect{Length: 4,Width: 4}
    //l := new(rect)
    bob := Employee{
        Id:12,
        Person: Person{
            Name: "Bob",
            Age: 23,
        },
        Salary: 234,
    }
    bob = Employee{12,Person{"Bob",23},234}
    fmt.Println(bob)
    susan := Person{
        Name: "Susan",
        Age: 24,
    }
    data,err := json.Marshal(susan)
    if err != nil{
        panic(err)
    }
    fmt.Println(string(data))
    fmt.Print(z.GetArea())
}
func (r rect) GetArea() int{
    return r.Width * r.Length
}
type Speaker interface {
    Speak() string
}

type Dog struct {
    Name string
}

func (d Dog) Speak() string {
    return "Woof!"
}

type Human struct {
    Name string
}

func (h Human) Speak() string {
    return "Hello!"
}


