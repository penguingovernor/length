package length

import "fmt"

// A Distance represents a physical distance
// as an float64 nanometer count. The representation limits the
// largest representable duration to approximately
// 1.900163142869793e283 light years ≈ diameter of the observable universe (≈ 93 billion ly )
type Distance float64

// Common distances.
//
// To count the number of units in a Distance, divide:
//	meter := length.Meter
//	fmt.Print(float64(meter/length.Millimeter)) // prints 1000
//
// To convert an integer number of units to a Distance, multiply:
//	meters := 10
//	fmt.Print(length.Distance(seconds)*length.Meter) // prints 10m
//
const (
	Nanometer  Distance = 1
	Micrometer          = 1e3 * Nanometer
	Millimeter          = 1e3 * Micrometer
	Centimeter          = 10 * Millimeter
	Meter               = 1e3 * Millimeter
	Kilometer           = 1e3 * Meter
	Inch                = 2.54 * Centimeter
	Feet                = 304.8 * Millimeter
	Yard                = 3 * Feet
	Mile                = 5280 * Feet
	Lightyear           = 9.461e12 * Kilometer
)

var usingMetric = true

// ToggleUnits toggles the units (metric <=> imperial) that are printed whenever the String
// function is called (as is the case in family of printing functions in the fmt package).
// By default the metric system is used.
func ToggleUnits() {
	usingMetric = !usingMetric
}

// UseMetric toggles the units to use the metric system.
// See ToggleUnits for more information.
func UseMetric() {
	usingMetric = true
}

// UseImperial toggles the units to use the imperial system.
// See ToggleUnits for more information.
func UseImperial() {
	usingMetric = false
}

// String returns a string representing the distance in the form "10m" or "10yd".
// The unit that is used is based on the state of the ToggleUnit function.
// As a special case, durations less than one
// meter (or yard) use a smaller unit to ensure
// that the leading digit is non-zero. The zero duration formats as 0m or 0yd.
func (d Distance) String() string {
	// If in metric mode
	if usingMetric {
		return d.printMetric()
	}
	return d.printImperial()
}

func (d Distance) printMetric() string {
	if d >= 1*Meter {
		return fmt.Sprintf("%fm", float64(d)/float64(Meter))
	}
	if d >= 1*Centimeter {
		return fmt.Sprintf("%fcm", float64(d)/float64(Centimeter))
	}
	if d >= 1*Millimeter {
		return fmt.Sprintf("%fmm", float64(d)/float64(Millimeter))
	}
	if d >= 1*Micrometer {
		return fmt.Sprintf("%fµm", float64(d)/float64(Micrometer))
	}
	if d == 0 {
		return fmt.Sprintf("0m")
	}
	return fmt.Sprintf("%fnm", float64(d)/float64(Nanometer))
}

func (d Distance) printImperial() string {
	if d >= 1*Yard {
		return fmt.Sprintf("%fyd", float64(d)/float64(Yard))
	}
	if d >= 1*Feet {
		return fmt.Sprintf("%fft", float64(d)/float64(Feet))
	}
	if d == 0 {
		return fmt.Sprintf("0yd")
	}
	return fmt.Sprintf("%fin", float64(d)/float64(Inch))
}

var unitMap = map[string]float64{
	"nm": float64(Nanometer),
	"um": float64(Micrometer), // U+03BC = Greek letter mu
	"µm": float64(Micrometer), // U+00B5 = micro symbol
	"μm": float64(Micrometer), // U+03BC = Greek letter mu
	"mm": float64(Millimeter),
	"cm": float64(Centimeter),
	"m":  float64(Meter),
	"km": float64(Kilometer),
	"in": float64(Inch),
	"ft": float64(Feet),
	"yd": float64(Yard),
	"mi": float64(Mile),
	"ly": float64(Lightyear),
}

// ParseDistance parses a distance string.
// A distance string is a possibly signed sequence of
// decimal numbers, each with optional fraction and a unit suffix,
// such as "300m" or "-1.5ly"
// Valid time units are "nm", "um" (or "µm"), "mm", "m", "km", "in", "ft", "yd", "mi", "ly".
func ParseDistance(s string) (Distance, error) {
	// WIP
	return Distance(0 * unitMap["m"]), nil
}
