package main

import (
	"log"
	"os"
	"strings"
	"unicode"
)

func puzzle1(data string) int {
    total := 0
    for _, s := range strings.Split(data, "\n") {
        s = strings.TrimSpace(s)
        var digits []int
        for _, c := range s {
            if unicode.IsDigit(c) {
                digits = append(digits, int(c - '0'))
            }
        }

        number := digits[0] * 10 + digits[len(digits) - 1]
        total += number
    }

    return total
}

func remapInput(s string) string {
    // To account for overlaps, pad replacement strings
    s = strings.ReplaceAll(s, "one",   "o1e")
    s = strings.ReplaceAll(s, "two",   "t2o")
    s = strings.ReplaceAll(s, "three", "t3e")
    s = strings.ReplaceAll(s, "four",  "f4r")
    s = strings.ReplaceAll(s, "five",  "f5e")
    s = strings.ReplaceAll(s, "six",   "s6x")
    s = strings.ReplaceAll(s, "seven", "s7n")
    s = strings.ReplaceAll(s, "eight", "e8t")
    s = strings.ReplaceAll(s, "nine",  "n9e")

    return s
}

func puzzle2(data string) int {
    data = remapInput(data)
    return puzzle1(data)
}

func main() {
    data, err := os.ReadFile("./input.txt")
    if err != nil {
        log.Fatalf("Unable to read input file: %s\n", err)
        return
    }

    s_data := string(data)
    log.Println("Puzzle 1:", puzzle1(s_data))
    log.Println("Puzzle 2:", puzzle2(s_data))
}
