package main

import (
  "bufio"
  "flag"
  "fmt"
  "io"
  "os"
  "strings"
  "unicode/utf8"
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

    fmt.Printf("%d %d %d %s\n", lineCount, wordCount, byteCount, fileName)
    defer file.Close()
  }
}


func countBytesInFile(file io.Reader) (int, error) {
  buf := make([]byte, 4096)
  totalBytes := 0

  for {
    bytesRead, err := file.Read(buf)

    if err != nil {
      if err == io.EOF {
        break
      }
      return 0, err
    }

    totalBytes += bytesRead
  }
  return totalBytes, nil
}

func countLinesInFile(file io.Reader) (int, error){
  scanner := bufio.NewScanner(file)
  lineCount := 0

  for scanner.Scan() {
    lineCount++
  }

  if err := scanner.Err(); err != nil {
    return 0, err
  }

  return lineCount, nil
}

func countWordsInFile(file io.Reader) (int, error){
  scanner := bufio.NewScanner(file)
  wordCount := 0

  for scanner.Scan() {
    line := scanner.Text()
    words := strings.Fields(line)
    wordCount += len(words)
  }

  if err := scanner.Err(); err != nil {
    return 0, err
  }

  return wordCount, nil
}

func countRunesInFile(file io.Reader)(int, error){
  scanner := bufio.NewScanner(file)
  runeCount := 0

  for scanner.Scan() {
    line := scanner.Bytes()

    for len(line) > 0 {
      _, size := utf8.DecodeRune(line)
      runeCount++
      line = line[size:]
    }
    runeCount++
  }

  if err := scanner.Err(); err != nil {
    return 0, err
  }

  return runeCount, nil
}
