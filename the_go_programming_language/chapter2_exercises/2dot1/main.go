// Package tempconv performs Celsius and Fahrenheit conversions.
package main

import "fmt"

// Celsius is the temperature in Celsius
type Celsius float64

// Fahrenheit is the temperature in Fahrenheit
type Fahrenheit float64

// Kelvin is a unit for measuring temperature
type Kelvin float64

// Temperature constants
const (
	AbsoluteZeroC Celsius = -273.15
	FreezingC     Celsius = 0
	BoilingC      Celsius = 100
)

func (c Celsius) String() string    { return fmt.Sprintf("%g°C", c) }
func (f Fahrenheit) String() string { return fmt.Sprintf("%g°F", f) }
func (k Kelvin) String() string     { return fmt.Sprintf("%g K", k) }

// CToF converts a Celsius tempearture to Fahrenheit
func CToF(c Celsius) Fahrenheit { return Fahrenheit(c*9/5 + 32) }

// FToC converts a Fahrenheit tempearture to Celsius
func FToC(f Fahrenheit) Celsius { return Celsius((f - 32) * 5 / 9) }

// KToC converts a Kelvin tempearture to Celsius
func KToC(k Kelvin) Celsius { return Celsius(k + Kelvin(AbsoluteZeroC)) }

func main() {
	fmt.Printf("Brr! %v\n", AbsoluteZeroC)
	fmt.Println(CToF(BoilingC))

	fmt.Println(KToC(45))
}
