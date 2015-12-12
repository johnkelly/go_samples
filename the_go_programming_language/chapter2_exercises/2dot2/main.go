//Converts numbers into various units
package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	for _, arg := range os.Args[1:] {
		value, err := strconv.ParseFloat(arg, 64)

		if err != nil {
			fmt.Fprintf(os.Stderr, "cf: %v\n", err)
			os.Exit(1)
		}
		f := Fahrenheit(value)
		c := Celsius(value)
		fmt.Println("Temperatures:")
		fmt.Printf("%s = %s, %s = %s\n\n", f, FToC(f), c, CToF(c))

		feet := Foot(value)
		meters := Meter(value)

		fmt.Println("Lengths:")
		fmt.Printf("%s = %s, %s = %s\n\n", feet, FootToMeters(feet), meters, MeterToFeet(meters))

		pounds := Pound(value)
		kilos := Kilogram(value)

		fmt.Println("Weight:")
		fmt.Printf("%s = %s, %s = %s\n\n", pounds, PoundToKilos(pounds), kilos, KiloToPounds(kilos))
	}
}

// Celsius is the temperature in Celsius
type Celsius float64

// Fahrenheit is the temperature in Fahrenheit
type Fahrenheit float64

// Foot is a measure of length
type Foot float64

// Meter is a measure of length
type Meter float64

// Pound is a measure of weight
type Pound float64

// Kilogram is a measure of weight
type Kilogram float64

func (c Celsius) String() string    { return fmt.Sprintf("%g°C", c) }
func (f Fahrenheit) String() string { return fmt.Sprintf("%g°F", f) }
func (f Foot) String() string       { return fmt.Sprintf("%g ft", f) }
func (m Meter) String() string      { return fmt.Sprintf("%g m", m) }
func (p Pound) String() string      { return fmt.Sprintf("%g lb", p) }
func (k Kilogram) String() string   { return fmt.Sprintf("%g kg", k) }

// CToF converts a Celsius tempearture to Fahrenheit
func CToF(c Celsius) Fahrenheit { return Fahrenheit(c*9/5 + 32) }

// FToC converts a Fahrenheit tempearture to Celsius
func FToC(f Fahrenheit) Celsius { return Celsius((f - 32) * 5 / 9) }

// FootToMeters converts a Foot measurement to Meters
func FootToMeters(f Foot) Meter { return Meter(f * 3) }

// MeterToFeet converts a Meter measurement to Feet
func MeterToFeet(m Meter) Foot { return Foot(m / 3) }

// KiloToPounds converts a kilogram to pounds
func KiloToPounds(k Kilogram) Pound { return Pound(k / 2.2) }

// PoundToKilos converts a pound to kilograms
func PoundToKilos(p Pound) Kilogram { return Kilogram(p * 2.2) }
