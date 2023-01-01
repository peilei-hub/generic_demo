package main

import (
	"fmt"
	"reflect"
	"strconv"
)

// 2. 一般接口，接口中包含类型约束，不能用来定义变量

// 2.1 简单类型约束
type CommonInterface interface {
	int | int8 | float32 | string
}

// 不能用来定义变量
//var commonInterface CommonInterface // error

// 2.2 集合操作
type Int interface {
	int | int8 | int32 | int64
}

type Float interface {
	float32 | float64
}

// 并集操作
type IntAndFloat interface {
	Int | Float
}

// 交集操作
type IntExceptInt8 interface {
	Int
	int8
}

// 空集
type Null interface {
	int
	int32
}

// 2.3 复杂类型约束1
type CommonInterface2 interface {
	~int | ~int8 | ~struct {
		Data string
	}

	Func1() string
}

// 不能用来定义变量
//var c CommonInterface2 // error

func DoCommonInterface2[T CommonInterface2](t T) {
	fmt.Println(t.Func1())
}

// 实现复杂类型约束1，需要满足：
// 1. 底层类型为int|int8|struct{Data string}
// 2. 有方法Func1() string

// 举例1 是CommonInterface2的实例化
type CommonInterface2_1 int

func (c CommonInterface2_1) Func1() string {
	return "CommonInterface2_1"
}

// 举例2 是CommonInterface2的实例化
type CommonInterface2_2 struct {
	Data string
}

func (c CommonInterface2_2) Func1() string {
	return c.Data
}

// 举例3 不是CommonInterface2的实例化
type CommonInterface2_3 int32

func (c CommonInterface2_3) Func1() string {
	panic("CommonInterface2_3")
}

// 2.4 复杂类型约束2
type CommonInterface3[T string | float32] interface {
	~int | ~int8 | ~struct {
		Data T
	}

	Func2() T
}

// 不能用来定义变量
//var a CommonInterface3[string]

func DoCommonInterface3_1[T CommonInterface3[string]](t T) {
	fmt.Println(reflect.TypeOf(t.Func2()))
}

func DoCommonInterface3_2[T CommonInterface3[float32]](t T) {
	fmt.Println(reflect.TypeOf(t.Func2()))
}

// 新增一个泛型D, 用来表示CommonInterface3里的泛型
func DoCommonInterface3[D string | float32, T CommonInterface3[D]](t T) {
	fmt.Println(reflect.TypeOf(t.Func2()))
}

// 实现复杂类型约束2，需要满足：
// 1. 底层类型为int|int8|struct{Data T} ，T的约束为string | float32
// 2. 有方法Func2() T, T的约束为string | float32
// 3. 如果底层类型定义的没有T, 则Func2() T的约束可以为string | float32
// 4. 如果底层类型定义的有T, 比如struct{Data T}, 此时Func2() T，两个T需要同时为string或者同时为float32

// 举例1 CommonInterface3的实例化
type CommonInterface3Impl1 int

func (c CommonInterface3Impl1) Func2() string {
	return strconv.Itoa(int(c))
}

// 举例2 CommonInterface3的实例化
type CommonInterface3Impl2 int

func (c CommonInterface3Impl2) Func2() float32 {
	return float32(c)
}

// 举例3 CommonInterface3的实例化
type CommonInterface3Impl3[T string | float32] struct {
	Data T
}

func (c CommonInterface3Impl3[T]) Func2() T {
	return c.Data
}

func main() {
	// test CommonInterface2
	commonInterface2_1 := CommonInterface2_1(1)
	DoCommonInterface2[CommonInterface2_1](commonInterface2_1)
	DoCommonInterface2(commonInterface2_1) // 类型推断
	commonInterface2_2 := CommonInterface2_2{}
	DoCommonInterface2(commonInterface2_2)
	//commonInterface2_3 := CommonInterface2_3(1)
	//DoCommonInterface2(commonInterface2_3) // error, commonInterface2_3不是CommonInterface2的实例化

	// test CommonInterface3
	commonInterface3_1 := CommonInterface3Impl1(1)
	DoCommonInterface3_1(commonInterface3_1)
	DoCommonInterface3[string](commonInterface3_1)

	commonInterface3_2 := CommonInterface3Impl2(1)
	DoCommonInterface3_2(commonInterface3_2)
	DoCommonInterface3[float32](commonInterface3_2)

	commonInterface3_3 := CommonInterface3Impl3[string]{
		Data: "data",
	}
	DoCommonInterface3_1(commonInterface3_3)
	DoCommonInterface3[string](commonInterface3_3)
}
