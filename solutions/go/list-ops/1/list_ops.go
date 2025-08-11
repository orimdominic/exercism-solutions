package listops

// IntList is an abstraction of a list of integers which we can define methods on
type IntList []int

func (s IntList) Foldl(fn func(int, int) int, initial int) int {
	n := initial
	for _, v := range s {
		n = fn(n, v)
	}

	return n
}

func (s IntList) Foldr(fn func(int, int) int, initial int) int {
	n := initial
	for i := len(s) - 1; i >= 0; i-- {
		n = fn(s[i], n)
	}

	return n
}

func (s IntList) Filter(fn func(int) bool) IntList {
	l := make(IntList, 0)
	for _, v := range s {
		if fn(v) {
			l = append(l, v)
		}
	}

	return l
}

func (s IntList) Length() int {
	n := 0
	for range s {
		n += 1
	}

	return n
}

func (s IntList) Map(fn func(int) int) IntList {
	l := make(IntList, 0)
	for _, v := range s {
		l = append(l, fn(v))
	}

	return l
}

func (s IntList) Reverse() IntList {
	for i := 0; i < int(len(s)/2); i++ {
		s[i], s[len(s)-i-1] = s[len(s)-i-1], s[i]
	}
	return s
}

func (s IntList) Append(lst IntList) IntList {
	return append(s, lst...)
}

func (s IntList) Concat(lists []IntList) IntList {
	for _, lst := range lists {
		s = append(s, lst...)
	}

	return s
}
