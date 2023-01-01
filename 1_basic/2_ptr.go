package main

import (
	"fmt"
	"reflect"
)

func Int32Ptr(in int32) *int32 {
	return &in
}

func StringPtr(in string) *string {
	return &in
}

// 获取指针
func Ptr[T any](in T) *T {
	return &in
}

func Int32Value(in *int32) int32 {
	return *in
}

func StringValue(in *string) string {
	return *in
}

// 获取指针的值
func PtrValue[T any](in *T) T {
	return *in
}

func main() {
	int32Ptr1 := Int32Ptr(1)
	fmt.Println(reflect.TypeOf(int32Ptr1)) // *int32
	stringPtr1 := StringPtr("a")
	fmt.Println(reflect.TypeOf(stringPtr1)) // *string
	// 使用泛型
	intPtr1 := Ptr[int](1)
	fmt.Println(reflect.TypeOf(intPtr1)) // *int
	intPtr2 := Ptr(1)                    // 类型推断
	fmt.Println(reflect.TypeOf(intPtr2)) // *int

	strPtr := Ptr("a")
	fmt.Println(reflect.TypeOf(strPtr)) // *string
	strValue1 := StringValue(strPtr)
	fmt.Println(reflect.TypeOf(strValue1)) // string
	// 泛型
	strValue2 := PtrValue(strPtr)
	fmt.Println(reflect.TypeOf(strValue2)) // string
}
