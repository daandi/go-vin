package main

import "strings"

type VIN struct {
	VIN string
	WMI string
	VDS string
	VIS string
}

type Manufacturer struct {
	Name    string
	Country string
	Code    string
}

func DecodeVIN(vin_str string) VIN {
	normalizedVINStr := strings.ToUpper(strings.TrimSpace(vin_str))
	wmi, vds, vis := SplitVIN(normalizedVINStr)

	return VIN{VIN: normalizedVINStr, WMI: wmi, VDS: vds, VIS: vis}
}

func DecodeWMI(wmi string) Manufacturer {
	return Manufacturer{Name: "", Country: "", Code: wmi}
}

/*SplitVIN splits a vin in wmi vds and vis*/
func SplitVIN(vin string) (wmi, vds, vis string) {
	wmi = vin[:3]
	vds = vin[3:9]
	vis = vin[9:17]
	return
}
