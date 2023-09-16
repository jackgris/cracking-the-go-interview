package main

import (
	"math/rand"
	"strconv"
	"testing"
)

func TestMinSum(t *testing.T) {
	tdt := []struct {
		num    []int32
		k      int32
		result int32
	}{
		{
			num:    []int32{8, 7, 6},
			k:      4,
			result: 2 + 4 + 3,
		},
		{
			num:    []int32{10, 6, 8},
			k:      3,
			result: 5 + 4 + 3,
		},
		{
			num:    []int32{12, 13, 28},
			k:      2,
			result: 7 + 13 + 12,
		},
		{
			num:    []int32{30, 10, 6, 2},
			k:      4,
			result: 4 + 5 + 6 + 2,
		},
	}

	for i, v := range tdt {
		t.Run(strconv.Itoa(i+1), func(t *testing.T) {
			result := minSum(v.num, v.k)
			if result != v.result {
				t.Fatalf(`expected "%v" not "%v"`, v.result, result)
			}
		})
	}
}

func BenchmarkMinSum(b *testing.B) {
	var (
		k   int32 = 10_000
		num       = make([]int32, 100_000)
	)

	for i := range num {
		num[i] = int32(rand.Intn(1_000))
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		minSum(num, k)
	}
}
