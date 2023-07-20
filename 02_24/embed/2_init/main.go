package main
import (
	"log"
	"os"
	"html/template"
)

var indexTemplate *template.Template

func init() {
	indexTemplate = template.Must(template.ParseFiles("resources/index.gohtml"))
}

func main() {

	err := indexTemplate.Execute(os.Stdout, map[string]interface{}{
		"UserAgent": "Mozilla/5.0",
	})
	if err != nil {
		log.Fatal(err)
	}
}

// go run main.go 
// Output:
// <html lang="en">
// <head>
// <meta charset="UTF-8"/>
// <title>go:embed demo</title>
// </head>
// <body>
// <div>
// <h1>Hello, Mozilla/5.0! </h1>
// <p>If you see this, then go:embed worked!</p>
// </div>
// </body>
// </html>

