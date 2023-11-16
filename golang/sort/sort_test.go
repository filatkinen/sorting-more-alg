package sort_test

import (
	"fmt"
	"github.com/filatkinen/sorting-more-alg/golang/sort"
	"math/rand"
	"testing"
	"time"
)

var (
	slice1 []int
	slice2 []int
)

func initSlices(l1 int) {
	r1 := rand.New(rand.NewSource(time.Now().UnixNano()))
	slice1 = make([]int, l1)
	slice2 = make([]int, l1)
	for i := 0; i < l1; i++ {
		slice1[i] = r1.Intn(1000)
	}
	copy(slice2, slice1)
}

func BenchmarkSort(b *testing.B) {
	data := []struct {
		l1 int
	}{
		{100},
		{1000},
		{10000},
	}

	for _, v := range data {
		initSlices(v.l1)
		b.Run(fmt.Sprintf("%d sort", v.l1), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				sort.QuickSort(slice1)
			}
		})
		b.Run(fmt.Sprintf("%d sort append", v.l1), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				sort.QuickSortAppend(slice2)
			}
		})
	}
}
