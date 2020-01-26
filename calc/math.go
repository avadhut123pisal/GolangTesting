package calc

func Add(numbers ...int) int {
	sum := 0
	if len(numbers) < 2 {
		return sum
	} else {
		for _, num := range numbers {
			sum += num
		}
		return sum
	}
}
