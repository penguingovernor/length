package length

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
