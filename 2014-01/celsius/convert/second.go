package main

import (
	"fmt"
)

type Celsius float32
type Fahrenheit float32

func main() {
	c := Celsius(-1.0)
	f := CelsiusToFahrenheit(c)
	fmt.Printf("%4.2fºC is %4.2fºF\n", c, f)

	// if c == f {
	// 	fmt.Println("Match")
	// }
}

// CelsiusToFahrenheit converts temperature to share with our American friends.
func CelsiusToFahrenheit(c Celsius) Fahrenheit {
	return Fahrenheit((c * 9 / 5) + 32)
}
