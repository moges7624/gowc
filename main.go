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

  if *byte {
    fmt.Println("Getting number of bytes")

    byteCount, err := countBytesInFile(file)
    if err != nil {
      fmt.Println("Error counting bytes:", err)
      os.Exit(1)
    }

    fmt.Println(byteCount)
  } else if *line {
    fmt.Println("Getting number of lines")

    lineCount, err := countLinesInFile(file)
    if err != nil {
      fmt.Println("Error counting lines:")
      os.Exit(1)
    }

    fmt.Println(lineCount)
  } else if *word {
    fmt.Println("Getting number of words")

    wordCount, err := countWordsInFile(file)
    if err != nil {
      fmt.Println("Error counting words:", err)
      os.Exit(1)
    }

    fmt.Println(wordCount)
  } else if *char {
    fmt.Println("Getting number of chars")

    charCount, err := countRunesInFile(file)
    if err != nil {
      fmt.Println("Error counting chars")
      os.Exit(1)
    }

    fmt.Println(charCount)
  } else {
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

    fmt.Printf("%6d %6d %6d %s\n", lineCount, wordCount, byteCount, fileName)
    defer file.Close()
  }
}


