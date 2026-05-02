package main

import "fmt"

func celsiusToFahrenheit(c float64) float64 {
	// F = C * 9/5 + 32
	return float64(c)*float64(9/5) + float64(32)
}

// TODO: Implement this function
func celsiusToKelvin(c float64) float64 {
	// K = C + 273.15
	return float64(c) + float64(273.15)
}

func main() {
	celsiusValues := []float64{-40, 0, 20, 37, 100, 10, 40, 32, 28}

	fmt.Println("Temperature Conversion Table")
	fmt.Println("─────────────────────────────────────")
	fmt.Printf("%-10s %-12s %-10s\n", "Celsius", "Fahrenheit", "Kelvin")
	fmt.Println("─────────────────────────────────────")

	for _, c := range celsiusValues {
		// TODO: Call your functions and print each row
		f := celsiusToFahrenheit(c)
		k := celsiusToKelvin(c)
		fmt.Printf("%8.2f°C  %8.2f°F  %7.2fK\n", c, f, k)
	}

	fmt.Println("─────────────────────────────────────")

}
