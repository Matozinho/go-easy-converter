package main

import "fmt"

const kEbullition float32 = 373.0

func main() {
	tempK := kEbullition
	tempC := tempK - 273.0

	fmt.Printf("A temperatura de ebulição da água em K = %g, temperatura de ebulição da água em °C = %g.\n", tempK, tempC)
}
