//Package weather provides tools to analyze the weather.
package weather

var (
    //CurrentCondition represents a certain weather condition as a string.
	CurrentCondition string
    //CurrentLocation represents a certain city as a string.
	CurrentLocation  string
)

//Forecast takes a location and condition and return a formatted weather forecast string.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
