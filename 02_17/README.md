# Лекция 2. Базовые конструкции языка
25 ключевых слов Golang:
````
break        default      func         interface    select
case         defer        go           map          struct
chan         else         goto         package      switch
const        fallthrough  if           range        type
continue     for          import       return       var
````
### flag
Пакет, чтоб парсить аргументы командной строки.

Чтобы зарегистрировать новый аргумент в пакет flag, необходимо сделать инициализацию нового флага с параметрами, как указано ниже. Когда мы вызываем flag.Parse(), у Вас в тех переменных, которые Вы просили через флаги зарегистрировать, будут находиться распаршенные значения этих аргументов:
````
import (
	"fmt"
	"flag"
	"strings"
)
var (
	n = flag.Bool("n", false, "omit training newline")
	sep = flag.String("s", " ", "separator")
)
func main() {
	flag.Parse()
	fmt.Print(strings.Join(flag.Args(), *sep))
	if !*n {
		fmt.Println()
	}
}
// go run main.go "ght" "tre" "rtf"
````
### Приведение типов
````
import "fmt"
type Celsius float64
type Farenheit float64
func CToF(c Celsius) Farenheit { return Farenheit(c * 5 / 9 + 32)}
func FToC(f Farenheit) Celsius { return Celsius((f - 32) * 5 / 9)}
func main() {
	var c Celsius
	var f Farenheit
	c = 36.6
	f = CToF(c)
	fmt.Println(f)
}
````
!!! init - магическая функция, которая вызывается неявно (её компилятор сам вызывает)

https://github.com/adonovan/gopl.io/blob/master/ch2/popcount/main.go

https://www.reddit.com/r/golang/comments/4i4ja9/purpose_of_the_init_instructions_in_popcount_the/
````
import "fmt"
// pc[i] is the population count of i.
var pc [256]byte
func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}
````
PopCount returns the population count (number of set bits) of x.
````
func PopCount(x uint64) int {
	return int(pc[byte(x>>(0*8))] +
		pc[byte(x>>(1*8))] +
		pc[byte(x>>(2*8))] +
		pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))])
}
func main() {
	var inpt uint64
	for _, err := fmt.Scan(&inpt); err == nil; _, err = fmt.Scan(&inpt){
		fmt.Println(PopCount(inpt))
	}
	// fmt.Println(PopCount(10))
}
````
Scope - уровень, на котором определяются имена. 

x и x в цикле for - разные!!!!!
````
import "fmt"
func main() {
	x := "hello"
	for _, x := range x {
		x := x + 'A' - 'a'
		fmt.Printf("%c", x) // HELLO
	}
	fmt.Println()
}
````
### Strings
Строки в Go объявляются с типом string:
````
var s string = "hello"

// сокращенная запись
s := "hey"
````
Практически всегда для строк используются двойные кавычки. Однако они не подходят, когда нужно написать несколько строк. Для этого используются обратные кавычки:
````
q := `
    SELECT *
    FROM person
    WHERE age > 18
`
````
Строки можно сравнивать операторами: ==, >, <, <=, >=, где строки сравниваются посимвольно в лексическом порядке и по длине. Это свойство часто используется при сортировке массива строк:
````
"привет" == "привет" // true
"golang" > "go" // true
"golang" > "lang" // false
"go" > "foobar" // true
````
Базовые операции со строками в любом языке — это конкатенация и интерполяция. Конкатенация осуществляется с помощью знака +:
````
"hello " + "world" // "hello world"
````
В Go нет привычной интерполяции, как в динамических языках. Она реализуется через форматирующую функцию fmt.Sprintf:
````
username := "Ivan"

greetings := fmt.Sprintf("hello, %s", username) // "hello, Ivan"
````
Узнать длину строки можно с помощью встроенной функции len:
````
len("go") // 2

// будьте внимательны! Функция считает кол-во байт, а не кол-во символов
len("го") // 4
````
Задание

Для работы со строками часто используется стандартная библиотека strings. В данном задании вам понадобятся следующие функции:
````
// обрезает символы, переданные вторым аргументом, с обеих сторон строки
Trim(s, cutset string) string
// пример
strings.Trim(" hello ", " ") // "hello"

// преобразует все буквы в строке в нижний регистр
strings.ToLower(s string) string
// пример
strings.ToLower("пРиВеТ") // "привет"

// озаглавливает первую букву в каждом слове в строке
strings.Title(s string) string
// пример
strings.Title("привет, джон") // "Привет, Джон"
````
Реализуйте функцию Greetings(name string) string, которая вернет строку-приветствие. Например, при передаче имени Иван, возвращается Привет, Иван!. Учтите, что имя может быть написано в разном регистре и содержать пробелы.

Ваше решение:
````
package solution

import (
	"fmt"
	"strings"
)

// BEGIN (write your solution here)
func Greetings(name string) string {
	return fmt.Sprintf("Привет, %s!", strings.Title(strings.ToLower(strings.Trim(name, " "))))
}
````
#### Стандартный пакет strings

Для работы со строками в Go существует стандартный пакет strings, который содержит основные функции. С некоторыми мы уже встречались в первом модуле (например strings.ReplaceAll). Теперь рассмотрим список самых часто встречающихся функций:
````
import "strings"

// проверяет наличие подстроки в строке
strings.Contains("hello", "h") // true

// разбивает строку по Юникод символам или по переданному разделителю
strings.Split("hello", "") // ["h", "e", "l", "l", "o"]

// склеивает строки из слайса с разделителем
strings.Join([]string{"hello","world!"}, " ") // "hello world!"

// обрезает символы из второго аргумента в строке
strings.Trim(" hey !", " ") // "hey !"
````
Очень важная часть пакета strings — это Builder. Когда необходимо собрать большую строку по каким-то правилам, использование конкатенации — не лучшее решение, потому что каждая операция создает новую строку, что сильно влияет на производительность при большом количестве операций. Такая задача решается с помощью билдера:
````
import "strings"

sb := &strings.Builder{}

sb.WriteString("hello")
sb.WriteString(" ")
sb.WriteString("world")

sb.String() // "hello world"
````
Задание

В пакете unicode есть функция unicode.Is(unicode.Latin, rune), которая проверяет, что руна является латинским символом или нет.

Реализуйте функцию latinLetters(s string) string, которая возвращает только латинские символы из строки s. Например:
````
latinLetters("привет world!") // "world"
````
Решение учителя:
````
package solution

import (
	"strings"
	"unicode"
)

// BEGIN (write your solution here)
func latinLetters(s string) string {
	sb := &strings.Builder{}

	for _, r := range s {
		if unicode.Is(unicode.Latin, r)  {
			sb.WriteRune(r)
		}
    }

    return sb.String()
}
// END
````
Ваше решение:
````
package solution

import (
	"strings"
	"unicode"
)

// BEGIN (write your solution here) (write your solution here)
func latinLetters(s string) string {
	sb := &strings.Builder{}
	for _, ch := range(s){
		if unicode.Is(unicode.Latin, ch) {
			var tmp []byte
			tmp = append(tmp, byte(ch))
			sb.Write(tmp)
		}
	}
	return sb.String()
}
````

i - обращение к конкретному байту (не символу!)
````
import (
	"fmt"
	"unicode/utf8"
)
func main() {
	str := "hello!"
	str1 := "Привет"
	str1r := []rune(str1)
	fmt.Printf("длина слова на латинице:\t%d\nдлина слова на кирилице:\t%d\nсколько рун в слове на кирилице:\t%d\nдлина после перевода в руны:\t%d\n", 
	len(str), len(str1), 
	utf8.RuneCountInString(str1), len(str1r))
}
// длина слова на латинице:        6
// длина слова на кирилице:        12
// сколько рун в слове на кирилице:        6
// длина после перевода в руны:    6
````
### Строки и байты

Строки в Go — это иммутабельные массивы байт. Для стандартного компилятора Go внутренняя структура строки описана как:
````
type _string struct {
    elements *byte // байты
    len      int   // кол-во байт
}
````
После инициализации строку нельзя изменить и такая иммутабельность позволяет избежать побочных эффектов в коде.
````
s := "hello"
s[4] = "" // ошибка компиляции: cannot assign to s[4] (strings are immutable)
````
Стоит отметить, что тип данных byte — это алиас к типу uint8 (0-255). Во-первых, потому что нужно абстрактно отличать типы в коде. Во-вторых, байты представляют ASCII символы, а в кодовой таблице ASCII символов 256 кодов:
````
package main

import "fmt"

func main() {
    s := "hey"

    fmt.Println(s[0], s[1], s[2]) // 104 101 121

    fmt.Println(string(s[0]), string(s[1]), string(s[2])) // h e y
}
````
Большинство библиотечных функций работают со слайсами байт []byte для производительности. Конвертация строки в слайс байт описывается в коде явно:
````
package main

import "fmt"

func main() {
    s := "hey"
    bs := []byte(s)

    fmt.Println([]byte(s)) // [104 101 121]

    fmt.Println(string(bs)) // hey
}
````
Отдельные ASCII символы можно объявлять сразу с типом byte. Для этого нужно обернуть символ в одинарные кавычки и указать тип byte:
````
package main

import (
    "fmt"
    "reflect"
)

func main() {
    asciiCh := byte('Z')
    asciiChStr := string(asciiCh)

    fmt.Println(reflect.TypeOf(asciiCh), asciiCh) // uint8 90

    fmt.Println(reflect.TypeOf(asciiChStr), asciiChStr) // string Z
}
````
Задание

Реализуйте функции nextASCII(b byte) byte и prevASCII(b byte) byte, которые возвращают следующий или предыдущий символ ASCII таблицы соответственно. Например:
````
nextASCII(byte('a')) // 'b'
prevASCII(byte('b')) // 'a'
````
Допускаем, что в функцию prevASCII не может прийти нулевой символ, а в функцию nextASCII — последний символ ASCII таблицы.

Решение учителя:
````
package solution

// BEGIN

// nextASCII returns the next byte after b according to the ASCII code table.
func nextASCII(b byte) byte {
	return b + 1
}

// prevASCII returns the previous byte after b according to the ASCII code table.
func prevASCII(b byte) byte {
	return b - 1
}

// END
````
Ваше решение:
````
package solution

// BEGIN (write your solution here)
func nextASCII(b byte) byte {
	if b < 255 && b >= 0 {
		b += 1
	}
	return b
}

func prevASCII(b byte) byte {
	if b > 0 && b <= 255 {
		b -= 1
	}
	return b
}
````
#### Обход строки

Так как строка — это массив байт, ее можно обойти с помощью цикла for:
````
package main

import (
    "fmt"
)

func main() {
    s := "hello"
    for i := 0; i < len(s); i++ {
        fmt.Println(string(s[i]))
    }

}
````
Вывод:
````
h
e
l
l
o
````
Таким способом можно обойти только строки, состоящие из ASCII символов. Если строка содержит мультибайтовые символы, вывод будет некорректен:
````
package main

import (
    "fmt"
)

func main() {
    s := "привет"
    for i := 0; i < len(s); i++ {
        fmt.Println(string(s[i]))
    }

}
````
Вывод проверьте сами в Go Playground https://play.golang.org/p/-G3ygH0rTIv

В Go присутствует синтаксический сахар при обходе строки. Если использовать конструкцию for range, строка автоматически будет преобразована в []rune, то есть обход будет по Юникод символам:
````
package main

import (
    "fmt"
)

func main() {
    emoji := []rune("cool😀")

    for _, ch := range emoji {
        fmt.Println(ch, string(ch)) // выводим код символа и его строковое представление
    }
}
````
Вывод:
````
99 c
111 o
111 o
108 l
128512 😀
````
Задание

Реализуйте функцию shiftASCII(s string, step int) string, которая принимает на вход состоящую из ASCII символов строку s и возвращает новую строку, где каждый символ из входящей строки сдвинут вперед на число step. Например:
````
shiftASCII("abc", 0) // "abc"
shiftASCII("abc1", 1) // "bcd2"
shiftASCII("bcd2", -1) // "abc1"
shiftASCII("hi", 10) // "rs"
````
Если после сдвига код символа выходит за рамки ASCII, то число должно быть взято по модулю 256:
````
shiftASCII("abc", 256) // "abc"
shiftASCII("abc", -511) // "bcd"
````
Решение учителя:
````
package solution

// BEGIN

// shiftASCII shifts each byte in the string s and if there is an overflow of ASCII it takes a new code by module of 256.
func shiftASCII(s string, step int) string {
	if step == 0 {
		return s
	}

	shifted := make([]byte, len(s))
	for i := 0; i < len(s); i++ {
		shifted[i] = byte(int(s[i]) + step)
	}

	return string(shifted)
}

// END
````
Ваше решение:
````
package solution

// BEGIN (write your solution here)
func shiftASCII(s string, step int) string {
	var tmp int
	var outp string
	for i := 0; i < len(s); i++ {
		tmp = (int(s[i]) + step) % 256
		if tmp < 0 {
			tmp += 256
		}
		outp += string(tmp)
	}
	return outp
}
````
#### Форматирование строк

В предыдущих уроках мы использовали пакет fmt для вывода переменных или результатов функций:
````
s := "hello world"

// печатает вывод на следующей строке
fmt.Println(s) // "hello world"
````
Пакет fmt так же используется для форматирования строк. Плейсхолдеры разных типов данных в основном не отличаются от других языков:
````
name := "Andy"

// подставляем строку
fmt.Sprintf("hello %s", name) // "hello Andy"

// число
fmt.Sprintf("there are %d kittens", 10) // "there are 10 kittens"

// логический тип
fmt.Sprintf("your story is %t", true) // "your story is true"
````
Так же существуют специализированные плейсхолдеры, которые преобразуют сложные структуры:
````
package main

import (
    "fmt"
)

type Person struct {
    Name string
    Age  int
}

func main() {
    p := Person{Name: "Andy", Age: 18}

    // вывод значений структуры
    fmt.Println("simple struct:", p)

    // вывод названий полей и их значений
    fmt.Printf("detailed struct: %+v\n", p)

    // вывод названий полей и их значений в виде инициализации
    fmt.Printf("Golang struct: %#v\n", p)
}
````
Вывод:
````
simple struct: {Andy 18}
detailed struct: {Name:Andy Age:18}
Golang struct: main.Person{Name:"Andy", Age:18}
````
Полезное: Go Formatting https://pkg.go.dev/fmt#hdr-Printing

Задание

Реализуйте функцию generateSelfStory(name string, age int, money float64) string, которая генерирует предложение, подставляя переданные данные в возвращаемую строку. Например:
````
generateSelfStory("Vlad", 25, 10.00000025) // "Hello! My name is Vlad. I'm 25 y.o. And I also have $10.00 in my wallet right now."
````
Шаблон возвращаемой строки: Hello! My name is [name]. I'm [age] y.o. And I also have $[money with precision 2] in my wallet right now.

Решение учителя:
````
package solution

import "fmt"

// BEGIN

// generateSelfStory generates a self story substituting vars.
func generateSelfStory(name string, age int, money float64) string {
	return fmt.Sprintf("Hello! My name is %s. I'm %d y.o. And I also have $%.2f in my wallet right now.", name, age, money)
}

// END
````
Ваше решение:
````
package solution

import "fmt"

// BEGIN (write your solution here)
func generateSelfStory(name string, age int, money float64) string {
	return fmt.Sprintf("Hello! My name is %s. I'm %d y.o. And I also have $%.2f in my wallet right now.", name, age, money)
}
````
### unicode
#### Руны

В современном мире невозможно работать только со строками, состоящими исключительно из ASCII символов. Везде используются нестандартные знаки, языки отличные от латиницы и эмодзи. Для работы с такими Юникод символами в Go представлен тип данных rune:
````
package main

import (
    "fmt"
)

func main() {
    emoji := []rune("привет😀")

    for i := 0; i < len(emoji); i++ {
        fmt.Println(emoji[i], string(emoji[i])) // выводим код символа и его строковое представление
    }
}
````
Вывод:
````
1087 п
1088 р
1080 и
1074 в
1077 е
1090 т
128512 😀
````
rune — это алиас к int32. Как и байты, руны были созданы для отличия от встроенного типа данных. Каждая руна представляет собой код символа стандарта Юникод. Строка свободно преобразуется в []byte и []rune, но эти 2 типа данных не конвертируются между собой напрямую:
````
s := "hey😉"

rs := []rune([]byte(s)) // cannot convert ([]byte)(s) (type []byte) to type []rune

bs := []byte([]rune(s)]) // cannot convert ([]rune)(s) (type []rune) to type []byte
````
В Go присутствует синтаксический сахар при обходе строки. Если использовать конструкцию for range, строка автоматически будет преобразована в []rune, то есть обход будет по Юникод символам:
````
package main

import (
    "fmt"
)

func main() {
    emoji := []rune("cool😀")

    for _, ch := range emoji {
        fmt.Println(ch, string(ch)) // выводим код символа и его строковое представление
    }
}
````
Вывод:
````
99 c
111 o
111 o
108 l
128512 😀
````
Задание

Реализуйте функцию isASCII(s string) bool, которая возвращает true, если строка s состоит только из ASCII символов.
type rune = int32

Решение учителя:
````
package solution

import "unicode"

// BEGIN

// isASCII checks whether the string s contains only ASCII chars.
func isASCII(s string) bool {
	for _, r := range s {
		if r > unicode.MaxASCII {
			return false
		}
	}

	return true
}

// END
````
Ваше решение:
````
package solution

//import "unicode"

// BEGIN (write your solution here)
func isASCII(s string) bool {
	if len(s) == len([]rune(s)) {
		return true
	}
	return false
}
````
### кодировка utf8
````
0xxxxxx								runes 0-127
11xxxxx 10xxxxxx						128-2047
110xxxx 10xxxxxx 10xxxxxx			2048-65535
1110xxx 10xxxxxx 10xxxxxx 10xxxxxx	65536-0x10fff
````
### Разница между рунами и байтами
````
import "unicode/utf8"
func count Runes() {
	s := "Привет"
	fmt.Println(len(s)) // "13"
	fmt.Println(utf8.RuneCountInString(s)) // "9"
}
````
### utf8
````
import (
	"fmt"
	"unicode/utf8"
)
func main() {
	s := "Привет"
	for i := 0; i < len(s); {
		r, size := utf8.DecodeRuneInString(s[i:])
		fmt.Printf("%d\t%c\n", i, r)
		i += size
	}
}
// 0       П
// 2       р
// 4       и
// 6       в
// 8       е
// 10      т
````
### Декодирование utf8 встроено в язык:
````
import (
	"fmt"
)
func main() {
	s := "Привет"
	for i, r := range(s) {
		fmt.Printf("%d\t%q\t%d\n", i, r, r)
	}
	// runes := []rune(s)
	// str := string(runes) - может быть больше исходной s
}
// 0       'П'     1055
// 2       'р'     1088
// 4       'и'     1080
// 6       'в'     1074
// 8       'е'     1077
// 10      'т'     1090
````
Некорректный байт превращается в unicode replacement character `\uFFFD`
### []byte
````
s := "abc"
b := []byte(s)
s2 := string(b)
bytes.Buffer
````
````
import "fmt"
type ByteCounter int
func (c *ByteCounter) Write(p []byte) (int, error) {
	*c += ByteCounter(len(p))
	return len(p), nil
}
func main() {
	var c ByteCounter
	c.Write([]byte("hello"))
	fmt.Println(c) // "5" = len(c)
	c.Write([]byte("привет"))
	fmt.Println(c) // "17" = len(c)
}
````
Программа уменьшает время работы: сначала огромное количество интов пишутся в буфер, который возвращается:
````
import (
	"fmt"
	"bytes"
)
func intsToString(values []int) string {
	var buf bytes.Buffer
	buf.WriteByte('[')
	for i, v := range values {
		if i > 0 {
			buf.WriteString(", ")
		}
		fmt.Fprintf(&buf, "%d", v)
	}
	buf.WriteByte(']')
	return buf.String()
}
func main() {
	var d []int
	for i := 0; i < 1000; i++ {
		d = append(d, 10100101001)
	}
	fmt.Printf("%s, %T\n", intsToString(d), intsToString(d))
}
// go run main.go 
// [10100101001, 10100101001, ... 10100101001, 10100101001], string
````
### constants
````
type Flags uint64const (
	FlugUp Flags = 1 << iota	// 1 is up
	FlagBroadcast				// 2 тоже тип Flags
	FlagLoopback				// 4
	FlagPointToPoint			// 8
	FlagMulticast				// 16
)
````
#### untyped constants (написать сверху правило, оно применится для всех констант в столбике)
1024 в десятичной = 10000000000
````
import "fmt"
const (
	_ = 1 << (10 * iota)
	KiB		//колобайт, равно 10000000000
	MiB
	GiB
	TiB		// (exceeds 1 << 32)
	PiB
	EiB
	ZiB		// (exceeds 1 << 64) больше, чем uint64
	YiB		// несмотря на то, что не помещается в 64 разрядную систему, вычислять с ней можно
)
func main() {
	fmt.Println(YiB/ZiB) // 1024
	fmt.Println(YiB) // fail
}
````
### arrays
Обязательно указывать размер. Или троеточие, тогда размер посчитается автоматически
````
var a [3]int
var q [3]int = [3]int{1, 2, 3}
````
````
import "fmt"
func main() {
	h := [...]int{5: 4, 9: 1}
	for i, j := range h {
		fmt.Println(i, j)
	}
}
// 0 0
// 1 0
// 2 0
// 3 0
// 4 0
// 5 4
// 6 0
// 7 0
// 8 0
// 9 1
````
### slices
````
s := make([]type, len, cap)
s := make([]int, 10, 10)
s := []int{1, 2, 3}
````
Для обхода коллекции в Go есть "синтаксический сахар" range. Эта конструкция обходит слайс, возвращая индекс и элемент на каждом шаге:
````
names := []string{"John", "Harold", "Vince"}

// i — это индекс, name — это значение на текущем шаге цикла
for i, name := range names {
    fmt.Println("Hello ", name, " at index ", i)
}
````
Вывод:
````
Hello  John  at index  0
Hello  Harold  at index  1
Hello  Vince  at index  2
````
#### reverse reverses a slice of ints in place
````
import "fmt"
func reverse(s []int) {
	for i, j := 0, len(s) - 1; i < j; i, j = i + 1, j - 1 {
		s[i], s[j] = s[j], s[i]
	}
}
func main() {
	s := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(s)
	reverse(s)
	fmt.Println(s)
}
// [1 2 3 4 5 6 7 8 9]
// [9 8 7 6 5 4 3 2 1]
````
#### сравнение слайсов
````
import "fmt"
func main() {
	var a, b []string	//создали два nil слайса
	fmt.Println(a == b)	//true
	var c []int			//создали nil слайс
	d := []int{}		// слайс на массив нулевого размера
	fmt.Println(c == d)	//false (c == nil && d != nil) (len(c) == 0 && len(d) == 0)
}
````
#### append, срезы и переаллокация слайсов
если надо к одному слайсу добавить другой слайс, то троеточие:
````
var a, b []int
a = append(a, b...)
````
````
import "fmt"
func main() {
	s := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(s)
	s = s[2:4]
	fmt.Println(s)
	s = s[0:4]
	fmt.Println(s)
	s = s[:2]
	fmt.Println(s)
	s = append(s, 1, 2, 4, 5, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1)
	fmt.Println(s)
}
// [1 2 3 4 5 6 7 8 9]
// [3 4]
// [3 4 5 6]
// [3 4]
// [3 4 1 2 4 5 1 1 1 1 1 1 1 1 1 1]
````
если надо создать копию слайса, то создаётся nil слайс и к нему аппендится необходимый слайс
````
import "fmt"
func main() {
	var a, b []int	//создали два nil слайса
	s := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Printf("s: len: %d, cap: %d,\tslice: %v\n", len(s), cap(s), s)
	a = append(b, s...)
	fmt.Printf("a: len: %d, cap: %d, \tslice: %v\n", len(a), cap(a), a)
	a[5] = 0
	fmt.Printf("s: len: %d, cap: %d,\tslice: %v\n", len(s), cap(s), s)
	fmt.Printf("a: len: %d, cap: %d,\tslice: %v\n", len(a), cap(a), a)
}
// s: len: 9, cap: 9,      slice: [1 2 3 4 5 6 7 8 9]
// a: len: 9, cap: 10,     slice: [1 2 3 4 5 6 7 8 9]
// s: len: 9, cap: 9,      slice: [1 2 3 4 5 6 7 8 9]
// a: len: 9, cap: 10,     slice: [1 2 3 4 5 0 7 8 9]
````
#### Копирование слайсов

Допустим, в вашей функции происходят изменения элементов, но вы не хотите затронуть входной слайс. В языке есть встроенная функция func copy(dst, src []Type) int, которая копирует слайс src в слайс dst и возвращает кол-во скопированных элементов:
````
nums := []int{1,2,3,4,5}

// важно инициализировать слайс той же длины
numsCp := make([]int, len(nums))

copy(numsCp, nums)

fmt.Println(numsCp) // [1,2,3,4,5]
````
Почему мы не можем просто перезаписать слайс в другую переменную и изменять ее? Как и с функциями, при присваивании слайса к переменной, копируется только длина и вместимость, но массив передается по ссылке:
````
nums := []int{1,2,3,4,5}

numsCp := nums

// исходный слайс nums тоже будет изменен
numsCp[0] = 10

fmt.Println(nums) // [10,2,3,4,5]
````
Существует распространенная ошибка, когда пытаются скопировать слайсы различной длины. В этом случае элементы, выходящие за рамки слайса dst, не будут скопированы:
````
nums := []int{1, 2, 3, 4, 5}

// создали слайс с длиной 0
numsCp := make([]int, 0)

// при копировании в пустой слайс ничего не произойдет
copy(numsCp, nums)

fmt.Println(numsCp) // []
````
Задание

Реализуйте функцию IntsCopy(src []int, maxLen int) []int, которая создает копию слайса src с длиной maxLen. Если maxLen равен нулю или отрицательный, то функция возвращает пустой слайс []int{}. Если maxLen больше длины src, то возвращается полная копия src.

Решение учителя:
````
package solution

// BEGIN

// IntsCopy copies a slice of ints src with max length len.
// If maxLen is greater than src it returns a full copy of src.
// If maxLen is zero or negative it returns empty slice.
func IntsCopy(src []int, maxLen int) []int {
	if maxLen <= 0 {
		return []int{}
	}

	if maxLen > len(src) {
		maxLen = len(src)
	}

	cp := make([]int, maxLen)
	copy(cp, src)

	return cp
}

// END
````
Ваше решение:
````
package solution

// BEGIN (write your solution here)
func IntsCopy(src []int, maxLen int) []int {
	if maxLen > 0 {
		var outp []int
		for i := 0; i < maxLen && i < len(src); i++ {
			outp = append(outp, src[i])
		}
		return outp
	} else {
		return []int{}
	}
}
````
#### Сортировка слайсов

Сортировка массива — распространненая задача в программировании. Во всех языках существуют готовые решения для этой задачи, и Go — не исключение. Стандартный пакет sort предоставляет функции для сортировки:
````
nums := []int{2,1,6,5,3,4}

sort.Slice(nums, func(i, j int) bool {
    return nums[i] < nums[j]
})

fmt.Println(nums) // [1 2 3 4 5 6]
````
Рассмотрим функцию Slice(x interface{}, less func(i, j int) bool). В описании функции присутствует неизвестный тип данных interface{}. Понятие интерфейса будет рассмотренно в следующих модулях курса. Следует запомнить, что пустой интерфейс interface{} в Go означает тип данных, под который подходит любой другой тип. Например:
````
func Print(arg interface{}) {
    fmt.Println(arg)
}

func main() {
    Print("hello!")
    Print(123)
    Print([]int{1,5,10})
}
````
Вывод:
````
hello!
123
[1 5 10]
````
То есть в функцию Slice(x interface{}, less func(i, j int) bool) передается слайс любого типа данных, как первый аргумент. Вторым аргументом передается функция, которая берет элементы по индексу и определяет должен ли элемент по индексу i находиться перед элементом по индексу j.

"Под капотом" в функции sort.Slice используется быстрая сортировка. В пакете также присутствует сортировка вставками sort.SliceStable:
````
nums := []int{2,1,6,5,3,4}

sort.SliceStable(nums, func(i, j int) bool {
    return nums[i] < nums[j]
})

fmt.Println(nums) // [1 2 3 4 5 6]
````
Выбор алгоритма зависит от набора и размера данных, архитектуры процессора, скорости доступа к памяти, то есть от многих факторов. Для большинства стандартных случаев используется sort.Slice, пока производительность или нестабильность алгоритма не станет "узким горлышком".

Задание

Реализуйте функцию UniqueSortedUserIDs(userIDs []int64) []int64, которая возвращает отсортированный слайс, состоящий из уникальных идентификаторов userIDs. Обработка должна происходить in-place, то есть без выделения доп. памяти.

Решение учителя:
````
package solution

import "sort"

// BEGIN

// UniqueSortedUserIDs sorts and removes duplicates from the userIDs slice.
func UniqueSortedUserIDs(userIDs []int64) []int64 {
	if len(userIDs) < 2 {
		return userIDs
	}

	sort.SliceStable(userIDs, func(i, j int) bool { return userIDs[i] < userIDs[j] })
	uniqPointer := 0
	for i := 1; i < len(userIDs); i++ {
		if userIDs[uniqPointer] != userIDs[i] {
			uniqPointer++
			userIDs[uniqPointer] = userIDs[i]
		}
	}

	return userIDs[:uniqPointer+1]
}

// END
````
Ваше решение:
````
package solution

import "sort"

// BEGIN (write your solution here)
func UniqueSortedUserIDs(userIDs []int64) []int64 {
	sort.SliceStable(userIDs, func(i, j int) bool {
		return userIDs[i] < userIDs[j]
	})
	for i := 0; i < len(userIDs) - 1;  {
		if userIDs[i] == userIDs[i+1] {
			copy(userIDs[i:], userIDs[i+1:])
			userIDs[len(userIDs)-1] = 0
			userIDs = userIDs[:len(userIDs)-1]
		} else {
			i++
		}
	}
	return userIDs
}
````
#### remove - удаление элемента в слайсе
````
import "fmt"
func remove(slice []int, i int) []int {
	copy(slice[i:], slice[i + 1:])
	return slice[:len(slice) - 1]
}
func main() {
	s := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Printf("s: len: %d, cap: %d, slice: %v\n", len(s), cap(s), s)
	remove(s, 5)
	fmt.Printf("s: len: %d, cap: %d, slice: %v\n", len(s), cap(s), s)
	a := remove(s, 5)
	fmt.Printf("s: len: %d, cap: %d, slice: %v\n", len(s), cap(s), s)
	fmt.Printf("a: len: %d, cap: %d, slice: %v\n", len(a), cap(a), a)
	a[5] = 0
	fmt.Printf("s: len: %d, cap: %d, slice: %v\n", len(s), cap(s), s)
	fmt.Printf("a: len: %d, cap: %d, slice: %v\n", len(a), cap(a), a)
}
// s: len: 9, cap: 9, slice: [1 2 3 4 5 6 7 8 9]
// s: len: 9, cap: 9, slice: [1 2 3 4 5 7 8 9 9]
// s: len: 9, cap: 9, slice: [1 2 3 4 5 8 9 9 9]
// a: len: 8, cap: 9, slice: [1 2 3 4 5 8 9 9]
// s: len: 9, cap: 9, slice: [1 2 3 4 5 0 9 9 9]
// a: len: 8, cap: 9, slice: [1 2 3 4 5 0 9 9]
````
#### использовать слайс как стек:
````
stack = append(stack, v) // push v
top := stack[len(stack) - 1] // top of stack
stack = stack[:len(stack) - 1] // pop
````
### maps
````
ages := make(map[string]int)
ages := make(map[string]int) {
 	"alice": 31,
 	"bob": 34,
}

// рекомендуемое создание с обозначением размера
m := make(map[int]string, 10) // 10-size
````
удаление элемента
````
delete(ages, "alice")
````
запрещено брать указатель элемента: нельзя &ages["bob"]

проверить, есть ли такой элемент:
````
age, ok := ages["bob"]
if !ok {/*...*/} //если такого ключа нет, то ...
````
Для проверки существования ключа можно использовать мапу с пустыми структурами
````
// пустая структура struct{} — это тип данных, который занимает 0 байт
// используется, когда нужно проверять в мапе только наличие ключа
cache := make(map[int64]struct{})

// проверяем есть ли ключ `key` в мапе
_, ok = cache['key']
Println.out(ok)  // false

// добавим ключ и проверим вновь
cache['key'] = struct{}{}
_, ok = cache['key']
Println.out(ok)  // true
````
перебор значений:
````
for name, age := range ages {/*...*/}
````
Ключом можно делать всё, что можно сравнивать: структуры, указатели, интерфейсы

Слайсы нельзя делать ключом

Мапы в Go всегда передаются по ссылке:
````
package main

import (
    "fmt"
)

func main() {
    m := map[int]string{1: "hello", 2: "world"}

    modifyMap(m)

    fmt.Println(m) // вывод: map[1:changed 2:world 200:added]
}

func modifyMap(m map[int]string) {
    m[200] = "added"

    m[1] = "changed"
}
````
Задание

Реализуйте функцию UniqueUserIDs(userIDs []int64) []int64, которая возвращает слайс, состоящий из уникальных идентификаторов userIDs. Порядок слайса должен сохраниться.

Решение учителя:
````
package solution

// BEGIN

// UniqueUserIDs removes duplicates from the userIDs slice saving the IDs order.
func UniqueUserIDs(userIDs []int64) []int64 {
	// пустая структура struct{} — это тип данных, который занимает 0 байт
	// используется, когда нужно проверять в мапе только наличие ключа
	processed := make(map[int64]struct{})

	uniqUserIDs := make([]int64, 0)
	for _, uid := range userIDs {
		if _, ok := processed[uid]; ok {
			continue
		}

		uniqUserIDs = append(uniqUserIDs, uid)
		processed[uid] = struct{}{}
	}

	return uniqUserIDs
}

// END
````
Ваше решение:
````
package solution

// BEGIN (write your solution here)
func UniqueUserIDs(userIDs []int64) []int64 {
	for i := 0; i < len(userIDs) - 1; i++ {
		for j := i + 1; j < len(userIDs); {
			if userIDs[i] == userIDs[j] {
				copy(userIDs[j:], userIDs[j+1:])
				userIDs[len(userIDs)-1] = 0
				userIDs = userIDs[:len(userIDs)-1]
			} else {
				j++
			}
		}

	}
	return userIDs
}
````
Стоит учитывать, что порядок ключей в мапе рандомизирован:
````
numExistence := make(map[int]bool, 0)

// записали по порядку числа от 0 до 9
for i := 0; i < 10; i++ {
    numExistence[i] = true
}

// обходим мапу и выводим ключи
for num := range numExistence {
    fmt.Println(num)
}
````
Вывод:
````
8
1
2
3
6
7
9
0
4
5
````
Задание

Реализуйте функцию MostPopularWord(words []string) string, которая возвращает самое часто встречаемое слово в слайсе. Если таких слов несколько, то возвращается первое из них.

Решение учителя:
````
package solution

// BEGIN

// MostPopularWord returns most popular word from the words slice.
// If there are multiple popular words it returns the first one depending on the words slice order.
func MostPopularWord(words []string) string {
	wordsCount := make(map[string]int, 0)
	mostPopWord := ""
	max := 0

	for _, word := range words {
		wordsCount[word]++
		if wordsCount[word] > max {
			max = wordsCount[word]
			mostPopWord = word
		}
	}

	return mostPopWord
}

// END
````
Ваше решение:
````
package solution

// BEGIN (write your solution here)
func MostPopularWord(words []string) string {
	tmpMap := make(map[string]int)
	var outp string
	tmpNmb := 0
	for _, word := range words {
		_, ok := tmpMap[word]
		if ok {
			tmpMap[word]++
		} else {
			tmpMap[word] = 1
		}
		if tmpNmb < tmpMap[word] {
			tmpNmb = tmpMap[word]
			outp = word
		}
	}
	return outp
}
````
### set
````
var s0 map[string]bool // так хуже, лучше как ниже:
var s1 map[string]struct{}
````
### struct
В Go нет классов и привычной реализации ООП. Вместо классов в языке используются структуры — наборы полей, имеющих название и тип данных. Объявление структуры имеет следующий вид:
````
type Person struct {
    // [название поля] [тип данных]
    Name string
    Age int
}

func main() {
    p := Person{Name: "John", Age: 25}

    p.Name // "John"
    p.Age // 25
}
````
Структуру можно инициализировать, не передавая значения. В этом случае каждое поле примет свое «нулевое» значение:
````
func main() {
    p := Person{}

    p.Name // ""
    p.Age // 0
}
````
Регистр первой буквы в названии структуры и полей означает публичность точно так же, как в переменных и функциях. Если первая буква заглавная, то структуру можно инициализировать во внешних пакетах. В противном случае она доступна только в рамках текущего пакета:
````
type Person struct { // структура публична
    Name string // поле публично

    wallet wallet // поле приватно: можно обращаться только внутри текущего пакета
}

type wallet struct { // структура приватна: можно инициализировать только внутри текущего пакета
    id string
    moneyAmount float64
}
````
У любого поля структуры можно указать теги. Они используются для метаинформации о поле для сериализации, валидации, маппинга данных из БД и тд. Тег указывается после типа данных через бектики:
````
type User struct {
    ID int64 `json:"id" validate:"required"`
    Email string `json:"email" validate:"required,email"`
    FirstName string `json:"first_name" validate:"required"`
}
````
Тег json используется для названий полей при сериализации/десериализации структуры в json и обратно:
````
package main

import (
    "encoding/json"
    "fmt"
)

type User struct {
    ID        int64  `json:"id"`
    Email     string `json:"email"`
    FirstName string `json:"first_name"`
}

func main() {
    u := User{}
    u.ID = 22
    u.Email = "test@test.com"
    u.FirstName = "John"

    bs, _ := json.Marshal(u)

    fmt.Println(string(bs)) // {"id":22,"email":"test@test.com","first_name":"John"}
}
````
Тег validate используется Go-валидатором. В следующем примере присутствует вызов функции у структуры v.Struct(u). Функции структур — методы — мы разберем подробно в соответствующем уроке, а пока просто посмотрите, как происходит вызов:
````
package main

import (
    "fmt"
    "github.com/go-playground/validator/v10"
)

type User struct {
    ID        int64  `validate:"required"`
    Email     string `validate:"required,email"`
    FirstName string `validate:"required"`
}

func main() {
    // создали пустую структуру, чтобы проверить валидацию
    u := User{}

    // создаем валидатор
    v := validator.New()

    // метод Struct валидирует переданную структуру и возвращает ошибку `error`, если какое-то поле некорректно
    fmt.Println(v.Struct(u))
}
````
Вывод программы:
````
Key: 'User.ID' Error:Field validation for 'ID' failed on the 'required' tag
Key: 'User.Email' Error:Field validation for 'Email' failed on the 'required' tag
Key: 'User.FirstName' Error:Field validation for 'FirstName' failed on the 'required' tag
````

````
type Point struct{ X, Y int }
p := Point{1, 2}
p := Point{X: 1, Y: 2}
````
````
import "fmt"
type Point struct{ X, Y int }
func main() {
	p1 := Point{1, 2}
	p2 := Point{1, 2}
	p3 := Point{2, 1}
	fmt.Println(p1 == p2) //true
	fmt.Println(p1 == p3) //false
}
````
Структуру можно использовать как ключ в map
#### struct embedding:
````
type Point struct {
    X, Y int
}
type Circle struct {
 	Point
 	Radius int
}
c := Circle{
 	Point: Point{X: 10, Y: 10},
 	Radius: 1,
}
c.X = 0 // == c.Point.X = 0
```` 
Синтаксис структуры позволет вписывать после поля 
произвольный строковой литерал, который можно использовать 
для передачи опций для функции конвертации форматов, серизации/десеризации. Например, требования к порлям могут отличаться, и можно указать имя, которое декодер Jsona будет использовать
#### json
````
type Movie struct {
	Title string
	Year int `json:"year"`
	Color bool `json:"color,omitempty"`
	Actors []string
}
````
#### marshal
````
data, err := json.Marshal(movies)
if err != nil {
	log.Fatalf("JSON marshaling failed: %s", err)
}
fmt.Printf("%s\n", data)
````
````
data, err := json.MarshalIndent(movies, "", "     ")
if err != nil {
	log.Fatalf("JSON marshaling failed: %s", err)
}
fmt.Printf("%s\n", data)
````
#### unmarshal
````
var movie Movie
if err := json.Unmarshal(data, &movie); err != nil {
	log.Fatalf("JSON unmarshaling failed: %s", err)
}
fmt.Println(movie)
````
#### github
````
func SearchIssues(terms []string) (*IssueSearchResult, error) {
	q := url.QueryEscape(strings.Join(terms, " "))
	resp, err := http.Get(IssueUrl + "?q=" + q)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("searc query failed: %s", resp.Status)
	}
	var result IssueSearchResult
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}
	return &result, nil
}
````
### functions
если возвращиемое значение ф-ции имеет имя, то этот элнмент ненадо объявлять и не надо указывать в ретёрне: просто указать ретёрн и вернутся те значения этих переменных, которые есть у них на момент вызова ретёрна.
````
func CountWordAndImages(url string)(words, images int, err error) {/*...*/}
````
### errors
если функция возвращает ошибку, то она всегда возвращает её последним аргументом. Текст ошибки должен быть в lowercase. 
Ошибку надо пояснить в соответствии с этапом программы:
````
doc, err := html.Parse(resp.Body)
if err != nil {
	return nil, fmt.Errorf("parsing %s as HTML: %w", url, err)
}
````
### EOF - нечего больше читать, конец файла
````
import "errors"
var EOF = errors/New("EOF")
in := bufio.NewReader(os.Stdin)
for {
	r, _, err := in.ReadRune()
	if err == io.EOF {
		break  // finishing reading
	}
	if err != nil {
		return fmt.Errorf("read failed: %v", err)
	}
	// use r
}
````
### function values
имя функции может быть передано как значение
````
import "fmt"
func Inc(i int) int { return i + 1 }
var f func(i int) int
func main() {
	// f = func(i int) int { return i * 2 }
	for i:= 0; i < 5; i++ {
		if f == nil {
			f = Inc
		}
		fmt.Println(f(i))
	}
}
````
### recursion
рекурсивная ф-ция объявляется отдельно от определения (var visit func(n *Node) ), иначе к ней невозможно обратиться рекурсивно
````
type Node struct {
	V int
	L, R *Node
}
func PrintAll(w io.writer, root *Node) {
	var visit func(n *Node) // объявляется отдельно 
	visit = func(n *Node) {
		fmt.Fprintln(w, n.V)
		if n.L != nil {
			visit(n.L)
		}
		if n/R != nil {
			visit(n.R)
		}
	}
	visit(root)
}
````