# Лекция 1. Введение
# Интерфейс коммандной строки
````
import (
	"fmt"
	"os"
)
func main() {
// 	var s, sep string
// 	for i := 1; i < len(os.Args); i++ {
// 		s += sep + os.Args[i]
// 		sep = " "
// 	}
// 	fmt.Println(s)

	s, sep := "", ""
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
}
//  go run main.go "privet"
````

посчитать число вхождения слов в Stdin (uniq):
````
import (
	"bufio"
	"fmt"
	"os"
)
func main() {
	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		counts[input.Text()]++
	}
	// NOTE: ignoring potential errors from input.Err()
	for line, n := range counts {
		fmt.Printf("%d\t%s\n", n, line)
	}
}
// go run main.go "privet hell privet gt privel gt d"
// "privet hell privet gt privel gt d"
// "go"
// "hyt"
// "hyt"
// Ctrl+D
````
(urlfetch) Принимает в качестве аргументов список урлов, делает запросы к этим урлам и печатает в стдаут то, что она скачала
````
import (
	"net/http"
	"fmt"
	"os"
	"io/ioutil"
)
func main() {
	for _, url := range os.Args[1:] {
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch:%v\n", err)
			os.Exit(1)
		}
		b, err := ioutil.ReadAll(resp.Body)
		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %s: %v\n", url, err)
			os.Exit(1)
		}
		fmt.Printf("%s\n", b)
	}
}
// go run main.go "https://google.com" "https://ya.ru"
````
(fetchall) то же самое, но с поддержкой многопоточности
````
import (
	"fmt"
	"time"
	"os"
	"io/ioutil"
	"net/http"
	"io"
)
func fetch(url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err) // send error to channel ch
	}
	defer resp.Body.Close() // don't leak resourcess
	nbytes, err := io.Copy(ioutil.Discard, resp.Body)
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v\n", url, err)
		return
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs %7d %s", secs, nbytes, url)
}
func main() {
	start := time.Now()
	ch := make(chan string)
	for _, url := range os.Args[1:] {
		go fetch(url, ch) // start a gorutine
	}
	for range os.Args[1:] {
		fmt.Println(<-ch) // receive from a channel ch
	}
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}
// go run main.go "https://google.com" "https://ya.ru"
````
webserver !!!!

handler echoes the Path component of the request URL r. (смысл этой функции: когда в ваш сервер прилетит запрос - будет вызвана эта функция. И эта функция посмотрит на запрос, который записан в переменной Request, и должна сформировать ответ и записать этот ответ в ResponseWriter). ResponseWriter можно использовать как Stdout: в него можно помещать ответ. Здесь в него просто записывается путь, по которому был сделан запрос
````
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}
````
````
import (
	"fmt"
	"log"
	"net/http"
)
func main() {
	http.HandleFunc("/", handler) // each request calls handler  // регистрируется handlrer. "/" - это путь (в данном случае - любой путь)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))	// стартуем сервер(адрес, nil = с дефолтными значениями)
															// эта функция вызывается навечно (если сервер удачно стартанул)
															// в случае ошибки - log.Fatal (ошибка пишется в лог и сервер не запускается)
}
// go run main.go
// В браузере: http://localhost:8000/1578/who are you
// Ctrl+C
````
(counter) - webserver, у которого есть состояние, считает число запросов, которые ему задали
````
import (
	"sync"
	"fmt"
	"log"
	"net/http"
)
var (
	counter int
	mu sync.Mutex
)
func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}
func handler(w http.ResponseWriter, r *http.Request) {
	var i int
	mu.Lock()	// так как сервер многопоточный, задо замьютить
	i = counter
	counter++
	mu.Unlock()	// размьютить
	fmt.Fprintf(w, "counter = %d\n", i)
}
// go run main.go
// В браузере: http://localhost:8000/
// В браузере: http://localhost:8000/
// Ctrl+C
````
