package main

import (
    "fmt"
    "math/rand"
)

func main() {
    fmt.Printf("Rand: %d", rand.New(rand.NewSource(1)))
}
