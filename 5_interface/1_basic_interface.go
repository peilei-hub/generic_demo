package main

// 1. 基本接口，方法集，go 1.18之前的接口的定义

// 1.1 基本接口-无泛型
type BasicInterface interface {
	Name() string
	Age() int
}

// 可以定义变量
var a BasicInterface

// 基本接口也代表一个类型集，可以用在类型约束中
type ATest[T BasicInterface] []T

// 可以当做类型集,用在泛型方法
func BasicInterfaceFunc1[T BasicInterface](b T) {
	b.Name()
	b.Age()
}

// 可以跟go1.18之前写法一致
func BasicInterfaceFunc2(b BasicInterface) {
	b.Name()
	b.Age()
}

// 1.2 基本接口-有泛型

type BasicInterface2[T int | int32 | string] interface {
	Func1(in T) (out T)
	Func2() T
}

// 可以定义变量
var b1 BasicInterface2[int]
var b2 BasicInterface2[string]

// 也可以用在类型约束
type BTest1[T BasicInterface2[int]] []T
type BTest2[T int] BasicInterface2[T]

// 可以当做类型集,用在泛型方法
func BTestFunc1[T BasicInterface2[int]](t T) {
	t.Func2()
}

// 可以跟go1.18之前写法一致
func BTestFunc2(t BasicInterface2[int]) {
	t.Func2()
}

// BasicInterface2的接口实现举例
// 只要满足三个条件：
// 1. 有方法  Func1(in T) (out T), 方法 Func2() T。
// 2. T满足约束为 int | int32 | string，
// 3. 方法的T,同时只能为一种类型。比如Func1中的T为int, Func2中的T也只能为int

// 举例1-是BasicInterface2的实现
type BasicInterface2Impl struct{}

func (b BasicInterface2Impl) Func1(in int) (out int) {
	panic("implement me")
}
func (b BasicInterface2Impl) Func2() int {
	panic("implement me")
}

// 举例2-不是BasicInterface2的实现
type BasicInterface2Impl2 struct{}

func (b BasicInterface2Impl2) Func1(in string) (out string) {
	panic("implement me")
}

func (b BasicInterface2Impl2) Func2() int {
	panic("implement me")
}

// 举例3-不是BasicInterface2的实现
type BasicInterface2Impl3 struct{}

func (b BasicInterface2Impl3) Func1(in float32) (out float32) {
	panic("implement me")
}

func (b BasicInterface2Impl3) Func2() float32 {
	panic("implement me")
}

func main() {
	var basicInterface2 BasicInterface2[int] = &BasicInterface2Impl{}
	_ = basicInterface2
}
