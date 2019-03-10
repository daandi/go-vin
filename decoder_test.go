package main

import (
	"testing"
)

func TestDecodeVINWMC(t *testing.T) {
	cases := []struct {
		in   string
		want VIN
	}{
		{"5GZCZ43D13S812715 ", VIN{VIN: "5GZCZ43D13S812715", Manufacturer: Manufacturer{"", "", "5GZ"}}},
		{"5GZCZ43D13S812715", VIN{VIN: "5GZCZ43D13S812715", Manufacturer: Manufacturer{"", "", "5GZ"}}},
	}
	for pos, c := range cases {
		vin := DecodeVIN(c.in)
		if vin != c.want {
			t.Errorf("DecodeVIN(%s) = %v, want %v case %d", c.in, vin, c.want, pos)
		}
	}
}
