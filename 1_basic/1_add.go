package main

type IntSlice []int
type Int32Slice []int32
type Float32Slice []float32
type MySlice[T int | int32 | float32] []T

func AddInt(a, b int) int {
	return a + b
}

func AddInt32(a, b int32) int32 {
	return a + b
}

func Add[T int | int32 | string](a, b T) T {
	return a + b
}

func main() {
	var intA int = 1
	var intB int = 2
	var int32A int32 = 3
	var int32B int32 = 4
	// go 1.18之前的使用
	AddInt(intA, intB)
	AddInt32(int32A, int32B)

	// 泛型
	Add(intA, intB) // 类型推断
	Add[int32](int32A, int32B)
	Add("a", "b") // 类型推断
}
