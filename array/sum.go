package main

func sum(numbers [5]int) (sum int) {
	for i := 0; i < 5; i++ {
		sum = sum + numbers[i]
	}
	return
}
