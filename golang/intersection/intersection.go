package intersection

func IntersectionSorted(a, b []int) []int {
	curA, curB := 0, 0
	c := make([]int, 0, min(len(a), len(b)))

	for {
		if a[curA] < b[curB] {
			curA++
			if curA == len(a) {
				break
			}
			continue
		}
		if a[curA] > b[curB] {
			curB++
			if curB == len(b) {
				break
			}
			continue
		}
		if a[curA] == b[curB] {
			if len(c) == 0 || c[len(c)-1] != a[curA] {
				c = append(c, a[curA])
				curA++
				curB++
				if curA == len(a) || curB == len(b) {
					break
				}
				continue
			}
		}
	}
	return c
}

func IntersectionUnSorted(a, b []int) []int {
	c := make([]int, 0, min(len(a), len(b)))
	m := make(map[int]int, min(len(a), len(b)))
	for _, v := range a {
		m[v]++
	}
	for _, v := range b {
		m[v]++
	}
	for k, v := range m {
		if v > 1 {
			c = append(c, k)
		}
	}
	return c
}
