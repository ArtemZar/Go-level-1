package operations

import (

	"testing"

)

func TestMultiply(t *testing.T) {
	cases := []struct{
		name string
		number1 float64
		number2 float64
		want float64
	}{
		{name: "Целые положительные",
			number1: 5,
			number2: 4,
			want:    20,
		},
		{name: "Одно отрицательное",
			number1: -3,
			number2: 5,
			want:    -15,
		},
		{name: "Дробные числа",
			number1: 1.3,
			number2: 1.2,
			want:    1.56,
		},
	}

	for _, tc := range cases {
		have, _ := Multiply(tc.number1, tc.number2)
		if !close(have, tc.want) {
			t.Errorf("%s: ожидания: %v, факт: %v", tc.name, tc.want, have)
		}
	}
}

func TestSum(t *testing.T) {
	cases := []struct{
		name string
		number1 float64
		number2 float64
		want float64
	}{
		{name: "Целые положительные",
			number1: 5,
			number2: 4,
			want:    9,
		},
		{name: "Одно отрицательное",
			number1: -3,
			number2: 5,
			want:    2,
		},
		{name: "Дробные числа",
			number1: 1.3,
			number2: 1.2,
			want:    2.5,
		},
	}

	for _, tc := range cases {
		have, _ := Sum(tc.number1, tc.number2)
		if !close(have, tc.want) {
			t.Errorf("%s: ожидания: %v, факт: %v", tc.name, tc.want, have)
		}
	}
}

func TestSubtract(t *testing.T) {
	cases := []struct{
		name string
		number1 float64
		number2 float64
		want float64
	}{
		{name: "Целые положительные",
			number1: 5,
			number2: 4,
			want:    1,
		},
		{name: "Одно отрицательное",
			number1: -3,
			number2: 5,
			want:    -8,
		},
		{name: "Дробные числа",
			number1: 1.3,
			number2: 1.2,
			want:    0.1,
		},
	}

	for _, tc := range cases {
		have, _ := Subtract(tc.number1, tc.number2)
		if !close(have, tc.want) {
			t.Errorf("%s: ожидания: %v, факт: %v", tc.name, tc.want, have)
		}
	}
}

func TestDivide(t *testing.T) {
	cases := []struct{
		name string
		number1 float64
		number2 float64
		want float64
	}{
		{name: "Целые положительные",
			number1: 5,
			number2: 4,
			want:    1.25,
		},
		{name: "Одно отрицательное",
			number1: -3,
			number2: 5,
			want:    -0.6,
		},
		{name: "Дробные числа",
			number1: 1.8,
			number2: 1.2,
			want:    1.5,
		},
	}

	for _, tc := range cases {
		have, _ := Divide(tc.number1, tc.number2)
		if !close(have, tc.want) {
			t.Errorf("%s: ожидания: %v, факт: %v", tc.name, tc.want, have)
		}
	}
}

func TestElevate(t *testing.T) {
	cases := []struct{
		name string
		number1 float64
		number2 float64
		want float64
	}{
		{name: "2 в квадрате",
			number1: 2,
			number2: 2,
			want:    4,
		},
		{name: "2 в кубе",
			number1: 2,
			number2: 3,
			want:    8,
		},
		{name: "5 в 4 степени",
			number1: 5,
			number2: 4,
			want:    625,
		},
	}

	for _, tc := range cases {
		have, _ := Elevate(tc.number1, tc.number2)
		if !close(have, tc.want) {
			t.Errorf("%s: ожидания: %v, факт: %v", tc.name, tc.want, have)
		}
	}
}

func tolerance(a, b, e float64) bool {
	// Multiplying by e here can underflow denormal values to zero.
	// Check a==b so that at least if a and b are small and identical
	// we say they match.
	if a == b {
		return true
	}
	d := a - b
	if d < 0 {
		d = -d
	}

	// note: b is correct (expected) value, a is actual value.
	// make error tolerance a fraction of b, not a.
	if b != 0 {
		e = e * b
		if e < 0 {
			e = -e
		}
	}
	return d < e
}
func close(a, b float64) bool      { return tolerance(a, b, 1e-7) }