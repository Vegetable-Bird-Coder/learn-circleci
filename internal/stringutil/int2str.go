package stringutil

func Int2Str(i int) string {
	if i < 0 {
		return "-" + Int2Str(-i)
	}
	if i == 0 {
		return "0"
	}

	var result string
	for i > 0 {
		digit := byte('0' + i%10)
		result = string(digit) + result
		i /= 10
	}
	return result
}
