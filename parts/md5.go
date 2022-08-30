package main

import (
    "fmt"
    "crypto/md5"
)

func main() {
    data := []byte("hello")
	fmt.Printf("%x", md5.Sum(data))
}
