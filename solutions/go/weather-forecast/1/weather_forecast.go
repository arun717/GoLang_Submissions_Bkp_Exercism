// Package weather contains Forecast function to help find weather conditions.
package weather

// CurrentCondition variable can be used to store the current weather condition value.
var CurrentCondition string

// CurrentLocation variable can be used to store the current location value.
var CurrentLocation string

//Forecast function is used to display the current location and current weather condition together.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
