package conversion

import (
	"errors"
	"strconv"
)

func StringsToFloat (s []string) ([]float64, error) {
	floats := make([]float64, len(s))
	for index, value := range s {
		float, err := strconv.ParseFloat(value, 64)
		if err != nil {
			return []float64{}, errors.New("failed to parse float")
		}
		floats[index] = float
	}
	return floats, nil
}