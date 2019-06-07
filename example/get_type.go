package main

import (
    "fmt"
    "reflect"
)

func main() {
    fmt.Println("The type of [42] is:      ", reflect.TypeOf(42))
    fmt.Println("The type of [3.1415] is:  ", reflect.TypeOf(3.1415))
    fmt.Println("The type of [true] is:    ", reflect.TypeOf(true))
    fmt.Println("The type of [be kind] is: ", reflect.TypeOf("be kind"))
}
