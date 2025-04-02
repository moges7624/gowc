package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
	"unicode"
)

type Options struct {
  printBytes bool
  printLines bool
  printWords bool
  printChars bool
}

type stats struct {
  bytes uint64
  words uint64
  lines uint64
  chars uint64
  fileName string
}

func CalculateStatsForFiles(fileNames []string, options Options){
  totals := stats{fileName: "total"}

  for _, filename := range fileNames {
    calculateStatsForFile(filename, options, &totals)
  }

  if len(fileNames) > 1 {
    fmt.Println(formatStats(options, totals, totals.fileName))
  }
}

func calculateStatsForFile(filename string,options Options, total *stats){
  file, err := os.Open(filename)

  if err != nil {
    log.Fatal(err)
  }

  defer file.Close()

  reader := bufio.NewReader(file)
  CalculateStatsWithTotals(reader, filename, options, total)

}

func CalculateStatsWithTotals(reader *bufio.Reader, filename string, options Options, total *stats){
  filestats := calculateStats(reader)
  filestats.fileName = filename

  fmt.Println(formatStats(options, filestats, filestats.fileName)) 

  total.lines += filestats.lines
  total.words += filestats.words
  total.bytes += filestats.bytes
}

func calculateStats(reader *bufio.Reader) stats {
  var prevChar rune
  var byteCount uint64
  var lineCount uint64
  var wordCount uint64
  var charCout uint64

  for {
    charRead, bytesRead, err := reader.ReadRune()

    if err != nil {
      if err == io.EOF {
        if prevChar != rune(0) && !unicode.IsSpace(prevChar){
          wordCount++
        }
        break
      }
      log.Fatal(err)
    }

    byteCount += uint64(bytesRead)
    charCout++

    if charRead == '\n' {
      lineCount++
    }

    if !unicode.IsSpace(prevChar) && unicode.IsSpace(charRead){
      wordCount++
    }

    prevChar = charRead
  }

  return stats{bytes: byteCount, words: wordCount, lines: lineCount, chars: charCout}
}

func formatStats(commandLineOptions Options, fileStats stats, filename string) string {
  var cols []string

  maxDigits := maxStatSize(fileStats)
  fmtString := fmt.Sprintf("%%%dd", maxDigits)

  if commandLineOptions.printLines {
    cols = append(cols, fmt.Sprintf(fmtString, fileStats.lines))
  }

  if commandLineOptions.printWords {
    cols = append(cols, fmt.Sprintf(fmtString, fileStats.words))
  }

  if commandLineOptions.printBytes {
    cols = append(cols, fmt.Sprintf(fmtString, fileStats.bytes))
  }

  if commandLineOptions.printChars {
    cols = append(cols, fmt.Sprintf(fmtString, fileStats.chars))
  }

  statsString := strings.Join(cols, " ") + " " + filename
  
  return statsString
}

func maxStatSize(fileStats stats) int {
  maxLen := 0

  lenLines := len(strconv.FormatUint(fileStats.lines, 10))
  if lenLines > maxLen {
    maxLen = lenLines
  }


  lenWords := len(strconv.FormatUint(fileStats.words, 10))
  if lenWords > maxLen {
    maxLen = lenWords
  }
  
  lenBytes:= len(strconv.FormatUint(fileStats.bytes, 10))
  if lenBytes> maxLen {
    maxLen = lenBytes
  }

  lenChars := len(strconv.FormatUint(fileStats.chars, 10))
  if lenChars > maxLen {
    maxLen = lenChars
  }

  return maxLen
}
