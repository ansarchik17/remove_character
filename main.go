package main

import (
	"fmt"
	"strconv"
)

func main() {
	arr := []any{9, "2", 10, "3"}
	result := SumMix(arr)
	fmt.Println(result)
}

func SumMix(arr []any) int {
	sum := 0

	for _, value := range arr {
		switch v := value.(type) {
		case int:
			sum += v
		case string:
			num, err := strconv.Atoi(v)
			if err == nil {
				sum += num
			}
		}
	}

	return sum
}
