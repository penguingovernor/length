package length

import "testing"

func TestDistance_String(t *testing.T) {
	type beforeFunc func()
	tests := []struct {
		name   string
		d      Distance
		want   string
		before beforeFunc
	}{
		{
			name:   "Metric - 0 meters",
			d:      Distance(0),
			want:   "0m",
			before: func() { UseMetric() },
		},
		{
			name:   "Metric - 2 Nanometers",
			d:      Distance(2 * Nanometer),
			want:   "2.000000nm",
			before: func() { UseMetric() },
		},
		{
			name:   "Metric - 2 Micrometers",
			d:      Distance(2 * Micrometer),
			want:   "2.000000µm",
			before: func() { UseMetric() },
		},
		{
			name:   "Metric - 2 Millimeters",
			d:      Distance(2 * Millimeter),
			want:   "2.000000mm",
			before: func() { UseMetric() },
		},
		{
			name:   "Metric - 2 Centimeters",
			d:      Distance(2 * Centimeter),
			want:   "2.000000cm",
			before: func() { UseMetric() },
		},
		{
			name:   "Metric - 2 Meters",
			d:      Distance(2 * Meter),
			want:   "2.000000m",
			before: func() { UseMetric() },
		},
		{
			name:   "Metric - 2 Kilometers",
			d:      Distance(2 * Kilometer),
			want:   "2000.000000m",
			before: func() { UseMetric() },
		},
		{
			name:   "Metric - 2 Lightyears",
			d:      Distance(2 * Lightyear),
			want:   "18922000000000000.000000m",
			before: func() { UseMetric() },
		},
		{
			name: "Imperial - 0 yd",
			d:    Distance(0),
			want: "0yd",
			before: func() {
				UseMetric()
				ToggleUnits()
			},
		},
		{
			name:   "Imperial - 2 in",
			d:      Distance(2 * Inch),
			want:   "2.000000in",
			before: func() { UseImperial() },
		},
		{
			name:   "Imperial - 2 ft",
			d:      Distance(2 * Feet),
			want:   "2.000000ft",
			before: func() { UseImperial() },
		},
		{
			name:   "Imperial - 2 yd",
			d:      Distance(2 * Yard),
			want:   "2.000000yd",
			before: func() { UseImperial() },
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.before()
			if got := tt.d.String(); got != tt.want {
				t.Errorf("Distance.String() = %v, want %v", got, tt.want)
			}
		})
	}
}
