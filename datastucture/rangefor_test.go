package datastucture

import (
	"fmt"
	"testing"
)

type Item struct {
	id int
	val [4096]byte
}

func BenchmarkForStruct(b *testing.B)  {
	var items [1024]Item
	for i := 0; i < b.N; i++ {
		length := len(items)
		var tmp int
		for k := 0; k < length; k++ {
			tmp = items[k].id
		}
		_ = tmp
	}
}

func BenchmarkRangeIndexStruct(b *testing.B)  {
	var items [1024]Item
	for i := 0; i < b.N; i++ {
		var tmp int
		for k := range items {
			tmp = items[k].id
		}
		_ = tmp
	}
}

func BenchmarkRangeStruct(b *testing.B)  {
	var items [1024]Item
	for i := 0; i < b.N; i++ {
		var tmp int
		for _, item := range items {
			tmp = item.id
		}
		_ = tmp
	}
}

func Test_RangeCopy(t *testing.T)  {
	persons := []struct{no int} {{no: 1}, {no: 2}, {no: 3}}
	// range迭代时，返回的是copy
	for _, person := range persons {
		person.no += 10
	}

	for i := 0; i < len(persons); i++ {
		persons[i].no += 100
	}

	fmt.Println(persons)
}

