package main

import "fmt"

const ebulicaoK = 373.15

func main() {
	tempKelvin := ebulicaoK
	tempCelsius := tempKelvin - 273.15
	fmt.Println("A temperatura de ebuliçao da água em Kelvin é", tempKelvin, "K, e em Celsius é", tempCelsius, "°C.")

}
