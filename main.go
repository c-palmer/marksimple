package main

import (
  "fmt"
  "flag"
)

func main() {
  markdownDir := flag.String("markdown-dir", ".",
  "Path to directory containing Markdown files to be converted into HTML")

  flag.Parse()

  fmt.Println("markdown-dir:", *markdownDir)
}
