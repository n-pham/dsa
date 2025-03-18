package main

import (
    "fmt"
	"io/ioutil"
    "log"
    "regexp"
    "strings"
)

func extractNumbersFromFunctionNames(fileContent string) []string {
    // Regular expression to match function names with numbers at the end (no `_suffix`)
    re := regexp.MustCompile(`func\s+[a-zA-Z]+(\d+)(_[a-zA-Z0-9])*\s*\(`)

    var numbers []string
    matches := re.FindAllStringSubmatch(fileContent, -1)
    for _, match := range matches {
        number := match[1]
		rest := match[2]
		fmt.Println(number, rest)
        // Exclude function names with `_` suffix
        if !strings.Contains(rest, "_") {
            numbers = append(numbers, number)
        }
    }
    return numbers
}

func main() {
    fileContent, err := ioutil.ReadFile("leetcode_202503.go")
    if err != nil {
        log.Fatal(err)
    }

    numbers := extractNumbersFromFunctionNames(string(fileContent))

    fmt.Println("Extracted Numbers:", numbers)
}