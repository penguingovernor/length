package length

import (
	"testing"
)

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

func TestParseDistance(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name    string
		args    args
		want    Distance
		wantErr bool
	}{
		{
			name: "12 Meters",
			args: args{
				s: "12m",
			},
			want:    Distance(12 * Meter),
			wantErr: false,
		},
		{
			name: "0 Meters",
			args: args{
				s: "0m",
			},
			want:    Distance(0),
			wantErr: false,
		},
		{
			name: "Bad Unit",
			args: args{
				s: "0ms",
			},
			want:    Distance(0),
			wantErr: true,
		},
		{
			name: "No Unit",
			args: args{
				s: "12",
			},
			want:    Distance(0),
			wantErr: true,
		},
		{
			name: "Negative Unit",
			args: args{
				s: "-1nm",
			},
			want:    Distance(-1 * Nanometer),
			wantErr: false,
		},
		{
			name: "Partial Unit",
			args: args{
				s: "-1.9mi",
			},
			want:    Distance(-1.9 * Mile),
			wantErr: false,
		},
		{
			name: "No Digits",
			args: args{
				s: "-.ly",
			},
			want:    0,
			wantErr: true,
		},
		{
			name: "Multiple Units",
			args: args{
				s: "5ft11in",
			},
			want:    Distance(5*Feet) + Distance(11*Inch),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseDistance(tt.args.s)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseDistance() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ParseDistance() = %v, want %v", got, tt.want)
			}
		})
	}
}
