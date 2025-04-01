package main

import (
	"flag"
	"fmt"
	"os"
)

func main(){
  line := flag.Bool("l", false, "count line")
  word := flag.Bool("w", false, "count words")
  byte := flag.Bool("c", false, "count bytes")
  char := flag.Bool("m", false, "count chars")

  flag.Parse()

  args := os.Args
  fileName := args[len(args)-1]

  file, err := os.Open(fileName)
  if err != nil {
    fmt.Printf("gowc: %s: %s", fileName, err)
    os.Exit(1)
  }
  defer file.Close()

  if flag.NFlag() == 0 {
    lineCount, err := countLinesInFile(file)
    if err != nil {
      fmt.Println("Error counting lines:")
      os.Exit(1)
    }

    file, err := os.Open(fileName)
    wordCount, err := countWordsInFile(file)
    if err != nil {
      fmt.Println("Error counting words:", err)
      os.Exit(1)
    }

    file, err = os.Open(fileName)
    byteCount, err := countBytesInFile(file)
    if err != nil {
      fmt.Println("Error counting bytes:", err)
      os.Exit(1)
    }

    fmt.Printf("%6d %6d %6d %6s\n", lineCount, wordCount, byteCount, fileName)
    defer file.Close()
    return
  }

  output := ""

  if *line {
    file, err := os.Open(fileName)
    lineCount, err := countLinesInFile(file)
    if err != nil {
      fmt.Println("Error counting lines:")
      os.Exit(1)
    }
    
    defer file.Close()
    output = fmt.Sprintf("%s %5d", output, lineCount)
  } 

  if *word {
    file, err := os.Open(fileName)
    wordCount, err := countWordsInFile(file)
    if err != nil {
      fmt.Println("Error counting words:", err)
      os.Exit(1)
    }

    defer file.Close()
    output = fmt.Sprintf("%s %6d", output, wordCount)
  }

  if *char {
    file, err := os.Open(fileName)
    charCount, err := countRunesInFile(file)
    if err != nil {
      fmt.Println("Error counting chars")
      os.Exit(1)
    }

    defer file.Close()
    output = fmt.Sprintf("%s %6d", output, charCount)
  }

  if *byte {
    file, err := os.Open(fileName)
    byteCount, err := countBytesInFile(file)
    if err != nil {
      fmt.Println("Error counting bytes:", err)
      os.Exit(1)
    }

    defer file.Close()
    output = fmt.Sprintf("%s %6d", output, byteCount)
  } 


  fmt.Println(output, fileName)
}

