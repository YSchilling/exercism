package meteorology

import "fmt"

type TemperatureUnit int

const (
	Celsius    TemperatureUnit = 0
	Fahrenheit TemperatureUnit = 1
)

func (p TemperatureUnit) String() string {
	converter := []string{"°C", "°F"}
	return converter[p]
}

type Temperature struct {
	degree int
	unit   TemperatureUnit
}

// Add a String method to the Temperature type

func (p Temperature) String() string {
	return fmt.Sprintf("%d %s", p.degree, p.unit)
}

type SpeedUnit int

const (
	KmPerHour    SpeedUnit = 0
	MilesPerHour SpeedUnit = 1
)

// Add a String method to SpeedUnit

func (p SpeedUnit) String() string {
	converter := []string{"km/h", "mph"}
	return converter[p]
}

type Speed struct {
	magnitude int
	unit      SpeedUnit
}

// Add a String method to Speed

func (p Speed) String() string {
	return fmt.Sprintf("%d %s", p.magnitude, p.unit)
}

type MeteorologyData struct {
	location      string
	temperature   Temperature
	windDirection string
	windSpeed     Speed
	humidity      int
}

// Add a String method to MeteorologyData

func (p MeteorologyData) String() string {
	return fmt.Sprintf("%s: %s, Wind %s at %s, %d%% Humidity", p.location, p.temperature, p.windDirection, p.windSpeed, p.humidity)
}
