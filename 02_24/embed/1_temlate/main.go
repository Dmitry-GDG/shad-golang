package main
import (
	"log"
	"os"
	// "text/template"
	"html/template"
)
const index = `<html lang="en">
<head>
<meta charset="UTF-8"/>
<title>go:embed demo</title>
</head>
<body>
<div>
<h1>Hello, {{ .UserAgent }}! </h1>
<p>If you see this, then go:embed worked!</p>
</div>
</body>
</html>`
func main() {
	t, err := template.New("index").Parse(index) // из пакета html/template
	if err != nil {
		log.Fatal(err)
	}

	err = t.Execute(os.Stdout, map[string]interface{}{
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


// // Помимо мапы сюда можно передать просто структуру, в которой определено поле UserAgent

// package main
// import (
// 	"log"
// 	"os"
// 	// "text/template"
// 	"html/template"
// )
// const index = `<html lang="en">
// <head>
// <meta charset="UTF-8"/>
// <title>go:embed demo</title>
// </head>
// <body>
// <div>
// <h1>Hello, {{ .UserAgent }}! </h1>
// <p>If you see this, then go:embed worked!</p>
// </div>
// </body>
// </html>`
// func main() {
// 	t, err := template.New("index").Parse(index) // из пакета html/template
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	// err = t.Execute(os.Stdout, map[string]interface{}{
// 	// 	"UserAgent": "Mozilla/5.0",
// 	// })
// 	// if err != nil {
// 	// 	log.Fatal(err)
// 	// }
// 	testStruct := struct {
// 		UserAgent string
// 	}{
// 		UserAgent: "Mozilla",
// 	}
// 	err = t.Execute(os.Stdout, &testStruct)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// }

// // go run main.go 
// // Output:
// // <html lang="en">
// // <head>
// // <meta charset="UTF-8"/>
// // <title>go:embed demo</title>
// // </head>
// // <body>
// // <div>
// // <h1>Hello, Mozilla! </h1>
// // <p>If you see this, then go:embed worked!</p>
// // </div>
// // </body>
// // </html>