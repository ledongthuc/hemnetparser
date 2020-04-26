package hemnetparser

import (
	"testing"
)

func TestPtrFloat64(t *testing.T) {
	valuePositive := float64(10)
	valueNegative := float64(-10)
	tests := []struct {
		name string
		val  float64
		want *float64
	}{
		{name: "Empty", val: 0, want: new(float64)},
		{name: "Positive", val: 10, want: &valuePositive},
		{name: "Negative", val: -10, want: &valueNegative},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := PtrFloat64(tt.val)
			if got != tt.want && (got == nil || tt.want == nil) {
				t.Errorf("PtrFloat64() = %v, want %v", got, tt.want)
				return
			}

			if got == nil && tt.want == nil {
				return
			}

			if *got != *tt.want {
				t.Errorf("PtrFloat64() = %v, want %v", *got, *tt.want)
			}
		})
	}
}

func TestPtrString(t *testing.T) {
	value := string("Value")
	tests := []struct {
		name string
		val  string
		want *string
	}{
		{name: "Empty", val: "", want: new(string)},
		{name: "Value", val: value, want: &value},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := PtrString(tt.val)
			if got != tt.want && (got == nil || tt.want == nil) {
				t.Errorf("PtrString() = %v, want %v", got, tt.want)
				return
			}

			if got == nil && tt.want == nil {
				return
			}

			if *got != *tt.want {
				t.Errorf("PtrString() = %v, want %v", *got, *tt.want)
			}
		})
	}
}

func TestGetFloat64(t *testing.T) {
	positive := float64(10)
	negative := float64(-10)
	tests := []struct {
		name string
		val  *float64
		want float64
	}{
		{name: "Empty", val: new(float64), want: 0},
		{name: "Nil", val: nil, want: 0},
		{name: "Positive", val: &positive, want: 10},
		{name: "Negative", val: &negative, want: -10},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetFloat64(tt.val); got != tt.want {
				t.Errorf("GetFloat64() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetString(t *testing.T) {
	value := string("Value")
	tests := []struct {
		name string
		val  *string
		want string
	}{
		{name: "Empty", val: new(string), want: ""},
		{name: "Nil", val: nil, want: ""},
		{name: "Value", val: &value, want: "Value"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetString(tt.val); got != tt.want {
				t.Errorf("GetString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStringValueOrNil(t *testing.T) {
	value := "Value"
	tests := []struct {
		name string
		val  string
		want *string
	}{
		{name: "Empty", val: "", want: nil},
		{name: "Value", val: "Value", want: &value},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := StringValueOrNil(tt.val)
			if got != tt.want && (got == nil || tt.want == nil) {
				t.Errorf("StringValueOrNil() = %v, want %v", got, tt.want)
				return
			}

			if got == nil && tt.want == nil {
				return
			}

			if *got != *tt.want {
				t.Errorf("StringValueOrNil() = %v, want %v", *got, *tt.want)
			}
		})
	}
}
