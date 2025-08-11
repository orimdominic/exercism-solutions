// Package weather holds operations for telling the weather.
package weather

// CurrentCondition tells the current weather condition of CurrentLocation.
var CurrentCondition string

// CurrentLocation tells the current location.
var CurrentLocation string

// Forecast returns the forecase for city/condition.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
