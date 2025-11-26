package helper

import "math"

func BasicArithmetic(a int, b int) (int, int, int, int) {
	addition := a + b
	subtract := a - b
	multiplication := subtract * b
	division := a / b
	return addition, subtract, multiplication, division
}

func C2FConvertor(temp float32) float32 {
	return temp*9/5 + 32
}

func CircleAreaCalculator(radius float32) float32 {
	return radius * radius * math.Pi
}
