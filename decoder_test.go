package main

import (
	"testing"
)

func TestDecodeVIN(t *testing.T) {
	cases := []struct {
		in   string
		want VIN
	}{
		{"5GZCZ43D13S812715", VIN{VIN: "5GZCZ43D13S812715", WMI: "5GZ", VDS: "CZ43D1", VIS: "3S812715"}},
		{"    5gzcZ43d13S812715  ", VIN{"5GZCZ43D13S812715", "5GZ", "CZ43D1", "3S812715"}},
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
		in  string
		wmi string
		vds string
		vis string
	}{
		{"SB164ABN1PE082986", "SB1", "64ABN1", "PE082986"},
		{"WVWZZZ1JZ3W386752", "WVW", "ZZZ1JZ", "3W386752"},
	}
	for pos, c := range cases {
		wmi, vds, vis := SplitVIN(c.in)
		if wmi != c.wmi || vds != c.vds || vis != c.vis {
			t.Errorf("SplitVIN(%s), want %s %s %s but got %s %s %s at case %d",
				c.in, c.wmi, c.vds, c.vis, wmi, vis, vds, pos)
		}
	}
}
