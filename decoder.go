package main

import "strings"

type VIN struct {
	VIN          string
	Manufacturer Manufacturer
}

type Manufacturer struct {
	Name    string
	Country string
	Code    string
}

func DecodeVIN(vin_str string) VIN {
	noramlized_vin_str := strings.ToUpper(strings.TrimSpace(vin_str))
	manunfacturer := decodeManufacturer(noramlized_vin_str)

	return VIN{VIN: noramlized_vin_str, Manufacturer: manunfacturer}
}

func decodeManufacturer(vin string) Manufacturer {
	wmc := extractWorldManufacturerCode(vin)

	return Manufacturer{Name: "", Country: "", Code: wmc}
}

func extractWorldManufacturerCode(vin string) string {
	return vin[:3]
}
