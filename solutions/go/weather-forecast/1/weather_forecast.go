//Package weather returns the weather forecast of a city.
package weather

//CurrentCondition represents the weather conditions of a place.
var CurrentCondition string
//CurrentLocation represents the place whose forecast is being described.
var CurrentLocation string

/*Forecast takes a city and its weather condition as arguments and returns a string which tells the forecast of a particular city.*/
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
