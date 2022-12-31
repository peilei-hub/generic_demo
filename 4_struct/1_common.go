package main

import "fmt"

// 1. 举个例子
type Map[K comparable, V any] struct {
	Data map[K]V
}

func NewMap[K comparable, V any]() *Map[K, V] {
	return &Map[K, V]{
		Data: make(map[K]V),
	}
}

func (m *Map[K, V]) Set(key K, value V) {
	m.Data[key] = value
}

func (m *Map[K, V]) Get(key K) V {
	return m.Data[key]
}

func (m *Map[K, V]) Exist(key K) bool {
	_, ok := m.Data[key]
	return ok
}

func (m *Map[K, V]) PrintAll() {
	for k, v := range m.Data {
		fmt.Println("key: ", k, ", val: ", v)
	}
}

// 2.1 不支持泛型方法
//
//	func (m *Map[K, V]) TestGeneric[T int | string](a, b T) T {
//		return a + b
//	}
//

// 2.2 只能通过receiver使用类型形参
func (m *Map[K, V]) Equal(a, b K) bool {
	return a == b
}

type Student struct {
	Num  int
	Name string
}

func main() {
	intStringMap := NewMap[int, string]()
	intStringMap.Set(1, "a")
	intStringMap.Set(2, "b")
	intStringMap.PrintAll()

	s1 := &Student{
		Num:  1,
		Name: "a",
	}
	s2 := &Student{
		Num:  2,
		Name: "b",
	}
	numStudentMap := NewMap[int, *Student]()
	numStudentMap.Set(s1.Num, s1)
	numStudentMap.Set(s2.Num, s2)
	numStudentMap.PrintAll()

	// 3. 匿名结构体不支持泛型
	//testCase := struct [T int | string] {
	//	a T
	//}[int] {
	//	a: 1
	//}
}
