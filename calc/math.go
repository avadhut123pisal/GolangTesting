package calc

import (
	"errors"
)

func Add(numbers ...int) (int, error) {
	sum := 0
	if len(numbers) < 2 {
		return sum, errors.New("Provide more than two numbers")
	} else {
		for _, num := range numbers {
			sum += num
		}
		return sum, nil
	}
}
