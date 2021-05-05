package factorial

import "fmt"

func main() {
	fmt.Println(factorial(9))
	fmt.Println(factorial(20))
}

func factorial(number int64) int64 {
	if number == 0 {
		return 1
	}
	return number * factorial(number-1)
}
