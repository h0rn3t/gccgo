package main

func main() {
	const c = min(8, 3)
	if c != 3 {
		panic("const")
	}
	if min(3, 1, 2) != 1 || max(3, 1, 2) != 3 {
		panic("int")
	}
	if min(3, 1.5) != 1.5 || max(1, 2.0, 10) != 10.0 {
		panic("const float")
	}
	x := float32(3)
	if max(float32(0), x) != 3 || min(x, float32(1)) != 1 {
		panic("float32")
	}
	zero := 0.0
	nan := zero / zero
	if min(1.0, nan) == 1.0 || max(nan, 1.0) == 1.0 {
		panic("nan")
	}
	if min("b", "a", "c") != "a" || max("", "foo", "bar") != "foo" {
		panic("string")
	}
	a, b := 4, 9
	if min(a, b, 5) != 4 || max(a, b) != b {
		panic("var")
	}
	s1, s2 := "pear", "apple"
	if min(s1, s2) != "apple" || max(s1, s2) != "pear" {
		panic("var string")
	}
	var f float64 = max(1, 2)
	if f != 2 {
		panic("assign")
	}

	m := map[string]int{"a": 1, "b": 2}
	clear(m)
	if len(m) != 0 || m["a"] != 0 {
		panic("clear map")
	}
	var nm map[string]int
	clear(nm)

	s := []int{1, 2, 3}
	clear(s)
	if len(s) != 3 || s[0] != 0 || s[1] != 0 || s[2] != 0 {
		panic("clear slice")
	}
	var ns []int
	clear(ns)

	p0, p1 := new(int), new(int)
	ps := []*int{p0, p1}
	clear(ps)
	if len(ps) != 2 || ps[0] != nil || ps[1] != nil {
		panic("clear ptr slice")
	}

	println("ok")
}
