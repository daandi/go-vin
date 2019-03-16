package vingo

import (
	"testing"
)

func TestDecodeVIN(t *testing.T) {
	cases := []struct {
		in   string
		want VIN
	}{
		{"5GZCZ43D13S812715", VIN{"5GZCZ43D13S812715", vinData{"5GZ", "CZ43D1", "3S812715"}}},
		{"    5gzcZ43d13S812715  ", VIN{"5GZCZ43D13S812715", vinData{"5GZ", "CZ43D1", "3S812715"}}},
	}
	for pos, c := range cases {
		vin := DecodeVIN(c.in)
		if vin != c.want {
			t.Errorf("DecodeVIN(%s) = %v, but want %v case %d", c.in, vin, c.want, pos)
		}
	}
}

func TestSplitVIN(t *testing.T) {
	cases := []struct {
		vin     string
		vinData vinData
	}{
		{"SB164ABN1PE082986", vinData{"SB1", "64ABN1", "PE082986"}},
		{"WVWZZZ1JZ3W386752", vinData{"WVW", "ZZZ1JZ", "3W386752"}},
	}
	for pos, want := range cases {
		vinData := SplitVIN(want.vin)
		if vinData != want.vinData {
			t.Errorf("SplitVIN(%s), wanted: %v but got: %v case %d",
				want.vin, want.vinData, vinData, pos)
		}
	}
}
