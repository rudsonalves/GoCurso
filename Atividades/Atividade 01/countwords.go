package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func lowCase(line string) string {
	return strings.ToLower(line)
}

// remove caracteres de pootuação, estrangeiros de e outros
// do início de do final das palavras
func removeChars(word string) string {
	runeW := []rune(word)
	if len(runeW) == 0 {
		return ""
	}

	// remove caracteres do final da []rune
	for {
		switch lastRune(runeW) {
		case ':', '.', ',', '!', '?', ';', '£', '©', '"', '–', ')', 'ª', '%', '/':
			runeW = removeLastRune(runeW)
			if len(runeW) == 0 {
				return ""
			}
			continue
		}
		if lastRune(runeW) >= 48 && lastRune(runeW) <= 57 {
			runeW = removeLastRune(runeW)
			if len(runeW) == 0 {
				return ""
			}
			continue
		}
		if lastRune(runeW) > 1000 {
			runeW = removeLastRune(runeW)
			if len(runeW) == 0 {
				return ""
			}
			continue
		}
		break
	}

	// remove caracteres do início da []rune
	for {
		r0 := runeW[0]
		if r0 == '(' || r0 == '-' || r0 == 12 || r0 == 186 || (r0 >= 48 && r0 <= 57) {
			runeW = runeW[1:]
			if len(runeW) == 0 {
				return ""
			}
			continue
		}

		if r0 > 1000 {
			runeW = runeW[1:]
			if len(runeW) == 0 {
				return ""
			}
			continue
		}
		break
	}

	return string(runeW)
}

// retorna o último rune
func lastRune(list []rune) rune {
	lastIndex := len(list) - 1
	if lastIndex < 0 {
		log.Panic("lastRune: lenght is zero")
	}
	return list[lastIndex]
}

// remove a última rune da slice
func removeLastRune(list []rune) []rune {
	lastIndex := len(list) - 1
	if lastIndex < 0 {
		log.Panic("removeLastRune: lenght is zero")
	}
	return list[:lastIndex]
}

// retorna o índicd do menor int
func getLast(list map[string]int) string {
	id, n := "", 99999
	for w, v := range list {
		if v < n {
			n = v
			id = w
		}
	}
	return id
}

func main() {
	read, err := os.ReadFile("A Riqueza das Nacoes - Adam Smith.txt")
	if err != nil {
		log.Fatal(err)
	}

	lines := strings.Split(string(read), "\n")

	words := map[string]int{}
	for _, line := range lines {
		book_words := strings.Split(line, " ")
		for _, w := range book_words {
			w = removeChars(lowCase(w))
			if len(w) < 1 {
				continue
			}
			if words[w] != 0 {
				words[w]++
			} else {
				words[w] = 1
			}
		}
	}

	fmt.Println("Total:", len(words))

	top10 := map[string]int{}
	last, lenght := "", 0
	for w, n := range words {
		if len(w) < 3 {
			continue
		}

		if lenght < 10 {
			top10[w] = n
			last = getLast(top10)
			lenght++
			continue
		}

		if n > top10[last] {
			delete(top10, last)
			last = w
			top10[w] = n
			last = getLast(top10)
			continue
		}
	}

	for w, n := range top10 {
		fmt.Printf("%30s: %5d\n", w, n)
	}
}
