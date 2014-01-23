package main

import (
	"fmt"
)

type Celsius float32
type Fahrenheit float32

func main() {
	c := Celsius(-1.0)
	fmt.Printf("%v is %v\n", c, c.ToFahrenheit())

	f := Fahrenheit(90)
	fmt.Printf("%v is %v\n", f, f.ToCelsius())
}

// ToFahrenheit converts temperature to share with our American friends.
func (c Celsius) ToFahrenheit() Fahrenheit {
	return Fahrenheit((c * 9 / 5) + 32)
}

// String returns a pretty printed temperature
func (c Celsius) String() string {
	return fmt.Sprintf("%4.2fºC", c)
}

// ToCelsius converts temperature from our American friends.
func (f Fahrenheit) ToCelsius() Celsius {
	return Celsius((f - 32) * 5 / 9)
}

// String returns a pretty printed temperature
func (f Fahrenheit) String() string {
	return fmt.Sprintf("%4.2fºF", f)
}
