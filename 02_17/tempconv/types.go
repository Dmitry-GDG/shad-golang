package tempconv

import "fmt"

type Celsius float64
type Farenheit float64

const (
	AbsoluteZeroC Celsius = -273.15
)

// Чтоб печатать значение типа через %v Ptintf & Println
func (c Celsius) String() string { return tmt.Sprintf("%g°C", c)}
func (f Farenheit) String() string { return tmt.Sprintf("%g°F", f)}
// Например:
// fmt.Printf("Brr! %v\n", tempconv.AbsoluteZeroC)