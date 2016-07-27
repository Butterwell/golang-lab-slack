//package shakebot
package main

import (
    "bufio"
    "fmt"
    "os"
    "math/rand"
)

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

  lineCount := 0

  var lines []string
  scanner := bufio.NewScanner(file)
  for scanner.Scan() {
    lineCount++
    //if (len([]rune(scanner.Text())) > 5) {
    lines = append(lines, scanner.Text())
    //}
  }
  fmt.Println("number of lines:", lineCount)
  return lines, scanner.Err()
}

func readSonnets() ([]string){
  lines, err := readLines("sonnets.txt")
  check(err)
  return lines
}

func getRandomLine(lines []string, r *rand.Rand) (string){
  line := ""
  for (len([]rune(line)) < 5) {
    line = lines[r.Intn(len(lines))]
  }
  return line
}
/*
func main() {
  lines := readSonnets()
  r := rand.New(rand.NewSource(99))
  fmt.Println(getRandomLine(lines, r))
  fmt.Println(lines[7])
}
*/
