package length

import (
	"errors"
	"fmt"
)

// A Distance represents a physical distance
// as a float64 nanometer count. The representation limits the
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
//	fmt.Print(length.Distance(meters)*length.Meter) // prints 10m
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
// As a special case, distances less than one
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

// This code was heavily inspired by the functions
// provided in https://golang.org/src/time/format.go .
// Many thanks to them for making this easier on myself.

var errLeadingInt = errors.New("time: bad [0-9]*") // never printed

// leadingInt consumes the leading [0-9]* from s.
func leadingInt(s string) (x float64, rem string, err error) {
	i := 0
	for ; i < len(s); i++ {
		c := s[i]
		if c < '0' || c > '9' {
			break
		}
		if x > (1<<63-1)/10 {
			// overflow
			return 0, "", errLeadingInt
		}
		x = x*10 + float64(int64(c)-'0')
		if x < 0 {
			// overflow
			return 0, "", errLeadingInt
		}
	}
	return x, s[i:], nil
}

// leadingFraction consumes the leading [0-9]* from s.
// It is used only for fractions, so does not return an error on overflow,
// it just stops accumulating precision.
func leadingFraction(s string) (x int64, scale float64, rem string) {
	i := 0
	scale = 1
	overflow := false
	for ; i < len(s); i++ {
		c := s[i]
		if c < '0' || c > '9' {
			break
		}
		if overflow {
			continue
		}
		if x > (1<<63-1)/10 {
			// It's possible for overflow to give a positive number, so take care.
			overflow = true
			continue
		}
		y := x*10 + int64(c) - '0'
		if y < 0 {
			overflow = true
			continue
		}
		x = y
		scale *= 10
	}
	return x, scale, s[i:]
}

// ParseDistance parses a distance string.
// A distance string is a possibly signed sequence of
// decimal numbers, each with optional fraction and a unit suffix,
// such as "300m" or "-1.5ly"
// Valid distance units are "nm", "um" (or "µm"), "mm", "m", "km", "in", "ft", "yd", "mi", "ly".
func ParseDistance(s string) (Distance, error) {

	// [-+]?([0-9]*(\.[0-9]*)?[a-z]+)+
	orig := s
	var d float64
	neg := false

	// Consume [-+]?
	if s != "" {
		c := s[0]
		if c == '-' || c == '+' {
			neg = c == '-'
			s = s[1:]
		}
	}
	// Special case: if all that is left is "0", this is zero.
	if s == "0" {
		return 0, nil
	}
	if s == "" {
		return 0, errors.New("length: invalid distance " + orig)
	}
	for s != "" {
		var (
			v     float64
			f     int64       // integers before, after decimal point
			scale float64 = 1 // value = v + f/scale
		)

		var err error

		// The next character must be [0-9.]
		if !(s[0] == '.' || '0' <= s[0] && s[0] <= '9') {
			return 0, errors.New("length: invalid distance " + orig)
		}
		// Consume [0-9]*
		pl := len(s)
		v, s, err = leadingInt(s)
		if err != nil {
			return 0, errors.New("length: invalid distance " + orig)
		}
		pre := pl != len(s) // whether we consumed anything before a period

		// Consume (\.[0-9]*)?
		post := false
		if s != "" && s[0] == '.' {
			s = s[1:]
			pl := len(s)
			f, scale, s = leadingFraction(s)
			post = pl != len(s)
		}
		if !pre && !post {
			// no digits (e.g. ".s" or "-.s")
			return 0, errors.New("length: invalid distance " + orig)
		}

		// Consume unit.
		i := 0
		for ; i < len(s); i++ {
			c := s[i]
			if c == '.' || '0' <= c && c <= '9' {
				break
			}
		}
		if i == 0 {
			return 0, errors.New("length: missing unit in distance " + orig)
		}
		u := s[:i]
		s = s[i:]
		unit, ok := unitMap[u]
		if !ok {
			return 0, errors.New("length: unknown unit " + u + " in distance " + orig)
		}
		if v > (1<<63-1)/unit {
			// overflow
			return 0, errors.New("length: invalid distance " + orig)
		}
		v *= unit
		if f > 0 {
			v += float64(f) * (float64(unit) / scale)
			if v < 0 {
				// overflow
				return 0, errors.New("length: invalid distance " + orig)
			}
		}
		d += v
		if d < 0 {
			// overflow
			return 0, errors.New("length: invalid distance " + orig)
		}
	}

	if neg {
		d = -d
	}
	return Distance(d), nil
}
