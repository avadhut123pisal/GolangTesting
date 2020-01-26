package calc

import (
	"errors"
	"github.com/fatih/color"
)

func Add(numbers ...int) (int, error) {
	sum := 0
	if len(numbers) < 2 {
		return sum, errors.New(color.RedString("Provide more than two numbers"))
	} else {
		for _, num := range numbers {
			sum += num
		}
		return sum, nil
	}
}
