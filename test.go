package main

import (
  // "fmt"
  "log"
  "github.com/gomarkdown/markdown"
  // "github.com/gomarkdown/markdown/ast"
  "github.com/gomarkdown/markdown/html"
  "github.com/gomarkdown/markdown/parser"
  "text/template"
  "os"
)

func main() {
  var md = []byte(
`
# header

Sample text.

$$
x = {-b \pm \sqrt{b^2-4ac} \over 2a}
$$

Hey, this is some $e=mc^2$ inline math

[link](http://example.com)
`)

  generatedHTML := string(mdToHTML(md))

//   var html =
// `
// <html>
//     <head>
//       <title>MD to HTML Example</title>
//       <script src="https://polyfill.io/v3/polyfill.min.js?features=es6"></script>
//       <script id="MathJax-script" async src="https://cdn.jsdelivr.net/npm/mathjax@3/es5/tex-mml-chtml.js"></script>
//     </head>
//     <body>
//       {{ .generatedHTML }}
//     </body>
// </html>
// `
//
//   templ := template.Must(template.New("html").Parse(html))
//   templ.Execute(os.Stdout, recipients)

  // fmt.Println(string(generatedHTML))

// Define a template.
	const letter = `
<html>
    <head>
      <title>MD to HTML Example</title>
      <script src="https://polyfill.io/v3/polyfill.min.js?features=es6"></script>
      <script id="MathJax-script" async src="https://cdn.jsdelivr.net/npm/mathjax@3/es5/tex-mml-chtml.js"></script>
    </head>
    <body>
      {{.Name}}

      WOW!!!!
    </body>
</html>
`

	// Prepare some data to insert into the template.
	type Recipient struct {
		Name string
	}
	var recipients = []Recipient{
		{generatedHTML},
	}

	// Create a new template and parse the letter into it.
	t := template.Must(template.New("letter").Parse(letter))

	// Execute the template for each recipient.
	for _, r := range recipients {
		err := t.Execute(os.Stdout, r)
		if err != nil {
			log.Println("executing template:", err)
		}
	}

}

func mdToHTML(md []byte) []byte {
  // create markdown parser with extensions
  extensions := parser.CommonExtensions | parser.MathJax
  p := parser.NewWithExtensions(extensions)
  doc := p.Parse(md)

  // create HTML renderer with extensions
  htmlFlags := html.CommonFlags | html.HrefTargetBlank
  opts := html.RendererOptions{Flags: htmlFlags}
  renderer := html.NewRenderer(opts)

  return markdown.Render(doc, renderer)
}
