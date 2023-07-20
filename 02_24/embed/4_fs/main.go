package main

import (
	"log"
	"embed"
	"net/http"
	"text/template"
)

var (
	//go:embed resources
	res embed.FS

	pages = map[string]string{  // это маппинг из урла в темплейт
		"/": "resources/index.gohtml",
	}
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) { // на путь "/" выполняет наш код хэндлера.
		page, ok := pages[r.URL.Path] // она ищет в pages путь
		if !ok { // если не находит - 404
			w.WriteHeader(http.StatusNotFound)
			return
		}

		tpl, err := template.ParseFS(res, page) // если находит, то используем спец ф-цию, которая умеет работать с FS, ищет темплейт по пути "resources/index.gohtml" и его парсит в tpl
		if err != nil {
			log.Printf("page %v not found in pages cache...", r.RequestURI)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "text/html") // заполняем Content-Type
		w.WriteHeader(http.StatusOK) // заполняем StatusOK
		data := map[string]interface{}{ // создаём мапу
			"UserAgent": r.UserAgent(),
		}
		if err := tpl.Execute(w, data); err != nil {  // выполняем tpl.Execute на нашей мапе
			return
		}
	})

	http.FileServer(http.FS(res))
	log.Println("...server started...")

	err := http.ListenAndServe(":8088", nil)
	if err != nil {
		// log.Fatal(err)
	}
}

// go run main.go 
// in browser: localhost:8088
// in terminal: curl -i localhost:8088
// Output:
// HTTP/1.1 200 OK
// Content-Type: text/html
// Date: Thu, 20 Jul 2023 14:39:29 GMT
// Content-Length: 196

// <html lang="en">
// <head>
// <meta charset="UTF-8"/>
// <title>go:embed demo</title>
// </head>
// <body>
// <div>
// <h1>Hello, curl/7.81.0! </h1>
// <p>If you see this, then go:embed worked!</p>
// </div>
// </body>
// </html>

