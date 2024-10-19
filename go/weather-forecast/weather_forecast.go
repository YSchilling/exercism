// Package weather provides you with current weather conditions.
package weather

// CurrentCondition holds the current weather conditions.
var CurrentCondition string

// CurrentLocation holds the current location where you want to get weather data for.
var CurrentLocation string

// Forecast function takes a city and a current condition and returns a little forecast string.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
