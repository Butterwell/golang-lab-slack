//package shakebot
package main

import (
    "bufio"
    "fmt"
    "os"
    "math/rand"
    "time"
    "strings"
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

func readSonnets() ([]string){
  lines, err := readLines("sonnets.txt")
  check(err)
  return lines
}

func getRandomLine(lines []string) (string){
  line := ""
  for (len([]rune(line)) < 10) {
    line = lines[rand.Intn(len(lines))]
  }
  if (strings.ContainsAny(line[len(line)-1:len(line)],",:.;")) {
    return line[:len(line)-1]
  }
  return line
}

/*
func main() {
  lines := readSonnets()
  fmt.Println(getRandomLine(lines))
  fmt.Println(getRandomLine(lines))
  fmt.Println(lines[0])
  fmt.Println(lines[1])
  fmt.Println(lines[2])
}
*/
