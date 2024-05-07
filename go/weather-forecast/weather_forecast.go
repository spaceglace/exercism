// Package weather provides a function to forecast the current weather condition
package weather

var CurrentCondition string
var CurrentLocation string

func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
