package main

//lint:file-ignore U1000 Ignore all unused code, it's generated

import (
	"log"
	"os"
	"runtime"
	"strings"
)

//lint:file-ignore U1000 Ignore all unused code, it's generated

var (
	debugEnabled = os.Getenv("DEBUG") == "true"
	debugLogger  = log.New(os.Stdout, "DEBUG: ", log.Ldate|log.Ltime)
)

func debugLog(v ...any) {
	if debugEnabled {
		pc, _, _, ok := runtime.Caller(1)
		if !ok {
			debugLogger.Println(v...)
			return
		}
		if fn := runtime.FuncForPC(pc); fn != nil {
			name := fn.Name()
			if lastDot := strings.LastIndex(name, "."); lastDot != -1 {
				name = name[lastDot+1:]
			}
			args := append([]any{name + ":"}, v...)
			debugLogger.Println(args...)
		}
	}
}

func ReverseVowels(s string) string {
	// 345
	vowels := map[rune]bool{
		'a': true, 'e': true, 'i': true, 'o': true, 'u': true,
		'A': true, 'E': true, 'I': true, 'O': true, 'U': true,
	}
	runeSlice := []rune(s)
	i1, i2 := 0, len(runeSlice)-1
	for i1 < i2 {
		if !vowels[runeSlice[i1]] {
			i1++
		} else if !vowels[runeSlice[i2]] {
			i2--
		} else {
			runeSlice[i1], runeSlice[i2] = runeSlice[i2], runeSlice[i1]
			i1++
			i2--
		}
	}
	return string(runeSlice)
}
