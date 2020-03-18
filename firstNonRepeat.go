package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

const CHARSET = "abcde"

// https://stackoverflow.com/a/31832326
func RandStringBytes(n int) string {
	b := make([]byte, n)
	rand.Seed(time.Now().UTC().UnixNano())
	for i := range b {
		b[i] = CHARSET[rand.Intn(len(CHARSET))]
	}
	return string(b)
}

func firstNonRepeatViaMap(s string) string {
	letterOccurrencesMap := make(map[string]int)
	for _, char := range s {
		c := string(char)
		if _, exists := letterOccurrencesMap[c]; !exists {
			letterOccurrencesMap[c] = 1
		} else {
			letterOccurrencesMap[c] += 1
		}
	}
	for _, char := range s {
		c := string(char)
		if letterOccurrencesMap[c] == 1 {
			return c
		}
	}

	return "_"
}

func firstNonRepeatViaStrMethods(s string) string {
	for _, char := range s {
		c := string(char)
		if strings.Index(s, c) == strings.LastIndex(s, c) {
			return c
		}
	}
	return "_"
}

func main() {
	// generate string
	sLength := 10
	s := RandStringBytes(sLength)
	fmt.Println(s)
	start := time.Now()
	fmt.Println(firstNonRepeatViaMap(s))
	totalTime := time.Since(start)
	fmt.Printf("took %s to find the first nonrepeating char via enumeration map\n", totalTime)

	start = time.Now()
	fmt.Println(firstNonRepeatViaStrMethods(s))
	totalTime = time.Since(start)
	fmt.Printf("took %s to find the first nonrepeating char via string methods\n", totalTime)
}
