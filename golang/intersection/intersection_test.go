package intersection_test

import (
	"fmt"
	"github.com/filatkinen/sorting-more-alg/golang/intersection"
	"math/rand"
	"sort"
	"testing"
	"time"
)

var (
	slice1 []int
	slice2 []int
)

func initSlices(l1, l2 int) {
	r1 := rand.New(rand.NewSource(time.Now().UnixNano()))
	slice1 = make([]int, l1)
	slice2 = make([]int, l2)
	for i := 0; i < l1; i++ {
		slice1[i] = r1.Intn(1000)
	}
	for i := 0; i < l2; i++ {
		slice2[i] = r1.Intn(1000)
	}
}

func BenchmarkIntersection(b *testing.B) {
	data := []struct {
		l1 int
		l2 int
	}{
		{100, 100},
		{1000, 1000},
		{5000, 5000},
	}

	for _, v := range data {
		initSlices(v.l1, v.l2)
		b.Run(fmt.Sprintf("%d,%d unsorted", v.l1, v.l2), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = intersection.IntersectionUnSorted(slice1, slice2)
			}
		})
		sort.Ints(slice1)
		sort.Ints(slice2)
		b.Run(fmt.Sprintf("%d,%d sorted", v.l1, v.l2), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = intersection.IntersectionSorted(slice1, slice2)
			}
		})
	}
}
