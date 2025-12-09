// Package weather provides methods to format a message
// providing details about condition and location.
package weather

var (
    // CurrentCondition represents the current condition.
	CurrentCondition string
    // CurrentLocation represents the current location.
	CurrentLocation  string
)

// Forecast function provide you the forecast for a specific location
// city is the city name and condition will be condition of it.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
