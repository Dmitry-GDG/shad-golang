# Лекция 3. Интерфейсы, ошибки, паники
# methods

в обычном методе аргумент (p Path) передаётся просто по значению. "то значит, что тут нельзя написать какой-то метод, который сам путь меняет (например, нельзя trim или append)
````
import (
	"fmt"
	"math"
)
type Point struct{ X, Y float64}
func (p Point) Distance(q Point) float64{
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}
type Path []Point
func (p Path) Distance() float64 {
	sum := 0.0
	for i := range p {
		if i > 0 {
			sum += p[i - 1].Distance(p[i])
		}
	}
	return sum
}
func main() {
	var p = Path{{1.0, 2.5}, {3.6, 3.8}, {4.5, 4.8},}
	fmt.Println(p.Distance())
	var p1 = Path{{1, 1}, {5, 1}, {5, 4}, {1, 1}, }
	fmt.Println(p1.Distance())
}
````
Метод м.б. определён не на самом типе, а на указателе на тип. И такорм случае аргумет является Pointer Receiver и смысл его в том, что он может сам объект поменять (может поменять сам тип, который передали по указателю). Особенность Go: в такой метод можно передать nil указатель, и такой метод продолжит работать нормально. Внутри такого метода можно проверить на nil и соответственно обработать, например:
````
type IntList struct {
	Value int
	Tail *IntList
}
func (list *IntList) Sum() int {
	if list == nil {
		return 0
	}
	return list.Value + list.Tail.Sum()
}
````
## Embedding
Методы взаимодействуют с эмбеддингом точно так же, как ембеддинг взаимодействует с полями. То есть: если у нас есть какая то структура (Point), мы можем её заимбеддить в другую структуру. Это значит: внутри новой структуры мы на отдельной строке должны указать имя типа, который эмбеддим, и не давать ему имени поля. 
````
type Point struct { X, Y float64}
type ColoredPoint struct {
	Point
	Color color.RGBA
}
````
В это момент все методы с Point переклеятся на нашу структуру ColoredPoint, и их можно будет позвать, как указано в примере (p.Distance). Но тут есть один момент: аргумент этого метода надо явно назад достать (q.Point)
````
type Point struct { X, Y float64}
type ColoredPoint struct {
	Point
	Color color.RGBA
}
red := color.RGBA{255, 0, 0,255}
blue := color.RGBA{, 0, 255, 255}
var p = ColoredPoint{Point{1, 1}, red}
var q = ColoredPoint{Point{5, 4}, blue}
fmt.Println(p.Distance(q.Point)) // 5
p.ScaleBy(2)
q.ScaleBy(2)
fmt.Println(p.Distance(q.Point)) // 10
````
Пример использования:
````
var cache = struct {
	sync.Mutex
	mapping map[string]string
} {
	mapping: make(map[string]string)
}
func Lookup(key string) string {
	cache.Lock()
	v := cache.mapping[key]
	cache.Unlock()
	return v
}
````
Из методов можно сделать функцию. В примере ниже мы захватили в ф-цию переменную p, и ф-ция стала ожидать только один аргумент q
````
p := Point{1, 2}
q := Point{4, 6}
distanceFromP := p.Distance
fmt.Println(distanceFromP(q))
var origin Point
fmt.Println(distanceFromP(origin))
scaleP := p.ScaleBy
scaleP(2)
scaleP(3)
scaleP(10)
````
## Method Values
````
type Rocket struct { /*...*/ }
func (r *Rocket) Launch() { /*///*/ }
r := new(Rocket)
time.AfterFunc(10 * time.Second, func90 { r.Launch() }) // эвивалентны, работают одинаково
time.AfterFunc(10 * time.Second, r.Launch) // эвивалентны, работают одинаково
````
## Method Expression
````
p := Point{1, 2}
q := Point{4, 6}
distance := Point.Distance  // method expression
fmt.Println(distance(p, q)) // "5"
fmt.Printf("%T\n", distance) // "func(Point, Point) float64"
````
В случае с методами, который принимают ресивер по указателю:
````
scale := (*Point).ScaleBy
scale(&p, 2)
fmt.Println(p) // "{2, 4}"
fmt.Printf("%T\n", scale) // "func(*Point, float64)"
````
Сеттеры и геттеры не приветствуются. 

Для изменения данных просто имя данных начинают с заглавной буквы

------------
------------
------------


# Interfaces
````
type Writer interface {
    Write(p []byte) (n int, err error)
}
````
## Embedding interfaces
````
package io
type Reader interface {
	Read(p []byte) (n int, err error)
}
type Closer interface {
	Close() error
}
type ReadWriter interface {
	Reader
	Writer
}
type ReadWriteCloser interface {
	Reader
	Writer
	Closer
}
````
Как пользоваться:

Можно создать переменную, у которой типом является Interface. Например, io.Writer. После этого этой переменной присвоить значение, которое удовлетворяет этому типу:
````
var w io.Writer
w = os.Stdout // OK: *os.File has Write method
w = New(bytes.Buffer) // OK: *bytes.Buffer has Write method
w = time.Second // compile error: time.Duration lacks Write method
````
### Особенности интерфейса:
Если у типа есть метод, определённый на указателе на этот тип, то мы можем вызвать этот метод как на указателе, так и на самом типе (происходит неявное приведение, неявное разыменование)

Но! На интерфейсе всё происходит иначе: не происходит неявного разыменования
````
package fmt
type Stringer interface {
	String() string
}
...
type IntSet struct { /*...*/ }
func (*IntSet) String() string
var s IntSet
var _=s.String() // OK: s is a variable and &s is a method
var _ fmt.Stringer = &s // OK
var _ fmt.Stringer = s // compiling error: InSet lacks String method
````
### interface{}
````
var any interface{}
any = true
any = 12.34
any = "hello"
any = map[string]int{"one":1}
any = new(bytes.Buffer)
````
### interface satisfaction
Что делать, если у меня есть где-то интерфейс, в другом пакете тип, и я хочу убедиться, что этот тип реализует этот интерфейс, потому что итерфейс могут поменять, а я забуду это поправить. Для этих целей мы в своём пакете на самом верхнем уровне создаёт переменную с названием "_" (переменная, к которой нельзя обратиться), у которой прописываем явный тип нужного нам интерфейса. Дальше туда присваиваем значение своего типа (для указателей - nil или new(..)). И получаем ошибку на самом раннем этапе компиляции
````
// *bytes.Buffer must satisfy io.Writer
var _ io.Writer = (*bytes.Buffer)(nil)
````
### interface values (реализация интерфейсов)
интерфейс состоит из двух указателей:
````
[указатель на конкретный тип (например *File)]
[указатель на ячейку памяти, где хранится значение переменной]
````
У nil в зависимости от контекста - разный смысл: если присвоить переменной интерфейса nil - в обе ячейки запишется nil, nil
````
var w io.Writer
w = os.Stdout
w = new(bytes.Buffer)
w = nil
````
А что будет, если присвоить этой переменной нулевой указатель данного типа?
````
var buf *bytes.Buffer // создаётся нулевой указатель на bytes.Buffer
if debug {
	buf = new(bytes.Buffer) // enable collection of output
}
f(buf) // NOTE: subtly incorrect!
/* If out is non-nul, output will be written to it/ */
func f(out io.Writer) {
	/* ... do something */
	if out != nil { // в нашем случае мы не присваивали переменной nil,
					// в её первой ячейке лежит указатель на тип, 
					// при этом значение не является нулевым, 
					// это не считается за нулевой интерфейс.
					// в нашем случае out всегда не равен nil
		out.Write([]byte("done!\n"))
	}
}
````
Как поправить наш пример? Изначально определить переменную как io.Writer, nогда её начальные значения будут nil, nil. И дальше её инициализировать, если надо
````
var buf io.Writer
if debug {
	buf = new(bytes.Buffer)
}
f(buf) // OK
````
## Что можно делать с интерфейсами?
#### - Если есть интерфейс, то можно вызывать метод на этом интерфейсе:
````
out.Write([]byte("done!\n")) // см пример выше
````
#### - Type Assertions (приведение интерфейса)
Это чем то похоже на приведение типов, только в начале мы пишем интерфейс - точка - новый тип
#### type assertiom Concrete type (на конкретный тип):
````
var w io.Writer
w = io.Stdout
f := w.(*os.File) // success: f == os.Stdout, мы угадали, внутри именно такой тип
c := w.(*bytes.Buffer) // panic: interface holds *os.File, not *bytes.Buffer
````
#### type assertion interface type (скастить интерфейс к другому интерфейсу):
То есть Go сходит к существующему интерфейсу и проверит, реализует ли он метод, который необходим новому интерфейсу? Если реализует - то ОК
````
var w io.Writer
w = io.Stdout
rw := w.(io.ReadWriter) // success: *os.File has both Read and Write
w = new(ByteCounter)
rw := w.(io.ReadWriter) // panic: *ByteCounter has no Read method
````
#### - Доставать тип интерфейса
Проверка (check type):
````
val, ok := w.(io.ReadWriter)
if ok {...}
````
````
var w io.Writer = os.Stdout
f, ok := w.(*os.File) // success: ok, f == os.Stdout
b, ok := w.(*bytes.Buffer) // failure: !ok, b == nil
````
### один интерфейс м.б. вложен в другой
При этом есть нюанс: в верхнем присваивании nil примется, в нижнем - будет паника
````
var w io.Writer
var rw io.ReadWriter
w = rw // OK: io.ReadWriter is assignable to io.Writer
w = rw.(io.Writer) // fails(panic) only if rw == nil
````
При работе с пустым интерфейсом часто бывает ситуация, когда необходимо написать лесенку из if-ов. Примерно такая лесенка написана внутри пакета fmt:
````
func sqlQuote(x interface{}) string {
	if x == nil {
		return "NULL"
	} else if _, ok := x.(int); ok {
		return fmt.Sprintf("%d", x)
	} else if _, ok := x.(uint); ok {
		return fmt.Sprintf("%d", x)
	} else if b, ok := x.(bool); ok {
		if b {
			return "TRUE"
		}
		return "FALSE"
	} else if s, ok := x.(string); ok {
		return sqlQuoteStrings(s) // not shown
	} else {
		panic(fmt.Sprintf("unexpected type %T: %v", x, x))
	}
}
````
или через type switch (здесь type - это не заглушка, надо писать именно type):
````
switch x.(type) {
case nil: 
case int: 
case bool:
case string: 
default:
}
````
или (что удобней) сразу присвоить переменной значение в зависимости от её типа:
````
switch x := x.(type) { /*...*/ }
````
Пример:
````
func sqlQuote(x interface{}) string {
	switch v := x.(type) {
	case nil:
		return "NULL"
	case int, uint:
		return fmt.Sprintf("%d", v) // v has type interface{} here (так как не м.б два разныз интерфейса одновременно)
	case bool: // v has type bool here
		if v {
			return "TRUE"
		}
		return "FALSE"
	case string: // v has type string here
		return sqlQuoteStrings(s) // not shown
	default:
		panic(fmt.Sprintf("unexpected type %T: %v", v, v))
	}
}
````
------------
------------
------------
# panic, defer, recover
## panic
Паника - для ситуаций, когда в коде баг. Она не для ситуаций, когда в коде что-то сломалось. Работает только с одним аргументом, обычно - строкой (можно передать и свой кастомный тип), который должен описать ошибку.  
````
switch s := suit(drawCard()); s {
case "Spades": //...
case "Hearts": //...
case "Diamonds": //...
case "Clubs": //...
default:
    panic(fmt.Sprintf("invalid suit %q", s)) // Joker?
}

func Reset(x *Buffer) {
    if x == nil {
        panic("x is nil") // unnecessary!
    }
    x.elements = nil
}
````
Паника работает в одной горутине. Внутри этой горутины паника раскручивает стек. Она идет по всем ф-циям внутри нашего вызова и выполнять все defer-ы, которые были созданы. Мы хотим, чтобы defer-ы исполнились в случае паники. Деферы вызываются по правилам стека LIFO.
````
//defer
func CopyFile(dstName, srcName string) (written int64, err error) {
    src, err := os.Open(srcName)
    if err != nil {
        return
    }
    defer src.Close()
    dst, err := os.Create(dstName)
    if err != nil {
        return
    }
    defer dst.Close()
    return io.Copy(dst, src)
}
````
Казалось бы: зачем нам defer, если программа всё-равно умрёт? - Чтоб файл в любом случае закрылся и никуда не утёк. 

От паники можно восстановиться (recover). Для этого внутри defer-а надо написать ф-цию recover, которая имеет следующий смысл: она возвращает тот аргумент, который был ей передан, а если был штатный вызов defer-а (не из-за паники), то она возвращает nil. Ещё она прекращает панику.

Функция должна вернуть некоторые переменные. При вызове программы эти переменные инициализируются нулями (например), на момент паники эти переменные уже имеют некоторые значения. В этом примере в случае паники мы выходим из ф-ции, но выходим без паники, возвращая переменные со значениями, которые у них были на момент пваники. Наша программа штатно продолжает работать, получив из ф-ции эти значения.
````
//recover
func Parse(input string) (s *Syntax, err error) {
    defer func() {
        if p := recover(); p != nil {
            err = fmt.Errorf("internal error: %v", p)
        }
    }()
    // ...parser...
}
````
defer может паниковать:
````
func main() {
    defer fync() {
        recover()
        fmt.Println("Checkpoint 1")
        panic(1)
    }()
    defer func() {
        recover()
        fmt.Println("Checkpoint 2")
        panic(2)
    }()
    panic(999)
}
// Output:
// Checkpoint 2
// Checkpoint 1
// panic: 999 [recovered]
//     panic: 2 [recovered]
//     panic: 1
````
