package stringutil

func Float2Str(f float64) string {
	if f < 0 {
		return "-" + Float2Str(-f)
	}
	if f == 0 {
		return "0"
	}

	var result string
	intPart := int(f)
	fractionalPart := f - float64(intPart)

	// Convert integer part to string
	for intPart > 0 {
		digit := byte('0' + intPart%10)
		result = string(digit) + result
		intPart /= 10
	}

	if fractionalPart > 0 {
		result += "."
		for fractionalPart > 0 {
			fractionalPart *= 10
			digit := byte('0' + int(fractionalPart))
			result += string(digit)
			fractionalPart -= float64(int(fractionalPart))
		}
	}

	return result
}
