package main

import (
    "fmt"
    "html"
)

func main() {
    const s = `"helllo html" <hello@example.com>`
    escapeString := html.EscapeString(s)
	fmt.Println(html.EscapeString(s))
    fmt.Println(html.UnescapeString(escapeString))
}
