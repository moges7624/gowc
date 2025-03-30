package main

import (
	"bufio"
	"fmt"
	"io"
	"strings"
	"unicode/utf8"
)

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
  fmt.Printf("Hello there")
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
