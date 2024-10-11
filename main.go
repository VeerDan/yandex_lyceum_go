package main

import (
    "fmt"
)

type Person struct {
    Name string
    Age int
    Address string
}

func (p Person) Print() {
    fmt.Printf("Name: %s\n", p.Name)
    fmt.Printf("Age: %d\n", p.Age)
    fmt.Printf("Address: %s\n", p.Address)
}
