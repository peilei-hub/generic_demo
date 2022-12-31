package main

import "reflect"

// 1. 类型形参不能单独使用
//type CommonType[T int | string | float32] T // error

// 2. 指针类型约束问题
// type Test1 [T * int] []T // error 会当做 T 乘 int
// 可以用逗号消除歧义
type Test2[T *int,] []T
type Test3[T *int | *int32,] []T

// 推荐写法
type Test4[T interface{ *int | *int32 }] []T

// 3. 类型套娃
type Slice[T int | int32 | float32 | float64] []T

//type StringSlice[T string] Slice[T] // error, Slice[T] 的T不支持string

type FloatSlice[T float32 | float64] Slice[T]

// 两种引用Slice[T]的写法
type MapSlice1[T int | float32] map[string]Slice[T]
type MapSlice2[T Slice[int] | Slice[float32]] map[string]T

// 4. 特殊泛型, 表示实际类型还是int
type Special[T int | string] int

// 5. 获取具体类型
func GetType[T int | string](t T) {
	// t.(int) // error，泛型类型定义的变量不能使用类型断言

	// 1. 反射
	v := reflect.ValueOf(t)
	switch v.Kind() {
	case reflect.Int:
	default:
	}

	// 2. 转换
	var i any = t
	switch i.(type) {
	case int:
	default:
	}
}

func main() {
	// 6. 特殊泛型测试
	var s1 Special[int] = 1
	var s2 Special[string] = 2
	_ = s1
	_ = s2
	//var s3 Special[string] = "3" // error, Special[T]的实际类型为int

	// 7. 匿名结构体不支持泛型
	//testCase := struct [T int | string] {
	//	a T
	//}[int] {
	//	a: 1
	//}
}
