package main

//lint:file-ignore U1000 Ignore all unused code, it's generated
import (
	"fmt"
	"log"
	"os"
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
	fileContent, err := os.ReadFile("leetcode_202503.go")
	if err != nil {
		log.Fatal(err)
	}

	// numbers := extractNumbersFromFunctionNames(string(fileContent))

	// fmt.Println("Extracted Numbers:", numbers)

	re := regexp.MustCompile(`(?ms)^func.*?\n\s*//\s*(\d+)\n`)
	matches := re.FindAllStringSubmatch(string(fileContent), -1)
	m := make(map[string]struct{})
	for _, match := range matches {
		// fmt.Println(match[1]) // Extracted number
		m[match[1]] = struct{}{}
	}

	solvedProblems, err := os.ReadFile("solved_problems.txt")
	if err != nil {
		log.Fatal(err)
	}

	solvedSet := make(map[string]struct{})
	for _, line := range strings.Split(string(solvedProblems), "\n") {
		line = strings.TrimSpace(line)
		if line != "" {
			solvedSet[line] = struct{}{}
		}
	}

	for num := range m {
		if _, exists := solvedSet[num]; !exists {
			fmt.Println(num)
		}
	}
}
