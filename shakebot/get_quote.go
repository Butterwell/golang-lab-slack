//package shakebot
package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"

	"github.com/texttheater/golang-levenshtein/levenshtein"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

// Reading files requires checking most calls for errors.
// This helper will streamline our error checks below.
func check(e error) {
	if e != nil {
		panic(e)
	}
}

// readLines reads a whole file into memory
// and returns a slice of its lines.
func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

func readSonnets() []string {
	lines, err := readLines("sonnets.txt")
	check(err)
	return lines
}

func getRandomLine(lines []string) string {
	line := ""
	for len([]rune(line)) < 10 {
		line = lines[rand.Intn(len(lines))]
	}
	if strings.ContainsAny(line[len(line)-1:len(line)], ",:.;") {
		return line[:len(line)-1]
	}
	return line
}

//Iterate through all of the quotes and decide the one that is most similar to
//the user input.
func compareToQuotes(userInput string, quotes []string) (string, int) {
	//I'd like to declare bestMatch and useQuote in a different way, this is
	//hacky, but I can't figure out another way of doing it.
	bestMatch := 9999
	useQuote := ""
	for _, quote := range quotes {
		editDistance := levenshtein.DistanceForStrings([]rune(quote),
			[]rune(userInput),
			levenshtein.DefaultOptions)
		if editDistance < bestMatch {
			bestMatch = editDistance
			useQuote = quote
		}
	}
	return useQuote, bestMatch
}

func main() {
	lines := readSonnets()
	fmt.Println(getRandomLine(lines))
	fmt.Println(getRandomLine(lines))
	fmt.Println(lines[0])
	fmt.Println(lines[1])
	fmt.Println(lines[2])
	comparison := "when winters"
	quote, distance := compareToQuotes(comparison, lines)
	fmt.Println(quote, distance)
}
