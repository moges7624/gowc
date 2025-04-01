package main

import (
	"flag"
)

func main(){
  var commandLineOptions Options

  flag.BoolVar(&commandLineOptions.printLines, "l", false, "count line")
  flag.BoolVar(&commandLineOptions.printWords, "w", false, "count words")
  flag.BoolVar(&commandLineOptions.printBytes, "c", false, "count bytes")
  flag.BoolVar(&commandLineOptions.printChars, "m", false, "count chars")
  flag.Parse()

  if flag.NFlag() == 0 {
    commandLineOptions.printBytes = true
    commandLineOptions.printLines = true
    commandLineOptions.printWords = true 
  }

  fileNames := flag.CommandLine.Args()

  if len(fileNames) == 0 {
    return
  }

  CalculateStatsForFiles(fileNames, commandLineOptions)
}
