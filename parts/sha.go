package main

import (
    "fmt"
    "crypto/sha256"
)

func main() {
    h := sha256.New()
    h.Write([]byte("hello"))
    fmt.Printf("%x", h.Sum(nil))
}



