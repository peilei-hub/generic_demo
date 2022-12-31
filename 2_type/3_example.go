package main

import "fmt"

type Set[K comparable] map[K]struct{}

func NewSet[K comparable](keys ...K) Set[K] {
	set := make(Set[K])
	for _, key := range keys {
		set.Add(key)
	}
	return set
}

func (s *Set[K]) Add(key K) bool {
	if _, ok := (*s)[key]; ok {
		return false
	}
	(*s)[key] = struct{}{}
	return true
}

func (s *Set[K]) Remove(key K) bool {
	if _, ok := (*s)[key]; ok {
		delete(*s, key)
		return true
	}
	return false
}

func (s *Set[K]) Keys() []K {
	res := make([]K, 0)
	for k := range *s {
		res = append(res, k)
	}
	return res
}

func (s *Set[K]) Contains(key K) bool {
	_, ok := (*s)[key]
	return ok
}

func main() {
	strSet := NewSet("c") // 类型推断
	strSet.Add("a")
	strSet.Add("a")
	strSet.Add("b")
	strKeys := strSet.Keys()
	fmt.Println(strKeys)

	intSet := NewSet(1, 2)
	intSet.Add(1)
	intSet.Add(1)
	intSet.Add(3)
	intKeys := intSet.Keys()
	fmt.Println(intKeys)
	fmt.Println(intSet.Contains(2))
}
