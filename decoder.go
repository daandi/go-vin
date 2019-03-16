package vingo

import "strings"

type VIN struct {
	VIN string
	vinData
}

type vinData struct {
	WMI, VDS, VIS string
}

type Manufacturer struct {
	Name, Country, Code string
}

func DecodeVIN(vin_str string) VIN {
	normalizedVINStr := strings.ToUpper(strings.TrimSpace(vin_str))
	vinData := SplitVIN(normalizedVINStr)

	return VIN{normalizedVINStr, vinData}
}

func DecodeWMI(wmi string) Manufacturer {
	return Manufacturer{Name: "", Country: "", Code: wmi}
}

/*SplitVIN splits a vin in wmi vds and vis*/
func SplitVIN(vin string) vinData {
	return vinData{WMI: vin[:3], VDS: vin[3:9], VIS: vin[9:17]}
}
