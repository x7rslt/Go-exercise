package convtool

func CToF(c Celsius) Fahrenheit {
	return Fahrenheit(c*9/5 + 32)
}
func FToC(f Fahrenheit) Celsius {
	return Celsius((f - 32) * 5 / 9)
}
func CToK(c Celsius) Kelvin {
	return Kelvin(c + 273.15)
}
func KToC(k Kelvin) Celsius {
	return Celsius(k - 273.15)
}
func IToM(i Inch) Meter {
	return Meter(i * 0.0254)
}
func MToI(m Meter) Inch {
	return Inch(m / 0.0254)
}
func UToC(u Usd) Cny {
	return Cny(u * 6.6114)
}
func CToU(c Cny) Usd {
	return Usd(c / 6.6114)
}
