package main

import (
    "fmt"
    // "math"
    // "slices"
	// "strings"
)

func ShareWith(name string) string {
	if name == "" { name = "you"}
	return fmt.Sprintf("One for %s, one for me.", name)
}

func Convert(number int) string {
	result := ""
	if number % 3 == 0 { result += "Pling" }
	if number % 5 == 0 { result += "Plang" }
	if number % 7 == 0 { result += "Plong" }
	if result == "" { result = fmt.Sprint(number) }
	return result
}

func main() {
	fmt.Println(Convert(27), Convert(105), Convert(52))
    // fmt.Println(ShareWith("Bob"))
	// fmt.Println(ShareWith(""))
}