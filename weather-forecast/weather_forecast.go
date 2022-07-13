// Package weather contains all the information needed to create a weather forecast.
package weather

// CurrentCondition is the current weather condition.
var CurrentCondition string

// CurrentLocation is the current weather location.
var CurrentLocation string

// Forecast receive the City and Location string parameters ans return the weather forecast as string.
// Those parameters are setted to the other two variables above.
// The return are a concatenated string of the weather forecast with location and this respective weather condition.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
