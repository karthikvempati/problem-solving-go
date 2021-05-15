package recursion

import "fmt"

func main() {
	fmt.Println(Factorial(9))
	fmt.Println(Factorial(20))
}

func Factorial(number int64) int64 {
	if number == 0 {
		return 1
	}
	return number * Factorial(number-1)
}
