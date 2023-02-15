package conv

// Konverterer Farhenheit til Celsius
func FarhenheitToCelsius(value float64) float64 {
	Celsius := (value - 32) * (5 / 9)
	return Celsius
}

func CelsiusToFahrenheit(value float64) float64 {
	Farhrenheit := value*(9/5) + 32
	return Farhrenheit
}

func KelvinToFarhenheit(value float64) float64 {
	Kelvin := (value-32)*(5/9) + 273.15
	return Kelvin
}

// De andre konverteringsfunksjonene implementere her
// ...
