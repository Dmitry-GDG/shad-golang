package main

// import (
// 	"fmt"
// 	"flag"
// 	"strings"
// )
// var (
// 	n = flag.Bool("n", false, "omit training newline")
// 	sep = flag.String("s", " ", "separator")
// )
// func main() {
// 	flag.Parse()
// 	fmt.Print(strings.Join(flag.Args(), *sep))
// 	if !*n {
// 		fmt.Println()
// 	}
// }
// // go run main.go "ght" "tre" "rtf"

// //
// Приведение типов
// //
// import "fmt"
// type Celsius float64
// type Farenheit float64
// func CToF(c Celsius) Farenheit { return Farenheit(c * 5 / 9 + 32)}
// func FToC(f Farenheit) Celsius { return Celsius((f - 32) * 5 / 9)}
// func main() {
// 	var c Celsius
// 	var f Farenheit
// 	c = 36.6
// 	f = CToF(c)
// 	fmt.Println(f)
// }

// !!! init - магическая функция, которая вызывается неявно (её компилятор сам вызывает)

// // // https://github.com/adonovan/gopl.io/blob/master/ch2/popcount/main.go
// // // https://www.reddit.com/r/golang/comments/4i4ja9/purpose_of_the_init_instructions_in_popcount_the/
// import "fmt"
// // pc[i] is the population count of i.
// var pc [256]byte
// func init() {
// 	for i := range pc {
// 		pc[i] = pc[i/2] + byte(i&1)
// 	}
// }
// // PopCount returns the population count (number of set bits) of x.
// func PopCount(x uint64) int {
// 	return int(pc[byte(x>>(0*8))] +
// 		pc[byte(x>>(1*8))] +
// 		pc[byte(x>>(2*8))] +
// 		pc[byte(x>>(3*8))] +
// 		pc[byte(x>>(4*8))] +
// 		pc[byte(x>>(5*8))] +
// 		pc[byte(x>>(6*8))] +
// 		pc[byte(x>>(7*8))])
// }
// func main() {
// 	var inpt uint64
// 	for _, err := fmt.Scan(&inpt); err == nil; _, err = fmt.Scan(&inpt){
// 		fmt.Println(PopCount(inpt))
// 	}
// 	// fmt.Println(PopCount(10))
// }

// // 
// Scope - уровень, на котором определяются имена
// // x и x в цикле for - разные!!!!!
// import "fmt"
// func main() {
// 	x := "hello"
// 	for _, x := range x {
// 		x := x + 'A' - 'a'
// 		fmt.Printf("%c", x) // HELLO
// 	}
// 	fmt.Println()
// }

// // 
// Strings
// // i - обращение к конкретному байту (не символу!)
// import (
// 	"fmt"
// 	"unicode/utf8"
// )
// func main() {
// 	str := "hello!"
// 	str1 := "Привет"
// 	str1r := []rune(str1)
// 	fmt.Printf("длина слова на латинице:\t%d\nдлина слова на кирилице:\t%d\nсколько рун в слове на кирилице:\t%d\nдлина после перевода в руны:\t%d\n", 
// 	len(str), len(str1), 
// 	utf8.RuneCountInString(str1), len(str1r))
// }
// // длина слова на латинице:        6
// // длина слова на кирилице:        12
// // сколько рун в слове на кирилице:        6
// // длина после перевода в руны:    6

// //
// unicode
// //
// type rune = int32

// //
// кодировка utf8
// //
// 0xxxxxx								runes 0-127
// 11xxxxx 10xxxxxx						128-2047
// 110xxxx 10xxxxxx 10xxxxxx			2048-65535
// 1110xxx 10xxxxxx 10xxxxxx 10xxxxxx	65536-0x10fff

// //
// Разница между рунами и байтами
// //
import "unicode/utf8"
func count Runes() {
	s := "Привет"
	fmt.Println(len(s)) // "13"
	fmt.Println(utf8.RuneCountInString(s)) // "9"
}

// // 
// utf8
// //
// import (
// 	"fmt"
// 	"unicode/utf8"
// )
// func main() {
// 	s := "Привет"
// 	for i := 0; i < len(s); {
// 		r, size := utf8.DecodeRuneInString(s[i:])
// 		fmt.Printf("%d\t%c\n", i, r)
// 		i += size
// 	}
// }
// // 0       П
// // 2       р
// // 4       и
// // 6       в
// // 8       е
// // 10      т
//
//// Декодирование utf8 встроено в язык:
// import (
// 	"fmt"
// )
// func main() {
// 	s := "Привет"
// 	for i, r := range(s) {
// 		fmt.Printf("%d\t%q\t%d\n", i, r, r)
// 	}
// 	// runes := []rune(s)
// 	// str := string(runes) - может быть больше исходной s
// }
// // 0       'П'     1055
// // 2       'р'     1088
// // 4       'и'     1080
// // 6       'в'     1074
// // 8       'е'     1077
// // 10      'т'     1090
//
// Некорректный байт превращается в unicode replacement character `\uFFFD`

// // 
// []byte
// //
// // s := "abc"
// // b := []byte(s)
// // s2 := string(b)
// // bytes.Buffer
// // программа уменьшает время работы: 
// // сначала огромное количество интов пишутся в буфер, 
// // который возвращается
// import (
// 	"fmt"
// 	"bytes"
// )
// func intsToString(values []int) string {
// 	var buf bytes.Buffer
// 	buf.WriteByte('[')
// 	for i, v := range values {
// 		if i > 0 {
// 			buf.WriteString(", ")
// 		}
// 		fmt.Fprintf(&buf, "%d", v)
// 	}
// 	buf.WriteByte(']')
// 	return buf.String()
// }
// func main() {
// 	var d []int
// 	for i := 0; i < 1000; i++ {
// 		d = append(d, 10100101001)
// 	}
// 	fmt.Printf("%s, %T\n", intsToString(d), intsToString(d))
// }
// // go run main.go 
// // [10100101001, 10100101001, ... 10100101001, 10100101001], string

// // 
// constants
// //
// type Flags uint64const (
// 	FlugUp Flags = 1 << iota	// 1 is up
// 	FlagBroadcast				// 2 тоже тип Flags
// 	FlagLoopback				// 4
// 	FlagPointToPoint			// 8
// 	FlagMulticast				// 16
// )
// // untyped constants (написать сверху правило, 
// //	оно применится для всех констант в столбике)
// // 1024 в десятичной = 10000000000
// import "fmt"
// const (
// 	_ = 1 << (10 * iota)
// 	KiB		//колобайт, равно 10000000000
// 	MiB
// 	GiB
// 	TiB		// (exceeds 1 << 32)
// 	PiB
// 	EiB
// 	ZiB		// (exceeds 1 << 64) больше, чем uint64
// 	YiB		// несмотря на то, что не помещается в 64 разрядную систему, вычислять с ней можно
// )
// func main() {
// 	fmt.Println(YiB/ZiB) // 1024
// 	fmt.Println(YiB) // fail
// }

// // 
// arrays - обязательно указывать размер. Или троеточие, 
// // тогда размер посчитается автоматически
// // var a [3]int
// // var q [3]int = [3]int{1, 2, 3}
// import "fmt"
// func main() {
// 	h := [...]int{5: 4, 9: 1}
// 	for i, j := range h {
// 		fmt.Println(i, j)
// 	}
// }
// // 0 0
// // 1 0
// // 2 0
// // 3 0
// // 4 0
// // 5 4
// // 6 0
// // 7 0
// // 8 0
// // 9 1

// // 
// slices
// //
// // s := make([]type, len, cap)
// // s := make([]int, 10, 10)
// // s := []int{1, 2, 3}
// // reverse reverses a slice of ints in place
// import "fmt"
// func reverse(s []int) {
// 	for i, j := 0, len(s) - 1; i < j; i, j = i + 1, j - 1 {
// 		s[i], s[j] = s[j], s[i]
// 	}
// }
// func main() {
// 	s := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
// 	fmt.Println(s)
// 	reverse(s)
// 	fmt.Println(s)
// }
// // [1 2 3 4 5 6 7 8 9]
// // [9 8 7 6 5 4 3 2 1]

// // сравнение слайсов
// import "fmt"
// func main() {
// 	var a, b []string	//создали два nil слайса
// 	fmt.Println(a == b)	//true
// 	var c []int			//создали nil слайс
// 	d := []int{}		// слайс на массив нулевого размера
// 	fmt.Println(c == d)	//false (c == nil && d != nil) (len(c) == 0 && len(d) == 0)
// }

// // append, срезы и переаллокация слайсов
// // если надо к одному слайсу добавить другой слайс, то троеточие:
// // var a, b []int
// // a = append(a, b...)
// import "fmt"
// func main() {
// 	s := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
// 	fmt.Println(s)
// 	s = s[2:4]
// 	fmt.Println(s)
// 	s = s[0:4]
// 	fmt.Println(s)
// 	s = s[:2]
// 	fmt.Println(s)
// 	s = append(s, 1, 2, 4, 5, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1)
// 	fmt.Println(s)
// }
// // [1 2 3 4 5 6 7 8 9]
// // [3 4]
// // [3 4 5 6]
// // [3 4]
// // [3 4 1 2 4 5 1 1 1 1 1 1 1 1 1 1]

// // если надо создать копию слайса, то создаётся nil слайс
// // и к нему аппендится необходимый слайс
// import "fmt"
// func main() {
// 	var a, b []int	//создали два nil слайса
// 	s := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
// 	fmt.Printf("s: len: %d, cap: %d,\tslice: %v\n", len(s), cap(s), s)
// 	a = append(b, s...)
// 	fmt.Printf("a: len: %d, cap: %d, \tslice: %v\n", len(a), cap(a), a)
// 	a[5] = 0
// 	fmt.Printf("s: len: %d, cap: %d,\tslice: %v\n", len(s), cap(s), s)
// 	fmt.Printf("a: len: %d, cap: %d,\tslice: %v\n", len(a), cap(a), a)
// }
// // s: len: 9, cap: 9,      slice: [1 2 3 4 5 6 7 8 9]
// // a: len: 9, cap: 10,     slice: [1 2 3 4 5 6 7 8 9]
// // s: len: 9, cap: 9,      slice: [1 2 3 4 5 6 7 8 9]
// // a: len: 9, cap: 10,     slice: [1 2 3 4 5 0 7 8 9]

// // remove - удаление элемента в слайсе
// import "fmt"
// func remove(slice []int, i int) []int {
// 	copy(slice[i:], slice[i + 1:])
// 	return slice[:len(slice) - 1]
// }
// func main() {
// 	s := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
// 	fmt.Printf("s: len: %d, cap: %d, slice: %v\n", len(s), cap(s), s)
// 	remove(s, 5)
// 	fmt.Printf("s: len: %d, cap: %d, slice: %v\n", len(s), cap(s), s)
// 	a := remove(s, 5)
// 	fmt.Printf("s: len: %d, cap: %d, slice: %v\n", len(s), cap(s), s)
// 	fmt.Printf("a: len: %d, cap: %d, slice: %v\n", len(a), cap(a), a)
// 	a[5] = 0
// 	fmt.Printf("s: len: %d, cap: %d, slice: %v\n", len(s), cap(s), s)
// 	fmt.Printf("a: len: %d, cap: %d, slice: %v\n", len(a), cap(a), a)
// }
// // s: len: 9, cap: 9, slice: [1 2 3 4 5 6 7 8 9]
// // s: len: 9, cap: 9, slice: [1 2 3 4 5 7 8 9 9]
// // s: len: 9, cap: 9, slice: [1 2 3 4 5 8 9 9 9]
// // a: len: 8, cap: 9, slice: [1 2 3 4 5 8 9 9]
// // s: len: 9, cap: 9, slice: [1 2 3 4 5 0 9 9 9]
// // a: len: 8, cap: 9, slice: [1 2 3 4 5 0 9 9]

// // использовать слайс как стек:
// // stack = append(stack, v) // push v
// // top := stack[len(stack) - 1] // top of stack
// // stack = stack[:len(stack) - 1] // pop

// // 
// maps
// // ages := make(map[string]int)
// // ages := make(map[string]int) {
// // 	"alice": 31,
// // 	"bob": 34,
// // }
// // удаление элемента
// // delete(ages, "alice")
// // запрещено брать указатель элемента: нельзя &ages["bob"]
// проверить, есть ли такой элемент:
// age, ok := ages["bob"]
// if !ok {/*...*/} //если такого ключа нет, то ...
// перебор значенийЖ
// for name, age := range ages {/*...*/}
// //Ключом можно делать всё, что можно сравнивать: структуры, указатели, интерфейсы
// //Слайсы нельзя делать ключом

// // 
// set
// // var s0 map[string]bool // так хуже, лучше как ниже:
// // var s1 map[string]struct{}

// // 
// struct
// // type Point struct{ X, Y int }
// // p := Point{1, 2}
// //p := Point{X: 1, Y: 2}
// import "fmt"
// type Point struct{ X, Y int }
// func main() {
// 	p1 := Point{1, 2}
// 	p2 := Point{1, 2}
// 	p3 := Point{2, 1}
// 	fmt.Println(p1 == p2) //true
// 	fmt.Println(p1 == p3) //false
// }
// //
// // Структуру можно использовать как ключ в map
// // struct embedding:
// // type Point struct {
// // X, Y int
// // }
// // type Circle struct {
// // 	Point
// // 	Radius int
// // }
// // c := Circle{
// // 	Point: Point{X: 10, Y: 10},
// // 	Radius: 1,
// // }
// // c.X = 0 // == c.Point.X = 0
// // 
// Синтаксис структуры позволет вписывать после поля 
// произвольный строковой литерал, который можно использовать 
// для передачи опций для функции конвертации форматов, серизации/десеризации
// Например, требования к порлям могут отличаться, и можно указать имя, 
// которое декодер Jsona будет использовать
// json
// type Movie struct {
// 	Title string
// 	Year int `json:"year"`
// 	Color bool `json:"color,omitempty"`
// 	Actors []string
// }
// //
// marshal
// data, err := json.Marshal(movies)
// if err != nil {
// 	log.Fatalf("JSON marshaling failed: %s", err)
// }
// fmt.Printf("%s\n", data)
// //
// data, err := json.MarshalIndent(movies, "", "     ")
// if err != nil {
// 	log.Fatalf("JSON marshaling failed: %s", err)
// }
// fmt.Printf("%s\n", data)
// //
// unmarshal
// var movie Movie
// if err := json.Unmarshal(data, &movie); err != nil {
// 	log.Fatalf("JSON unmarshaling failed: %s", err)
// }
// fmt.Println(movie)
// //
// github
// func SearchIssues(terms []string) (*IssueSearchResult, error) {
// 	q := url.QueryEscape(strings.Join(terms, " "))
// 	resp, err := http.Get(IssueUrl + "?q=" + q)
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer resp.Body.Close()
// 	if resp.StatusCode != http.StatusOK {
// 		return nil, fmt.Errorf("searc query failed: %s", resp.Status)
// 	}
// 	var result IssueSearchResult
// 	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
// 		return nil, err
// 	}
// 	return &result, nil
// }
// //

// //
// functions
// если возвращиемое значение ф-ции имеет имя, 
// то этот элнмент ненадо объявлять и не надо указывать в ретёрне: 
// просто указать ретёрн и вернутся те значения этих переменных, 
// которые есть у них на момент вызова ретёрна
// func CountWordAndImages(url string)(words, images int, err error) {/*...*/}

// //
// errors
// если функция возвращает ошибку, то она всегда 
// возвращает её последним аргументом
// Текст ошибки должен быть в lowercase
// Ошибку надо пояснить в соответствии с этапом программы:
// doc, err := html.Parse(resp.Body)
// if err != nil {
// 	return nil, fmt.Errorf("parsing %s as HTML: %w", url, err)
// }

// //
// EOF - нечего больше читать, конец файла
// import "errors"
// var EOF = errors/New("EOF")
// in := bufio.NewReader(os.Stdin)
// for {
// 	r, _, err := in.ReadRune()
// 	if err == io.EOF {
// 		break  // finishing reading
// 	}
// 	if err != nil {
// 		return fmt.Errorf("read failed: %v", err)
// 	}
// 	// use r
// }

// //
// function values
// // имя функции может быть передано как значение
// import "fmt"
// func Inc(i int) int { return i + 1 }
// var f func(i int) int
// func main() {
// 	// f = func(i int) int { return i * 2 }
// 	for i:= 0; i < 5; i++ {
// 		if f == nil {
// 			f = Inc
// 		}
// 		fmt.Println(f(i))
// 	}
// }

// //
// recursion
// // рекурсивная ф-ция объявляется отдельно 
// от определения (var visit func(n *Node) ), 
// иначе к ней невозможно обратиться рекурсивно
// //
// type Node struct {
// 	V int
// 	L, R *Node
// }
// func PrintAll(w io.writer, root *Node) {
// 	var visit func(n *Node) // объявляется отдельно 
// 	visit = func(n *Node) {
// 		fmt.Fprintln(w, n.V)
// 		if n.L != nil {
// 			visit(n.L)
// 		}
// 		if n/R != nil {
// 			visit(n.R)
// 		}
// 	}
// 	visit(root)
// }

