package main

func Add[T int | float32](a, b T) T {
	return a + b
}

// 1. 匿名函数不能自己定义类型形参
//func test() {
//	fn1 := func[T int | int32](a, b T) T {
//		return a + b
//	}
//}

// 2. 匿名函数可以使用定义好的类型形参
func test[T int | int32](a, b T) T {
	result := func(a, b T) T {
		return a + b
	}(a, b)
	return result
}
