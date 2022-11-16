package utils

func CalculateGas(gasUnit, gasPrice int64) (float64, error) {
	return float64(gasUnit*gasPrice) / 1000000000, nil
}
