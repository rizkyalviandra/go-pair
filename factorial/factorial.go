package factorial

func Factorial(number int) int {
	result := number
		if result != 0 {
			for i:=number-1;i>0;i-- {
				if (i-1) != 0 {
					result *= i
				}
			}
		} else if result == 0 {
			result = 1
		} else {
			return result
		}
	return result
}