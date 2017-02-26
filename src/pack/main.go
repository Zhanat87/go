package pack

func Add(numbers ...int) int {
	var result int
	if len(numbers) == 0 {
		println("No arguments provided")
		return 0
	}
	for _, num := range numbers {
		result += num
	}
	return result
}

func Subtract(initial int, numbers ...int) int {
	if len(numbers) == 0 {
		println("No arguments provided")
		return initial
	}
	for _, num := range numbers {
		initial -= num
	}
	return initial
}
