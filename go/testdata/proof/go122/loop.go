package main

func main() {
	var got []int
	var fs []func() int
	for i := range 3 {
		fs = append(fs, func() int { return i })
		got = append(got, i)
	}
	if len(got) != 3 || got[0] != 0 || got[1] != 1 || got[2] != 2 {
		panic("range values")
	}
	if fs[0]() != 0 || fs[1]() != 1 || fs[2]() != 2 {
		panic("range closure")
	}

	var gs []func() int
	for i := range 4 {
		i += 10
		gs = append(gs, func() int { return i })
	}
	if gs[0]() != 10 || gs[1]() != 11 || gs[2]() != 12 || gs[3]() != 13 {
		panic("range mutate")
	}

	var hs []func() int
	for i := 0; i < 3; i++ {
		hs = append(hs, func() int { return i })
	}
	if hs[0]() != 0 || hs[1]() != 1 || hs[2]() != 2 {
		panic("for closure")
	}

	var ps []func() int
	count := 0
	for i := 0; i < 5; i++ {
		ps = append(ps, func() int { return i })
		i++
		count++
	}
	if count != 3 || ps[0]() != 1 || ps[1]() != 3 || ps[2]() != 5 {
		panic("post")
	}

	count = 0
	for i := 0; i < 3; i++ {
		count++
		i = 10
	}
	if count != 1 {
		panic("assign")
	}

	n := int8(2)
	c := 0
	for i := range n {
		c += int(i)
	}
	if c != 1 {
		panic("int8")
	}

	c = 0
	for range 0 {
		c++
	}
	neg := -3
	for range neg {
		c++
	}
	if c != 0 {
		panic("empty")
	}

	var u uint8 = 255
	c = 0
	var last uint8
	for i := range u {
		last = i
		c++
	}
	if c != 255 || last != 254 {
		panic("uint8")
	}

	var k uint8
	sum := 0
	for k = range 4 {
		sum += int(k)
	}
	if sum != 6 || k != 3 {
		panic("const range")
	}

	println("ok")
}
