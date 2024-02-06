package main

import (
  "fmt"
  "flag"
  "os"
)

func main() {
  var markdownPath string
  flag.StringVar(&markdownPath, "markdown-path", ".",
`Path to file or directory containing Markdown to be converted into HTML.
Defaults to current directory.`)

  flag.Parse()

  fileInfo, err := os.Stat(markdownPath);
  if  err != nil {
    fmt.Fprintf(os.Stderr, "Error: %s\n", err)
    os.Exit(1)
  }

  if !fileInfo.IsDir() {
    fmt.Println("Path is a file", markdownPath)
    return
  }

  fmt.Println("Path is a directory", markdownPath)
}
